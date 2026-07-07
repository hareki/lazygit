# iconsgen_ext

Regenerates the `nameIconMap` / `extIconMap` tables (and the catppuccin color
constants) in [`../file_icons.go`](../file_icons.go) so lazygit's file icons and
colors match what neovim renders via **mini.icons** (the source of truth).

`file_icons.go`'s two maps are **generated** — don't hand-edit them. Change the
neovim config (below) and re-run this generator instead.

## How it works

`nvim --headless` runs [`oracle.lua`](./oracle.lua), which loads mini.icons,
applies the personal override, and reports the effective glyph + highlight group
for each key. The Go driver ([`main.go`](./main.go)) then:

- **re-skins** every existing key mini.icons resolves (new glyph + color),
- **keeps** lazygit's current icon where mini.icons has only its generic default,
- **adds** the few keys defined only in the personal override,
- maps mini.icons highlight groups to catppuccin-mocha colors, with the personal
  `MiniIconsAzure → teal` override applied (so Azure and Cyan both become teal).

Glyphs are written as Go `\u` / `\U` escapes via normal file I/O, so no editor
glyph sanitization occurs.

## Sources of truth (read, never modified)

- mini.icons plugin — default icon tables + resolver.
- `.../nvim/lua/core/mini-icons.lua` — personal override. **If you change this
  file's icon tables or the Azure→teal recolor, mirror it in `oracle.lua` /
  `hlToConst` before regenerating.**
- `.../nvim/lua/config/icons.lua` — glyph values the override references (read
  directly by `oracle.lua`, so glyph edits there flow through automatically).

## Run

Requires `nvim` on `PATH` with mini.icons installed.

```sh
# from the repo root
go generate ./pkg/gui/presentation/icons
# or directly
go run ./pkg/gui/presentation/icons/iconsgen_ext
```

Non-default plugin/config locations can be overridden:

```sh
MINI_ICONS_PATH=/path/to/mini.icons \
ICONS_LUA_PATH=/path/to/nvim/lua/config/icons.lua \
  go generate ./pkg/gui/presentation/icons
```

The generator is idempotent — re-running with an unchanged neovim config
produces no diff. Re-run it after an upstream merge that rewrites
`file_icons.go` to re-apply the sync.

## Verify

```sh
gofmt -l pkg/gui/presentation/icons/file_icons.go   # expect no output
go test ./pkg/gui/presentation/icons/...            # single-rune invariant
```
