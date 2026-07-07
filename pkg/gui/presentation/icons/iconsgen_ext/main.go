// Command iconsgen_ext regenerates the nameIconMap and extIconMap tables in
// ../file_icons.go so that lazygit's file icons and colors match what the
// user's neovim renders via mini.icons (the source of truth).
//
// It uses `nvim --headless` (oracle.lua) as an oracle: for every existing key
// it asks mini.icons for the effective glyph + highlight group, maps the
// highlight to a catppuccin-mocha color, and rewrites the map entry. Entries
// mini.icons has no specific icon for (is_default) keep their current lazygit
// icon. A few keys defined only in the personal override are added.
//
// The nerd-font glyphs are written as Go \u / \U escapes with normal file I/O,
// avoiding any editor glyph sanitization. Run it with:
//
//	go generate ./pkg/gui/presentation/icons
//
// Requirements: `nvim` on PATH with mini.icons installed. Paths can be
// overridden via MINI_ICONS_PATH and ICONS_LUA_PATH env vars.
package main

import (
	_ "embed"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

//go:embed oracle.lua
var oracleLua string

// Catppuccin mocha palette. The whole palette is emitted as constants (unused
// package-level constants are legal in Go) so the file is self-documenting.
var mochaPalette = [][2]string{
	{"ctpRosewater", "#F5E0DC"},
	{"ctpFlamingo", "#F2CDCD"},
	{"ctpPink", "#F5C2E7"},
	{"ctpMauve", "#CBA6F7"},
	{"ctpRed", "#F38BA8"},
	{"ctpMaroon", "#EBA0AC"},
	{"ctpPeach", "#FAB387"},
	{"ctpYellow", "#F9E2AF"},
	{"ctpGreen", "#A6E3A1"},
	{"ctpTeal", "#94E2D5"},
	{"ctpSky", "#89DCEB"},
	{"ctpSapphire", "#74C7EC"},
	{"ctpBlue", "#89B4FA"},
	{"ctpLavender", "#B4BEFE"},
	{"ctpText", "#CDD6F4"},
	{"ctpSubtext1", "#BAC2DE"},
	{"ctpSubtext0", "#A6ADC8"},
	{"ctpOverlay2", "#9399B2"},
	{"ctpOverlay1", "#7F849C"},
	{"ctpOverlay0", "#6C7086"},
	{"ctpSurface2", "#585B70"},
	{"ctpSurface1", "#45475A"},
	{"ctpSurface0", "#313244"},
	{"ctpBase", "#1E1E2E"},
	{"ctpMantle", "#181825"},
	{"ctpCrust", "#11111B"},
}

// hlToConst maps a mini.icons highlight group to a catppuccin constant, using
// the effective colors of the user's neovim: catppuccin's mini integration
// with the personal `MiniIconsAzure -> teal` override applied (so both Azure
// and Cyan resolve to teal).
var hlToConst = map[string]string{
	"MiniIconsAzure":  "ctpTeal", // overridden from sapphire to teal by the user
	"MiniIconsBlue":   "ctpBlue",
	"MiniIconsCyan":   "ctpTeal",
	"MiniIconsGreen":  "ctpGreen",
	"MiniIconsGrey":   "ctpText",
	"MiniIconsOrange": "ctpPeach",
	"MiniIconsPurple": "ctpMauve",
	"MiniIconsRed":    "ctpRed",
	"MiniIconsYellow": "ctpYellow",
}

// newKey is a key defined only in the personal mini.icons override that lazygit
// is currently missing. It is inserted immediately before insertBefore so the
// diff stays readable.
type newKey struct {
	mapName      string // "name" or "ext"
	key          string
	insertBefore string
}

var newKeys = []newKey{
	{"name", ".eslintrc.js", ".git"},
	{"name", ".keep", ".luacheckrc"},
	{"name", ".node-version", ".npmignore"},
	{"name", ".yarnrc.yml", ".zprofile"},
	{"name", "devcontainer.json", "docker-compose.yaml"},
	{"name", "tsconfig.build.json", "tsconfig.json"},
	{"ext", ".dbml", ".dconf"},
}

func main() {
	target := findTarget()
	src, err := os.ReadFile(target)
	must(err)

	hexToConst := map[string]string{}
	for _, p := range mochaPalette {
		hexToConst[p[1]] = p[0]
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, target, src, parser.ParseComments)
	must(err)

	nameLit := findMapLiteral(file, "nameIconMap")
	extLit := findMapLiteral(file, "extIconMap")
	if nameLit == nil || extLit == nil {
		fatal("could not locate nameIconMap / extIconMap in %s", target)
	}

	nameFinals := buildFinals(parseEntries(nameLit, hexToConst), "name")
	extFinals := buildFinals(parseEntries(extLit, hexToConst), "ext")

	// Assign query ids and run the oracle.
	all := append(append([]*finalEntry{}, nameFinals...), extFinals...)
	results := runOracle(all)

	kept, reskinned, added := applyResults(all, results)

	// Splice new map bodies in. Replace the later literal (extIconMap) first so
	// byte offsets of the earlier one stay valid.
	src = spliceLiteral(src, fset, extLit, renderMap(extFinals))
	src = spliceLiteral(src, fset, nameLit, renderMap(nameFinals))

	src = insertConstBlock(src)
	src = convertLooseColors(src, hexToConst)

	out, err := format.Source(src)
	if err != nil {
		fatal("gofmt failed on generated source: %v", err)
	}
	must(os.WriteFile(target, out, 0o644))

	fmt.Printf("regenerated %s: %d re-skinned, %d kept (mini.icons default), %d added\n",
		target, reskinned, kept, added)
}

type parsed struct {
	key        string
	iconRune   rune
	colorConst string
}

type finalEntry struct {
	key         string
	resolveName string
	isNew       bool

	// existing values, used when mini.icons has no specific icon (is_default)
	existingRune  rune
	existingColor string

	// resolved output
	iconRune   rune
	colorConst string
}

// findMapLiteral returns the composite literal for `var <name> = map[...]{...}`.
func findMapLiteral(file *ast.File, name string) *ast.CompositeLit {
	var lit *ast.CompositeLit
	for _, decl := range file.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.VAR {
			continue
		}
		for _, spec := range gd.Specs {
			vs, ok := spec.(*ast.ValueSpec)
			if !ok || len(vs.Names) != 1 || vs.Names[0].Name != name || len(vs.Values) != 1 {
				continue
			}
			if cl, ok := vs.Values[0].(*ast.CompositeLit); ok {
				lit = cl
			}
		}
	}
	return lit
}

// parseEntries extracts key -> {icon rune, color constant} from a map literal,
// preserving order. Existing colors are normalized to catppuccin constant names.
func parseEntries(lit *ast.CompositeLit, hexToConst map[string]string) []parsed {
	entries := make([]parsed, 0, len(lit.Elts))
	for _, elt := range lit.Elts {
		kv := elt.(*ast.KeyValueExpr)
		key := unquote(kv.Key.(*ast.BasicLit).Value)
		props := kv.Value.(*ast.CompositeLit)

		var iconRune rune
		var colorConst string
		for _, field := range props.Elts {
			fkv := field.(*ast.KeyValueExpr)
			switch fkv.Key.(*ast.Ident).Name {
			case "Icon":
				runes := []rune(unquote(fkv.Value.(*ast.BasicLit).Value))
				if len(runes) != 1 {
					fatal("icon for %q is not a single rune", key)
				}
				iconRune = runes[0]
			case "Color":
				switch cv := fkv.Value.(type) {
				case *ast.BasicLit: // "#RRGGBB"
					hex := strings.ToUpper(unquote(cv.Value))
					c, ok := hexToConst[hex]
					if !ok {
						fatal("color %s for %q is not in the catppuccin palette", hex, key)
					}
					colorConst = c
				case *ast.Ident: // already a constant (idempotent re-run)
					colorConst = cv.Name
				default:
					fatal("unexpected Color expression for %q", key)
				}
			}
		}
		entries = append(entries, parsed{key: key, iconRune: iconRune, colorConst: colorConst})
	}
	return entries
}

func buildFinals(entries []parsed, mapName string) []*finalEntry {
	finals := make([]*finalEntry, 0, len(entries)+len(newKeys))
	for _, p := range entries {
		finals = append(finals, &finalEntry{
			key:           p.key,
			resolveName:   resolveName(mapName, p.key),
			existingRune:  p.iconRune,
			existingColor: p.colorConst,
		})
	}
	for _, nk := range newKeys {
		if nk.mapName != mapName {
			continue
		}
		if containsKey(finals, nk.key) {
			continue // already present (idempotent re-run)
		}
		fe := &finalEntry{key: nk.key, resolveName: resolveName(mapName, nk.key), isNew: true}
		finals = insertBefore(finals, nk.insertBefore, fe)
	}
	return finals
}

// resolveName maps a lazygit map key to the filename handed to MiniIcons.get.
// Extension keys (".py") become a synthetic filename ("x.py") so mini.icons'
// filetype-detection fallback runs, mirroring a real file's rendering.
func resolveName(mapName, key string) string {
	if mapName == "ext" {
		return "x" + key
	}
	return key
}

func containsKey(finals []*finalEntry, key string) bool {
	for _, f := range finals {
		if f.key == key {
			return true
		}
	}
	return false
}

func insertBefore(finals []*finalEntry, anchor string, fe *finalEntry) []*finalEntry {
	for i, f := range finals {
		if f.key == anchor {
			return append(finals[:i:i], append([]*finalEntry{fe}, finals[i:]...)...)
		}
	}
	fmt.Fprintf(os.Stderr, "warning: anchor %q not found for new key %q; appending\n", anchor, fe.key)
	return append(finals, fe)
}

type result struct {
	glyph     string
	hl        string
	isDefault bool
}

func runOracle(all []*finalEntry) map[int]result {
	tmp, err := os.MkdirTemp("", "iconsgen_ext")
	must(err)
	defer os.RemoveAll(tmp)

	oraclePath := filepath.Join(tmp, "oracle.lua")
	queriesPath := filepath.Join(tmp, "queries.tsv")
	resultsPath := filepath.Join(tmp, "results.tsv")
	must(os.WriteFile(oraclePath, []byte(oracleLua), 0o644))

	var q strings.Builder
	for id, fe := range all {
		fmt.Fprintf(&q, "%d\t%s\n", id, fe.resolveName)
	}
	must(os.WriteFile(queriesPath, []byte(q.String()), 0o644))

	home, err := os.UserHomeDir()
	must(err)
	miniPath := envOr("MINI_ICONS_PATH", filepath.Join(home, ".local/share/nvim/lazy/mini.icons"))
	iconsLua := envOr("ICONS_LUA_PATH", filepath.Join(home,
		"Repositories/personal/dotfiles/nvim/.config/nvim/lua/config/icons.lua"))
	assertExists(miniPath, "mini.icons plugin (set MINI_ICONS_PATH)")
	assertExists(iconsLua, "personal icons.lua (set ICONS_LUA_PATH)")

	cmd := exec.Command("nvim", "--headless", "-l", oraclePath)
	cmd.Env = append(os.Environ(),
		"MINI_ICONS_PATH="+miniPath,
		"ICONS_LUA_PATH="+iconsLua,
		"QUERIES_PATH="+queriesPath,
		"RESULTS_PATH="+resultsPath,
	)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fatal("running nvim oracle: %v", err)
	}

	raw, err := os.ReadFile(resultsPath)
	must(err)
	results := map[int]result{}
	for line := range strings.SplitSeq(strings.TrimRight(string(raw), "\n"), "\n") {
		if line == "" {
			continue
		}
		f := strings.Split(line, "\t")
		if len(f) != 4 {
			fatal("malformed oracle result line: %q", line)
		}
		id, err := strconv.Atoi(f[0])
		must(err)
		results[id] = result{glyph: f[1], hl: f[2], isDefault: f[3] == "true"}
	}
	if len(results) != len(all) {
		fatal("oracle returned %d results for %d queries", len(results), len(all))
	}
	return results
}

func applyResults(all []*finalEntry, results map[int]result) (kept, reskinned, added int) {
	for id, fe := range all {
		r := results[id]
		switch {
		case fe.isNew:
			fe.iconRune = singleRune(r.glyph, fe.key)
			fe.colorConst = hlConst(r.hl, fe.key)
			added++
		case r.isDefault:
			// mini.icons has no specific icon; keep lazygit's existing one.
			fe.iconRune = fe.existingRune
			fe.colorConst = fe.existingColor
			kept++
		default:
			fe.iconRune = singleRune(r.glyph, fe.key)
			fe.colorConst = hlConst(r.hl, fe.key)
			reskinned++
		}
	}
	return kept, reskinned, added
}

func renderMap(finals []*finalEntry) string {
	var b strings.Builder
	b.WriteString("{\n")
	for _, fe := range finals {
		fmt.Fprintf(&b, "\t%s: {Icon: \"%s\", Color: %s}, // %s\n",
			strconv.Quote(fe.key), iconEscape(fe.iconRune), fe.colorConst, string(fe.iconRune))
	}
	b.WriteString("}")
	return b.String()
}

// iconEscape formats a rune as the file's \u (BMP) or \U (astral) escape style.
func iconEscape(r rune) string {
	if r <= 0xFFFF {
		return fmt.Sprintf(`\u%04x`, r)
	}
	return fmt.Sprintf(`\U%08x`, r)
}

func spliceLiteral(src []byte, fset *token.FileSet, lit *ast.CompositeLit, body string) []byte {
	lb := fset.Position(lit.Lbrace).Offset
	rb := fset.Position(lit.Rbrace).Offset
	out := make([]byte, 0, len(src)+len(body))
	out = append(out, src[:lb]...)
	out = append(out, body...)
	out = append(out, src[rb+1:]...)
	return out
}

func insertConstBlock(src []byte) []byte {
	if strings.Contains(string(src), `ctpRosewater = "#`) {
		return src // const block already present (idempotent re-run)
	}
	var b strings.Builder
	b.WriteString("//go:generate go run ./iconsgen_ext\n\n")
	b.WriteString("const (\n")
	for _, p := range mochaPalette {
		fmt.Fprintf(&b, "\t%s = %q\n", p[0], p[1])
	}
	b.WriteString(")\n\n")

	anchor := "var (\n\tDEFAULT_FILE_ICON"
	idx := strings.Index(string(src), anchor)
	if idx < 0 {
		fatal("could not find insertion anchor %q", anchor)
	}
	out := make([]byte, 0, len(src)+b.Len())
	out = append(out, src[:idx]...)
	out = append(out, b.String()...)
	out = append(out, src[idx:]...)
	return out
}

// convertLooseColors rewrites `Color: "#RRGGBB"` literals outside the two maps
// (the DEFAULT_* vars and the NerdFonts v2 patch) into catppuccin constants.
// The positional #4E4E4E linked-worktree color has no `Color:` label and is
// intentionally left as a literal (not part of the palette).
func convertLooseColors(src []byte, hexToConst map[string]string) []byte {
	re := regexp.MustCompile(`Color: "(#[0-9A-Fa-f]{6})"`)
	return re.ReplaceAllFunc(src, func(m []byte) []byte {
		hex := strings.ToUpper(string(re.FindSubmatch(m)[1]))
		if c, ok := hexToConst[hex]; ok {
			return []byte("Color: " + c)
		}
		return m
	})
}

func singleRune(glyph, key string) rune {
	runes := []rune(glyph)
	if len(runes) != 1 {
		fatal("mini.icons glyph %q for %q is not a single rune (%d runes)", glyph, key, len(runes))
	}
	return runes[0]
}

func hlConst(hl, key string) string {
	c, ok := hlToConst[hl]
	if !ok {
		fatal("unmapped highlight group %q for %q", hl, key)
	}
	return c
}

func findTarget() string {
	for _, cand := range []string{"file_icons.go", "pkg/gui/presentation/icons/file_icons.go"} {
		if _, err := os.Stat(cand); err == nil {
			return cand
		}
	}
	fatal("could not find file_icons.go; run via `go generate ./pkg/gui/presentation/icons`")
	return ""
}

func unquote(s string) string {
	v, err := strconv.Unquote(s)
	must(err)
	return v
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func assertExists(path, what string) {
	if _, err := os.Stat(path); err != nil {
		fatal("%s not found at %s", what, path)
	}
}

func must(err error) {
	if err != nil {
		fatal("%v", err)
	}
}

func fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "iconsgen_ext: "+format+"\n", args...)
	os.Exit(1)
}
