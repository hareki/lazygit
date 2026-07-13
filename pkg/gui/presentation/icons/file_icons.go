package icons

import (
	"path/filepath"
	"strings"

	"github.com/jesseduffield/lazygit/pkg/config"
)

// NOTE: Visit next links for inspiration:
// https://github.com/eza-community/eza/blob/main/src/output/icons.rs
// https://github.com/nvim-tree/nvim-web-devicons/tree/master/lua/nvim-web-devicons/default

//go:generate go run ./iconsgen_ext

const (
	ctpRosewater = "#F5E0DC"
	ctpFlamingo  = "#F2CDCD"
	ctpPink      = "#F5C2E7"
	ctpMauve     = "#CBA6F7"
	ctpRed       = "#F38BA8"
	ctpMaroon    = "#EBA0AC"
	ctpPeach     = "#FAB387"
	ctpYellow    = "#F9E2AF"
	ctpGreen     = "#A6E3A1"
	ctpTeal      = "#94E2D5"
	ctpSky       = "#89DCEB"
	ctpSapphire  = "#74C7EC"
	ctpBlue      = "#89B4FA"
	ctpLavender  = "#B4BEFE"
	ctpText      = "#CDD6F4"
	ctpSubtext1  = "#BAC2DE"
	ctpSubtext0  = "#A6ADC8"
	ctpOverlay2  = "#9399B2"
	ctpOverlay1  = "#7F849C"
	ctpOverlay0  = "#6C7086"
	ctpSurface2  = "#585B70"
	ctpSurface1  = "#45475A"
	ctpSurface0  = "#313244"
	ctpBase      = "#1E1E2E"
	ctpMantle    = "#181825"
	ctpCrust     = "#11111B"
)

var (
	DEFAULT_FILE_ICON      = IconProperties{Icon: "\uf15b", Color: ctpText}      // 
	DEFAULT_SUBMODULE_ICON = IconProperties{Icon: "\U000f02a2", Color: ctpPeach} // 󰊢
	DEFAULT_DIRECTORY_ICON = IconProperties{Icon: "\uf07b", Color: ctpBlue}      // 
)

// NOTE: The filename map is case sensitive.
var nameIconMap = map[string]IconProperties{
	".atom":                      {Icon: "\U000f05c0", Color: ctpPeach},    // 󰗀
	".babelrc":                   {Icon: "\U000f0626", Color: ctpYellow},   // 󰘦
	".bash_profile":              {Icon: "\U000f0493", Color: ctpGreen},    // 󰒓
	".bashprofile":               {Icon: "\ue615", Color: ctpGreen},        // 
	".bashrc":                    {Icon: "\U000f0493", Color: ctpGreen},    // 󰒓
	".clang-format":              {Icon: "\ue6a8", Color: ctpMauve},        // 
	".clang-tidy":                {Icon: "\ue6a8", Color: ctpMauve},        // 
	".codespellrc":               {Icon: "\U000f04c6", Color: ctpGreen},    // 󰓆
	".condarc":                   {Icon: "\ue6a8", Color: ctpMauve},        // 
	".dockerignore":              {Icon: "\U000f0868", Color: ctpPeach},    // 󰡨
	".ds_store":                  {Icon: "\uf302", Color: ctpSubtext0},     // 
	".editorconfig":              {Icon: "\ue652", Color: ctpText},         // 
	".env":                       {Icon: "\ueb52", Color: ctpYellow},       // 
	".eslintignore":              {Icon: "\U000f0c7a", Color: ctpSubtext0}, // 󰱺
	".eslintrc":                  {Icon: "\U000f0626", Color: ctpYellow},   // 󰘦
	".eslintrc.js":               {Icon: "\U000f0c7a", Color: ctpYellow},   // 󰱺
	".git":                       {Icon: "\U000f02a2", Color: ctpPeach},    // 󰊢
	".git-blame-ignore-revs":     {Icon: "\U000f02a2", Color: ctpPeach},    // 󰊢
	".gitattributes":             {Icon: "\U000f02a2", Color: ctpYellow},   // 󰊢
	".gitconfig":                 {Icon: "\U000f0493", Color: ctpPeach},    // 󰒓
	".github":                    {Icon: "\uf408", Color: ctpSubtext0},     // 
	".gitignore":                 {Icon: "\U000f02a2", Color: ctpMauve},    // 󰊢
	".gitlab-ci.yml":             {Icon: "\U000f0ba0", Color: ctpPeach},    // 󰮠
	".gitmodules":                {Icon: "\U000f0493", Color: ctpPeach},    // 󰒓
	".gtkrc-2.0":                 {Icon: "\U000f0493", Color: ctpTeal},     // 󰒓
	".gvimrc":                    {Icon: "\ue7c5", Color: ctpGreen},        // 
	".idea":                      {Icon: "\ue7b5", Color: ctpSubtext0},     // 
	".justfile":                  {Icon: "\U000f05b7", Color: ctpPeach},    // 󰖷
	".keep":                      {Icon: "\U000f02a2", Color: ctpText},     // 󰊢
	".luacheckrc":                {Icon: "\U000f08b1", Color: ctpTeal},     // 󰢱
	".luaurc":                    {Icon: "\U000f0626", Color: ctpYellow},   // 󰘦
	".mailmap":                   {Icon: "\U000f02a2", Color: ctpTeal},     // 󰊢
	".nanorc":                    {Icon: "\U000f0493", Color: ctpYellow},   // 󰒓
	".node-version":              {Icon: "\ue718", Color: ctpGreen},        // 
	".npmignore":                 {Icon: "\U000f0493", Color: ctpText},     // 󰒓
	".npmrc":                     {Icon: "\U000f0bc2", Color: ctpTeal},     // 󰯂
	".nuxtrc":                    {Icon: "\U000f1106", Color: ctpSapphire}, // 󱄆
	".nvmrc":                     {Icon: "\U000f0493", Color: ctpGreen},    // 󰒓
	".pre-commit-config.yaml":    {Icon: "\ue6a8", Color: ctpMauve},        // 
	".prettierignore":            {Icon: "\ue6b4", Color: ctpPeach},        // 
	".prettierrc":                {Icon: "\ue6b4", Color: ctpMauve},        // 
	".prettierrc.json":           {Icon: "\U000f0626", Color: ctpYellow},   // 󰘦
	".prettierrc.json5":          {Icon: "\U000f0626", Color: ctpPeach},    // 󰘦
	".prettierrc.toml":           {Icon: "\ue6b2", Color: ctpPeach},        // 
	".prettierrc.yaml":           {Icon: "\ue6a8", Color: ctpMauve},        // 
	".prettierrc.yml":            {Icon: "\ue6a8", Color: ctpMauve},        // 
	".pylintrc":                  {Icon: "\U000f0bc2", Color: ctpTeal},     // 󰯂
	".rvm":                       {Icon: "\ue21e", Color: ctpRed},          // 
	".settings.json":             {Icon: "\U000f0626", Color: ctpYellow},   // 󰘦
	".SRCINFO":                   {Icon: "\uf129", Color: ctpSapphire},     // 
	".tmux.conf":                 {Icon: "\U000f0493", Color: ctpText},     // 󰒓
	".tmux.conf.local":           {Icon: "\U000f0493", Color: ctpGreen},    // 󰒓
	".Trash":                     {Icon: "\uf1f8", Color: ctpLavender},     // 
	".vimrc":                     {Icon: "\ue7c5", Color: ctpGreen},        // 
	".vscode":                    {Icon: "\ue70c", Color: ctpSapphire},     // 
	".Xauthority":                {Icon: "\uf369", Color: ctpPeach},        // 
	".Xresources":                {Icon: "\U000f0493", Color: ctpBlue},     // 󰒓
	".xinitrc":                   {Icon: "\U000f0493", Color: ctpBlue},     // 󰒓
	".xsession":                  {Icon: "\uf369", Color: ctpPeach},        // 
	".yarnrc.yml":                {Icon: "\ue6a7", Color: ctpBlue},         // 
	".zprofile":                  {Icon: "\ue691", Color: ctpGreen},        // 
	".zshenv":                    {Icon: "\ue691", Color: ctpGreen},        // 
	".zshrc":                     {Icon: "\U000f0493", Color: ctpGreen},    // 󰒓
	"_gvimrc":                    {Icon: "\ue7c5", Color: ctpGreen},        // 
	"_vimrc":                     {Icon: "\ue7c5", Color: ctpGreen},        // 
	"AUTHORS":                    {Icon: "\U000f09a8", Color: ctpText},     // 󰦨
	"AUTHORS.txt":                {Icon: "\U000f09a8", Color: ctpText},     // 󰦨
	"bin":                        {Icon: "\U000f12a7", Color: ctpSapphire}, // 󱊧
	"brewfile":                   {Icon: "\ue791", Color: ctpRed},          // 
	"bspwmrc":                    {Icon: "\uf355", Color: ctpSubtext0},     // 
	"BUILD":                      {Icon: "\ue63a", Color: ctpGreen},        // 
	"build.gradle":               {Icon: "\ue775", Color: ctpTeal},         // 
	"build.zig.zon":              {Icon: "\ue6a9", Color: ctpPeach},        // 
	"bun.lockb":                  {Icon: "\ue76f", Color: ctpRosewater},    // 
	"cantorrc":                   {Icon: "\uf373", Color: ctpSapphire},     // 
	"Cargo.lock":                 {Icon: "\ue6b2", Color: ctpPeach},        // 
	"Cargo.toml":                 {Icon: "\ue6b2", Color: ctpPeach},        // 
	"checkhealth":                {Icon: "\U000f04d9", Color: ctpBlue},     // 󰓙
	"CMakeLists.txt":             {Icon: "\U000f1064", Color: ctpPeach},    // 󱁤
	"CODE_OF_CONDUCT":            {Icon: "\U000f10f1", Color: ctpRed},      // 󱃱
	"CODE_OF_CONDUCT.md":         {Icon: "\U000f10f1", Color: ctpRed},      // 󱃱
	"CODE-OF-CONDUCT.md":         {Icon: "\U000f0354", Color: ctpText},     // 󰍔
	"commit_editmsg":             {Icon: "\ue702", Color: ctpPeach},        // 
	"COMMIT_EDITMSG":             {Icon: "\U000f02a2", Color: ctpGreen},    // 󰊢
	"commitlint.config.js":       {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	"commitlint.config.ts":       {Icon: "\U000f06e6", Color: ctpTeal},     // 󰛦
	"compose.yaml":               {Icon: "\ue6a8", Color: ctpMauve},        // 
	"compose.yml":                {Icon: "\ue6a8", Color: ctpMauve},        // 
	"config":                     {Icon: "\uf013", Color: ctpSubtext0},     // 
	"containerfile":              {Icon: "\uf21f", Color: ctpSapphire},     // 
	"copying":                    {Icon: "\U000f0124", Color: ctpPeach},    // 󰄤
	"copying.lesser":             {Icon: "\ue60a", Color: ctpPeach},        // 
	"devcontainer.json":          {Icon: "\uf4b7", Color: ctpTeal},         // 
	"docker-compose.yaml":        {Icon: "\ue6a8", Color: ctpMauve},        // 
	"docker-compose.yml":         {Icon: "\ue6a8", Color: ctpMauve},        // 
	"dockerfile":                 {Icon: "\U000f0868", Color: ctpBlue},     // 󰡨
	"Dockerfile":                 {Icon: "\U000f0868", Color: ctpBlue},     // 󰡨
	"ds_store":                   {Icon: "\uf179", Color: ctpRosewater},    // 
	"eslint.config.cjs":          {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	"eslint.config.js":           {Icon: "\U000f0c7a", Color: ctpYellow},   // 󰱺
	"eslint.config.mjs":          {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	"eslint.config.ts":           {Icon: "\U000f06e6", Color: ctpTeal},     // 󰛦
	"ext_typoscript_setup.txt":   {Icon: "\U000f09a8", Color: ctpText},     // 󰦨
	"favicon.ico":                {Icon: "\uf03e", Color: ctpGreen},        // 
	"fp-info-cache":              {Icon: "\uf34c", Color: ctpRosewater},    // 
	"fp-lib-table":               {Icon: "\uf34c", Color: ctpRosewater},    // 
	"FreeCAD.conf":               {Icon: "\U000f0493", Color: ctpText},     // 󰒓
	"gemfile$":                   {Icon: "\ue791", Color: ctpRed},          // 
	"gitignore_global":           {Icon: "\U000f02a2", Color: ctpPeach},    // 󰊢
	"gnumakefile":                {Icon: "\U000f1064", Color: ctpText},     // 󱁤
	"GNUmakefile":                {Icon: "\U000f1064", Color: ctpText},     // 󱁤
	"go.mod":                     {Icon: "\U000f0afa", Color: ctpTeal},     // 󰫺
	"go.sum":                     {Icon: "\U000f07d3", Color: ctpTeal},     // 󰟓
	"go.work":                    {Icon: "\U000f07d3", Color: ctpMauve},    // 󰟓
	"gradle":                     {Icon: "\ue660", Color: ctpSapphire},     // 
	"gradle-wrapper.properties":  {Icon: "\U000f0b37", Color: ctpGreen},    // 󰬷
	"gradle.properties":          {Icon: "\U000f0b37", Color: ctpGreen},    // 󰬷
	"gradlew":                    {Icon: "\ue660", Color: ctpSapphire},     // 
	"gruntfile.babel.js":         {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	"gruntfile.coffee":           {Icon: "\ue611", Color: ctpPeach},        // 
	"gruntfile.js":               {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	"gruntfile.ls":               {Icon: "\ue611", Color: ctpPeach},        // 
	"gruntfile.ts":               {Icon: "\U000f06e6", Color: ctpTeal},     // 󰛦
	"gtkrc":                      {Icon: "\U000f0493", Color: ctpTeal},     // 󰒓
	"gulpfile.babel.js":          {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	"gulpfile.coffee":            {Icon: "\ue610", Color: ctpRed},          // 
	"gulpfile.js":                {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	"gulpfile.ls":                {Icon: "\ue610", Color: ctpRed},          // 
	"gulpfile.ts":                {Icon: "\U000f06e6", Color: ctpTeal},     // 󰛦
	"hidden":                     {Icon: "\uf023", Color: ctpSubtext0},     // 
	"hypridle.conf":              {Icon: "\U000f0493", Color: ctpText},     // 󰒓
	"hyprland.conf":              {Icon: "\U000f0493", Color: ctpText},     // 󰒓
	"hyprlock.conf":              {Icon: "\U000f0493", Color: ctpText},     // 󰒓
	"hyprpaper.conf":             {Icon: "\U000f0493", Color: ctpText},     // 󰒓
	"i3blocks.conf":              {Icon: "\U000f0493", Color: ctpText},     // 󰒓
	"i3status.conf":              {Icon: "\U000f0493", Color: ctpText},     // 󰒓
	"include":                    {Icon: "\ue5fc", Color: ctpRosewater},    // 
	"index.theme":                {Icon: "\U000f031f", Color: ctpMauve},    // 󰌟
	"ionic.config.json":          {Icon: "\U000f0626", Color: ctpYellow},   // 󰘦
	"justfile":                   {Icon: "\U000f05b7", Color: ctpPeach},    // 󰖷
	"kalgebrarc":                 {Icon: "\uf373", Color: ctpSapphire},     // 
	"kdeglobals":                 {Icon: "\uf373", Color: ctpSapphire},     // 
	"kdenlive-layoutsrc":         {Icon: "\uf33c", Color: ctpBlue},         // 
	"kdenliverc":                 {Icon: "\uf33c", Color: ctpBlue},         // 
	"kritadisplayrc":             {Icon: "\uf33d", Color: ctpMauve},        // 
	"kritarc":                    {Icon: "\uf33d", Color: ctpMauve},        // 
	"lib":                        {Icon: "\U000f1517", Color: ctpGreen},    // 󱔗
	"LICENSE":                    {Icon: "\ue60a", Color: ctpTeal},         // 
	"LICENSE.md":                 {Icon: "\ue60a", Color: ctpTeal},         // 
	"localized":                  {Icon: "\uf179", Color: ctpRosewater},    // 
	"lxde-rc.xml":                {Icon: "\U000f05c0", Color: ctpPeach},    // 󰗀
	"lxqt.conf":                  {Icon: "\U000f0493", Color: ctpText},     // 󰒓
	"Makefile":                   {Icon: "\U000f1064", Color: ctpText},     // 󱁤
	"mix.lock":                   {Icon: "\ue62d", Color: ctpMauve},        // 
	"mpv.conf":                   {Icon: "\U000f0493", Color: ctpText},     // 󰒓
	"node_modules":               {Icon: "\ue718", Color: ctpRed},          // 
	"npmignore":                  {Icon: "\ue71e", Color: ctpRed},          // 
	"nuxt.config.cjs":            {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	"nuxt.config.js":             {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	"nuxt.config.mjs":            {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	"nuxt.config.ts":             {Icon: "\U000f06e6", Color: ctpTeal},     // 󰛦
	"package-lock.json":          {Icon: "\U000f0626", Color: ctpYellow},   // 󰘦
	"package.json":               {Icon: "\ue718", Color: ctpGreen},        // 
	"PKGBUILD":                   {Icon: "\U000f1064", Color: ctpMauve},    // 󱁤
	"platformio.ini":             {Icon: "\U000f0bc2", Color: ctpTeal},     // 󰯂
	"pom.xml":                    {Icon: "\U000f05c0", Color: ctpPeach},    // 󰗀
	"prettier.config.cjs":        {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	"prettier.config.js":         {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	"prettier.config.mjs":        {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	"prettier.config.ts":         {Icon: "\U000f06e6", Color: ctpTeal},     // 󰛦
	"PrusaSlicer.ini":            {Icon: "\U000f0bc2", Color: ctpTeal},     // 󰯂
	"PrusaSlicerGcodeViewer.ini": {Icon: "\U000f0bc2", Color: ctpTeal},     // 󰯂
	"py.typed":                   {Icon: "\ue606", Color: ctpPeach},        // 
	"QtProject.conf":             {Icon: "\U000f0493", Color: ctpText},     // 󰒓
	"R":                          {Icon: "\U000f07d4", Color: ctpSapphire}, // 󰟔
	"README":                     {Icon: "\uf4ed", Color: ctpYellow},       // 
	"README.md":                  {Icon: "\uf4ed", Color: ctpYellow},       // 
	"robots.txt":                 {Icon: "\U000f06a9", Color: ctpText},     // 󰚩
	"rubydoc":                    {Icon: "\ue73b", Color: ctpRed},          // 
	"SECURITY":                   {Icon: "\U000f0483", Color: ctpSubtext1}, // 󰒃
	"SECURITY.md":                {Icon: "\U000f0354", Color: ctpText},     // 󰍔
	"settings.gradle":            {Icon: "\ue775", Color: ctpTeal},         // 
	"svelte.config.js":           {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	"sxhkdrc":                    {Icon: "\uf355", Color: ctpSubtext0},     // 
	"sym-lib-table":              {Icon: "\uf34c", Color: ctpRosewater},    // 
	"tailwind.config.js":         {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	"tailwind.config.mjs":        {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	"tailwind.config.ts":         {Icon: "\U000f06e6", Color: ctpTeal},     // 󰛦
	"tmux.conf":                  {Icon: "\U000f0493", Color: ctpText},     // 󰒓
	"tmux.conf.local":            {Icon: "\U000f0493", Color: ctpGreen},    // 󰒓
	"tsconfig.build.json":        {Icon: "\ue8ca", Color: ctpTeal},         // 
	"tsconfig.json":              {Icon: "\ue8ca", Color: ctpTeal},         // 
	"unlicense":                  {Icon: "\ue60a", Color: ctpPeach},        // 
	"vagrantfile$":               {Icon: "\uf2b8", Color: ctpSapphire},     // 
	"vlcrc":                      {Icon: "\U000f057c", Color: ctpPeach},    // 󰕼
	"webpack":                    {Icon: "\U000f072b", Color: ctpSapphire}, // 󰜫
	"weston.ini":                 {Icon: "\U000f0bc2", Color: ctpTeal},     // 󰯂
	"WORKSPACE":                  {Icon: "\ue63a", Color: ctpGreen},        // 
	"WORKSPACE.bzlmod":           {Icon: "\ue63a", Color: ctpGreen},        // 
	"xmobarrc":                   {Icon: "\uf35e", Color: ctpRed},          // 
	"xmobarrc.hs":                {Icon: "\U000f0c92", Color: ctpMauve},    // 󰲒
	"xmonad.hs":                  {Icon: "\U000f0c92", Color: ctpMauve},    // 󰲒
	"xorg.conf":                  {Icon: "\U000f0493", Color: ctpText},     // 󰒓
	"xsettingsd.conf":            {Icon: "\U000f0493", Color: ctpText},     // 󰒓
	"yarn.lock":                  {Icon: "\ue6a7", Color: ctpBlue},         // 
}

var extIconMap = map[string]IconProperties{
	".3gp":            {Icon: "\U000f022b", Color: ctpYellow},   // 󰈫
	".3mf":            {Icon: "\U000f01a7", Color: ctpSubtext0}, // 󰆧
	".7z":             {Icon: "\U000f05c4", Color: ctpBlue},     // 󰗄
	".DS_store":       {Icon: "\uf179", Color: ctpSubtext0},     // 
	".a":              {Icon: "\ue637", Color: ctpMauve},        // 
	".aac":            {Icon: "\U000f0223", Color: ctpYellow},   // 󰈣
	".adb":            {Icon: "\U000f1077", Color: ctpTeal},     // 󱁷
	".ads":            {Icon: "\U000f1077", Color: ctpTeal},     // 󱁷
	".ai":             {Icon: "\ue7b4", Color: ctpYellow},       // 
	".aif":            {Icon: "\U000f0223", Color: ctpTeal},     // 󰈣
	".aiff":           {Icon: "\U000f0386", Color: ctpRed},      // 󰎆
	".android":        {Icon: "\ue70e", Color: ctpGreen},        // 
	".ape":            {Icon: "\uf001", Color: ctpSapphire},     // 
	".apk":            {Icon: "\ue70e", Color: ctpGreen},        // 
	".app":            {Icon: "\ueae8", Color: ctpRed},          // 
	".apple":          {Icon: "\ue635", Color: ctpSubtext0},     // 
	".applescript":    {Icon: "\U000f0035", Color: ctpYellow},   // 󰀵
	".asc":            {Icon: "\U000f0306", Color: ctpSapphire}, // 󰌆
	".asm":            {Icon: "\ue637", Color: ctpMauve},        // 
	".ass":            {Icon: "\U000f0a16", Color: ctpPeach},    // 󰨖
	".astro":          {Icon: "\ue6b3", Color: ctpPeach},        // 
	".avi":            {Icon: "\U000f022b", Color: ctpText},     // 󰈫
	".avif":           {Icon: "\U000f021f", Color: ctpSapphire}, // 󰈟
	".avro":           {Icon: "\ue60b", Color: ctpPeach},        // 
	".awk":            {Icon: "\ue691", Color: ctpText},         // 
	".azcli":          {Icon: "\uebd8", Color: ctpSapphire},     // 
	".bak":            {Icon: "\U000f006f", Color: ctpSubtext0}, // 󰁯
	".bash":           {Icon: "\ue691", Color: ctpText},         // 
	".bash_history":   {Icon: "\ue795", Color: ctpGreen},        // 
	".bash_profile":   {Icon: "\ue795", Color: ctpGreen},        // 
	".bashrc":         {Icon: "\ue795", Color: ctpGreen},        // 
	".bat":            {Icon: "\U000f0bc2", Color: ctpGreen},    // 󰯂
	".bats":           {Icon: "\ue691", Color: ctpText},         // 
	".bazel":          {Icon: "\ue63a", Color: ctpGreen},        // 
	".bib":            {Icon: "\U000f125f", Color: ctpYellow},   // 󱉟
	".bicep":          {Icon: "\ue63b", Color: ctpTeal},         // 
	".bicepparam":     {Icon: "\ue63b", Color: ctpMauve},        // 
	".blade.php":      {Icon: "\U000f0ad0", Color: ctpRed},      // 󰫐
	".blend":          {Icon: "\U000f00ab", Color: ctpPeach},    // 󰂫
	".blp":            {Icon: "\U000f0821", Color: ctpBlue},     // 󰠡
	".bmp":            {Icon: "\U000f021f", Color: ctpGreen},    // 󰈟
	".brep":           {Icon: "\U000f0eeb", Color: ctpGreen},    // 󰻫
	".bz":             {Icon: "\U000f05c4", Color: ctpPeach},    // 󰗄
	".bz2":            {Icon: "\U000f05c4", Color: ctpPeach},    // 󰗄
	".bz3":            {Icon: "\U000f05c4", Color: ctpPeach},    // 󰗄
	".bzl":            {Icon: "\ue63a", Color: ctpGreen},        // 
	".c":              {Icon: "\U000f0671", Color: ctpBlue},     // 󰙱
	".c++":            {Icon: "\U000f0672", Color: ctpTeal},     // 󰙲
	".cab":            {Icon: "\ue70f", Color: ctpSubtext0},     // 
	".cache":          {Icon: "\uf49b", Color: ctpRosewater},    // 
	".cast":           {Icon: "\U000f022b", Color: ctpRed},      // 󰈫
	".cbl":            {Icon: "\U000f133c", Color: ctpBlue},     // 󱌼
	".cc":             {Icon: "\U000f0672", Color: ctpTeal},     // 󰙲
	".ccm":            {Icon: "\U000f0672", Color: ctpTeal},     // 󰙲
	".cfg":            {Icon: "\U000f0493", Color: ctpBlue},     // 󰒓
	".cjs":            {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	".class":          {Icon: "\U000f076b", Color: ctpRed},      // 󰝫
	".clj":            {Icon: "\ue768", Color: ctpGreen},        // 
	".cljc":           {Icon: "\ue768", Color: ctpGreen},        // 
	".cljd":           {Icon: "\ue76a", Color: ctpSapphire},     // 
	".cljs":           {Icon: "\ue768", Color: ctpGreen},        // 
	".cls":            {Icon: "\U000f011a", Color: ctpPeach},    // 󰄚
	".cmake":          {Icon: "\U000f1064", Color: ctpPeach},    // 󱁤
	".cmd":            {Icon: "\U000f0bc2", Color: ctpGreen},    // 󰯂
	".cob":            {Icon: "\U000f133c", Color: ctpBlue},     // 󱌼
	".cobol":          {Icon: "\u2699", Color: ctpSapphire},     // ⚙
	".coffee":         {Icon: "\ue61b", Color: ctpSubtext0},     // 
	".conda":          {Icon: "\ue715", Color: ctpGreen},        // 
	".conf":           {Icon: "\U000f0493", Color: ctpText},     // 󰒓
	".config.ru":      {Icon: "\U000f0d2d", Color: ctpRed},      // 󰴭
	".cp":             {Icon: "\ue646", Color: ctpSapphire},     // 
	".cpio":           {Icon: "\uf410", Color: ctpPeach},        // 
	".cpp":            {Icon: "\U000f0672", Color: ctpTeal},     // 󰙲
	".cppm":           {Icon: "\U000f0672", Color: ctpTeal},     // 󰙲
	".cpy":            {Icon: "\U000f133c", Color: ctpBlue},     // 󱌼
	".cr":             {Icon: "\ue62f", Color: ctpText},         // 
	".crdownload":     {Icon: "\uf019", Color: ctpSapphire},     // 
	".cs":             {Icon: "\U000f031b", Color: ctpGreen},    // 󰌛
	".csh":            {Icon: "\ue691", Color: ctpText},         // 
	".cshtml":         {Icon: "\U000f0214", Color: ctpText},     // 󰈔
	".cson":           {Icon: "\ue61b", Color: ctpSubtext0},     // 
	".csproj":         {Icon: "\U000f05c0", Color: ctpPeach},    // 󰗀
	".css":            {Icon: "\U000f031c", Color: ctpTeal},     // 󰌜
	".csv":            {Icon: "\ue64a", Color: ctpGreen},        // 
	".csx":            {Icon: "\U000f031b", Color: ctpGreen},    // 󰌛
	".cts":            {Icon: "\U000f06e6", Color: ctpTeal},     // 󰛦
	".cu":             {Icon: "\ue64b", Color: ctpGreen},        // 
	".cue":            {Icon: "\U000f075a", Color: ctpYellow},   // 󰝚
	".cuh":            {Icon: "\ue64b", Color: ctpGreen},        // 
	".cxx":            {Icon: "\U000f0672", Color: ctpTeal},     // 󰙲
	".cxxm":           {Icon: "\U000f0672", Color: ctpTeal},     // 󰙲
	".d":              {Icon: "\ue7af", Color: ctpGreen},        // 
	".d.ts":           {Icon: "\U000f06e6", Color: ctpTeal},     // 󰛦
	".dart":           {Icon: "\ue798", Color: ctpBlue},         // 
	".db":             {Icon: "\uf1c0", Color: ctpPeach},        // 
	".dbml":           {Icon: "\U000f01bc", Color: ctpTeal},     // 󰆼
	".dconf":          {Icon: "\ue706", Color: ctpRosewater},    // 
	".deb":            {Icon: "\uebc5", Color: ctpRed},          // 
	".desktop":        {Icon: "\U000f0379", Color: ctpMauve},    // 󰍹
	".diff":           {Icon: "\U000f0993", Color: ctpRed},      // 󰦓
	".djvu":           {Icon: "\uf02d", Color: ctpSubtext0},     // 
	".dll":            {Icon: "\U000f107c", Color: ctpSapphire}, // 󱁼
	".doc":            {Icon: "\U000f1392", Color: ctpTeal},     // 󱎒
	".docx":           {Icon: "\U000f1392", Color: ctpTeal},     // 󱎒
	".dot":            {Icon: "\U000f1392", Color: ctpTeal},     // 󱎒
	".download":       {Icon: "\uf019", Color: ctpSapphire},     // 
	".drl":            {Icon: "\ue28c", Color: ctpMaroon},       // 
	".dropbox":        {Icon: "\ue707", Color: ctpBlue},         // 
	".ds_store":       {Icon: "\uf179", Color: ctpSubtext0},     // 
	".dump":           {Icon: "\uf1c0", Color: ctpRosewater},    // 
	".dwg":            {Icon: "\U000f0eeb", Color: ctpGreen},    // 󰻫
	".dxf":            {Icon: "\U000f0eeb", Color: ctpGreen},    // 󰻫
	".ebook":          {Icon: "\ue28b", Color: ctpPeach},        // 
	".ebuild":         {Icon: "\ue691", Color: ctpText},         // 
	".editorconfig":   {Icon: "\ue615", Color: ctpSubtext0},     // 
	".edn":            {Icon: "\ue768", Color: ctpGreen},        // 
	".eex":            {Icon: "\ue62d", Color: ctpYellow},       // 
	".ejs":            {Icon: "\ue618", Color: ctpPeach},        // 
	".el":             {Icon: "\ue6b0", Color: ctpText},         // 
	".elc":            {Icon: "\ue632", Color: ctpSubtext0},     // 
	".elf":            {Icon: "\ueae8", Color: ctpRed},          // 
	".elm":            {Icon: "\ue62c", Color: ctpTeal},         // 
	".eln":            {Icon: "\ue632", Color: ctpSubtext0},     // 
	".env":            {Icon: "\ueb52", Color: ctpYellow},       // 
	".eot":            {Icon: "\ue659", Color: ctpRed},          // 
	".epp":            {Icon: "\ue631", Color: ctpYellow},       // 
	".epub":           {Icon: "\ue28b", Color: ctpPeach},        // 
	".erb":            {Icon: "\U000f0d2d", Color: ctpPeach},    // 󰴭
	".erl":            {Icon: "\ue7b1", Color: ctpRed},          // 
	".ex":             {Icon: "\ue62d", Color: ctpMauve},        // 
	".exe":            {Icon: "\U000f05b3", Color: ctpRed},      // 󰖳
	".exs":            {Icon: "\ue653", Color: ctpMauve},        // 
	".f#":             {Icon: "\ue7a7", Color: ctpSapphire},     // 
	".f3d":            {Icon: "\U000f0eeb", Color: ctpGreen},    // 󰻫
	".f90":            {Icon: "\U000f121a", Color: ctpMauve},    // 󱈚
	".fbx":            {Icon: "\uea8c", Color: ctpSapphire},     // 
	".fcbak":          {Icon: "\uf336", Color: ctpSubtext0},     // 
	".fcmacro":        {Icon: "\uf336", Color: ctpRed},          // 
	".fcmat":          {Icon: "\uf336", Color: ctpRed},          // 
	".fcparam":        {Icon: "\uf336", Color: ctpRed},          // 
	".fcscript":       {Icon: "\uf336", Color: ctpRed},          // 
	".fcstd":          {Icon: "\uf336", Color: ctpRed},          // 
	".fcstd1":         {Icon: "\uf336", Color: ctpRed},          // 
	".fctb":           {Icon: "\uf336", Color: ctpRed},          // 
	".fctl":           {Icon: "\uf336", Color: ctpRed},          // 
	".fdmdownload":    {Icon: "\uf019", Color: ctpSapphire},     // 
	".fish":           {Icon: "\ue691", Color: ctpText},         // 
	".flac":           {Icon: "\U000f0223", Color: ctpPeach},    // 󰈣
	".flc":            {Icon: "\uf031", Color: ctpRosewater},    // 
	".flf":            {Icon: "\uf031", Color: ctpRosewater},    // 
	".flv":            {Icon: "\U000f0381", Color: ctpPeach},    // 󰎁
	".fnl":            {Icon: "\ue6af", Color: ctpYellow},       // 
	".fodg":           {Icon: "\uf379", Color: ctpPeach},        // 
	".fodp":           {Icon: "\uf37a", Color: ctpPeach},        // 
	".fods":           {Icon: "\uf378", Color: ctpGreen},        // 
	".fodt":           {Icon: "\uf37c", Color: ctpSapphire},     // 
	".font":           {Icon: "\ue659", Color: ctpRed},          // 
	".fs":             {Icon: "\ue7a7", Color: ctpBlue},         // 
	".fsi":            {Icon: "\ue7a7", Color: ctpBlue},         // 
	".fsscript":       {Icon: "\ue7a7", Color: ctpSapphire},     // 
	".fsx":            {Icon: "\ue7a7", Color: ctpBlue},         // 
	".gcode":          {Icon: "\U000f0af4", Color: ctpSubtext0}, // 󰫴
	".gd":             {Icon: "\ue65f", Color: ctpYellow},       // 
	".gdoc":           {Icon: "\uf1c2", Color: ctpGreen},        // 
	".gem":            {Icon: "\ue21e", Color: ctpRed},          // 
	".gemfile":        {Icon: "\ueb48", Color: ctpRed},          // 
	".gemspec":        {Icon: "\U000f0d2d", Color: ctpRed},      // 󰴭
	".gform":          {Icon: "\uf298", Color: ctpGreen},        // 
	".gif":            {Icon: "\U000f0d78", Color: ctpTeal},     // 󰵸
	".git":            {Icon: "\U000f02a2", Color: ctpPeach},    // 󰊢
	".glb":            {Icon: "\uf1b2", Color: ctpPeach},        // 
	".gnumakefile":    {Icon: "\U000f1064", Color: ctpText},     // 󱁤
	".go":             {Icon: "\U000f07d3", Color: ctpTeal},     // 󰟓
	".godot":          {Icon: "\ue65f", Color: ctpSapphire},     // 
	".gpr":            {Icon: "\U000f1077", Color: ctpTeal},     // 󱁷
	".gql":            {Icon: "\U000f0877", Color: ctpRed},      // 󰡷
	".gradle":         {Icon: "\ue775", Color: ctpTeal},         // 
	".graphql":        {Icon: "\U000f0877", Color: ctpRed},      // 󰡷
	".gresource":      {Icon: "\uf362", Color: ctpRosewater},    // 
	".groovy":         {Icon: "\ue775", Color: ctpTeal},         // 
	".gsheet":         {Icon: "\uf1c3", Color: ctpGreen},        // 
	".gslides":        {Icon: "\uf1c4", Color: ctpPeach},        // 
	".guardfile":      {Icon: "\ue21e", Color: ctpSubtext0},     // 
	".gv":             {Icon: "\U000f1049", Color: ctpTeal},     // 󱁉
	".gz":             {Icon: "\U000f05c4", Color: ctpText},     // 󰗄
	".h":              {Icon: "\U000f0af5", Color: ctpMauve},    // 󰫵
	".haml":           {Icon: "\U000f0174", Color: ctpText},     // 󰅴
	".hbs":            {Icon: "\ue7f7", Color: ctpPeach},        // 
	".hc":             {Icon: "\U000f00a2", Color: ctpPeach},    // 󰂢
	".heex":           {Icon: "\ue62d", Color: ctpRed},          // 
	".hex":            {Icon: "\U000f02d8", Color: ctpYellow},   // 󰋘
	".hh":             {Icon: "\U000f0672", Color: ctpTeal},     // 󰙲
	".hpp":            {Icon: "\U000f0672", Color: ctpTeal},     // 󰙲
	".hrl":            {Icon: "\ue7b1", Color: ctpRed},          // 
	".hs":             {Icon: "\U000f0c92", Color: ctpMauve},    // 󰲒
	".htm":            {Icon: "\U000f031d", Color: ctpPeach},    // 󰌝
	".html":           {Icon: "\U000f031d", Color: ctpPeach},    // 󰌝
	".huff":           {Icon: "\U000f0858", Color: ctpText},     // 󰡘
	".hurl":           {Icon: "\U000f0af5", Color: ctpGreen},    // 󰫵
	".hx":             {Icon: "\U000f0af5", Color: ctpText},     // 󰫵
	".hxx":            {Icon: "\U000f0672", Color: ctpTeal},     // 󰙲
	".ical":           {Icon: "\uf073", Color: ctpSapphire},     // 
	".icalendar":      {Icon: "\uf073", Color: ctpSapphire},     // 
	".ico":            {Icon: "\uf03e", Color: ctpGreen},        // 
	".ics":            {Icon: "\U000f01ee", Color: ctpSapphire}, // 󰇮
	".ifb":            {Icon: "\uf073", Color: ctpSapphire},     // 
	".ifc":            {Icon: "\U000f0eeb", Color: ctpGreen},    // 󰻫
	".ige":            {Icon: "\U000f0eeb", Color: ctpGreen},    // 󰻫
	".iges":           {Icon: "\U000f0eeb", Color: ctpGreen},    // 󰻫
	".igs":            {Icon: "\U000f0eeb", Color: ctpGreen},    // 󰻫
	".image":          {Icon: "\uf1c5", Color: ctpPeach},        // 
	".img":            {Icon: "\U000f021f", Color: ctpSapphire}, // 󰈟
	".iml":            {Icon: "\U000f022e", Color: ctpGreen},    // 󰈮
	".import":         {Icon: "\uf0c6", Color: ctpRosewater},    // 
	".info":           {Icon: "\uf129", Color: ctpRosewater},    // 
	".ini":            {Icon: "\U000f0bc2", Color: ctpTeal},     // 󰯂
	".ino":            {Icon: "\uf34b", Color: ctpTeal},         // 
	".ipynb":          {Icon: "\U000f082e", Color: ctpPeach},    // 󰠮
	".iso":            {Icon: "\uede9", Color: ctpSubtext0},     // 
	".ixx":            {Icon: "\U000f0672", Color: ctpTeal},     // 󰙲
	".j2c":            {Icon: "\uf1c5", Color: ctpSubtext0},     // 
	".j2k":            {Icon: "\uf1c5", Color: ctpSubtext0},     // 
	".jad":            {Icon: "\ue256", Color: ctpPeach},        // 
	".jar":            {Icon: "\U000f06ca", Color: ctpPeach},    // 󰛊
	".java":           {Icon: "\U000f0b37", Color: ctpPeach},    // 󰬷
	".jfi":            {Icon: "\uf1c5", Color: ctpSubtext0},     // 
	".jfif":           {Icon: "\U000f021f", Color: ctpSapphire}, // 󰈟
	".jif":            {Icon: "\uf1c5", Color: ctpSubtext0},     // 
	".jl":             {Icon: "\ue624", Color: ctpMauve},        // 
	".jmd":            {Icon: "\uf48a", Color: ctpSapphire},     // 
	".jp2":            {Icon: "\uf1c5", Color: ctpSubtext0},     // 
	".jpe":            {Icon: "\uf1c5", Color: ctpSubtext0},     // 
	".jpeg":           {Icon: "\U000f0225", Color: ctpPeach},    // 󰈥
	".jpg":            {Icon: "\U000f0225", Color: ctpPeach},    // 󰈥
	".jpx":            {Icon: "\uf1c5", Color: ctpSubtext0},     // 
	".js":             {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	".json":           {Icon: "\U000f0626", Color: ctpYellow},   // 󰘦
	".json5":          {Icon: "\U000f0626", Color: ctpPeach},    // 󰘦
	".jsonc":          {Icon: "\U000f0626", Color: ctpYellow},   // 󰘦
	".jsx":            {Icon: "\ue625", Color: ctpTeal},         // 
	".jwmrc":          {Icon: "\uf35b", Color: ctpSapphire},     // 
	".jxl":            {Icon: "\uf1c5", Color: ctpSubtext0},     // 
	".kbx":            {Icon: "\U000f0bc4", Color: ctpSubtext0}, // 󰯄
	".kdb":            {Icon: "\uf23e", Color: ctpGreen},        // 
	".kdbx":           {Icon: "\uf23e", Color: ctpGreen},        // 
	".kdenlive":       {Icon: "\uf33c", Color: ctpBlue},         // 
	".kdenlivetitle":  {Icon: "\uf33c", Color: ctpBlue},         // 
	".kicad_dru":      {Icon: "\uf34c", Color: ctpRosewater},    // 
	".kicad_mod":      {Icon: "\uf34c", Color: ctpRosewater},    // 
	".kicad_pcb":      {Icon: "\uf34c", Color: ctpRosewater},    // 
	".kicad_prl":      {Icon: "\uf34c", Color: ctpRosewater},    // 
	".kicad_pro":      {Icon: "\uf34c", Color: ctpRosewater},    // 
	".kicad_sch":      {Icon: "\uf34c", Color: ctpRosewater},    // 
	".kicad_sym":      {Icon: "\uf34c", Color: ctpRosewater},    // 
	".kicad_wks":      {Icon: "\uf34c", Color: ctpRosewater},    // 
	".ko":             {Icon: "\uf17c", Color: ctpRosewater},    // 
	".kpp":            {Icon: "\uf33d", Color: ctpMauve},        // 
	".kra":            {Icon: "\uf33d", Color: ctpMauve},        // 
	".krz":            {Icon: "\uf33d", Color: ctpMauve},        // 
	".ksh":            {Icon: "\ue691", Color: ctpText},         // 
	".kt":             {Icon: "\U000f1219", Color: ctpBlue},     // 󱈙
	".kts":            {Icon: "\U000f1219", Color: ctpBlue},     // 󱈙
	".latex":          {Icon: "\ue69b", Color: ctpGreen},        // 
	".lck":            {Icon: "\ue672", Color: ctpSubtext0},     // 
	".leex":           {Icon: "\ue62d", Color: ctpYellow},       // 
	".less":           {Icon: "\U000f031c", Color: ctpMauve},    // 󰌜
	".lff":            {Icon: "\uf031", Color: ctpRosewater},    // 
	".lhs":            {Icon: "\ue61f", Color: ctpMauve},        // 
	".license":        {Icon: "\U000f0124", Color: ctpPeach},    // 󰄤
	".liquid":         {Icon: "\ue670", Color: ctpGreen},        // 
	".localized":      {Icon: "\uf179", Color: ctpSubtext0},     // 
	".lock":           {Icon: "\uf023", Color: ctpPeach},        // 
	".log":            {Icon: "\uf0f6", Color: ctpPeach},        // 
	".lrc":            {Icon: "\U000f0af9", Color: ctpPeach},    // 󰫹
	".lua":            {Icon: "\U000f08b1", Color: ctpTeal},     // 󰢱
	".luac":           {Icon: "\ue620", Color: ctpSapphire},     // 
	".luau":           {Icon: "\U000f08b1", Color: ctpGreen},    // 󰢱
	".lz":             {Icon: "\uf410", Color: ctpPeach},        // 
	".lz4":            {Icon: "\uf410", Color: ctpPeach},        // 
	".lzh":            {Icon: "\uf410", Color: ctpPeach},        // 
	".lzma":           {Icon: "\uf410", Color: ctpPeach},        // 
	".lzo":            {Icon: "\uf410", Color: ctpPeach},        // 
	".m":              {Icon: "\U000f0fc8", Color: ctpPeach},    // 󰿈
	".m3u":            {Icon: "\U000f0cb8", Color: ctpPeach},    // 󰲸
	".m3u8":           {Icon: "\U000f0cb8", Color: ctpPeach},    // 󰲸
	".m4a":            {Icon: "\U000f0223", Color: ctpMauve},    // 󰈣
	".m4v":            {Icon: "\U000f022b", Color: ctpPeach},    // 󰈫
	".magnet":         {Icon: "\uf076", Color: ctpRed},          // 
	".makefile":       {Icon: "\U000f1064", Color: ctpText},     // 󱁤
	".markdown":       {Icon: "\U000f0354", Color: ctpText},     // 󰍔
	".material":       {Icon: "\U000f0509", Color: ctpRed},      // 󰔉
	".md":             {Icon: "\U000f0354", Color: ctpText},     // 󰍔
	".md5":            {Icon: "\U000f0565", Color: ctpSubtext0}, // 󰕥
	".mdx":            {Icon: "\uf48a", Color: ctpYellow},       // 
	".mint":           {Icon: "\ue7a4", Color: ctpGreen},        // 
	".mjs":            {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	".mk":             {Icon: "\U000f1064", Color: ctpText},     // 󱁤
	".mkd":            {Icon: "\U000f0354", Color: ctpText},     // 󰍔
	".mkv":            {Icon: "\U000f022b", Color: ctpGreen},    // 󰈫
	".ml":             {Icon: "\ue67a", Color: ctpPeach},        // 
	".mli":            {Icon: "\ue67a", Color: ctpPeach},        // 
	".mm":             {Icon: "\U000f0afb", Color: ctpTeal},     // 󰫻
	".mo":             {Icon: "\U000f0af4", Color: ctpBlue},     // 󰫴
	".mobi":           {Icon: "\ue28b", Color: ctpPeach},        // 
	".mojo":           {Icon: "\U000f0238", Color: ctpRed},      // 󰈸
	".mov":            {Icon: "\U000f022b", Color: ctpTeal},     // 󰈫
	".mp3":            {Icon: "\U000f0223", Color: ctpTeal},     // 󰈣
	".mp4":            {Icon: "\U000f022b", Color: ctpTeal},     // 󰈫
	".mpp":            {Icon: "\U000f0672", Color: ctpTeal},     // 󰙲
	".msf":            {Icon: "\U000f0b05", Color: ctpYellow},   // 󰬅
	".msi":            {Icon: "\uf2d0", Color: ctpPeach},        // 
	".mts":            {Icon: "\U000f06e6", Color: ctpTeal},     // 󰛦
	".mustache":       {Icon: "\U000f15de", Color: ctpTeal},     // 󱗞
	".nfo":            {Icon: "\uf129", Color: ctpRosewater},    // 
	".nim":            {Icon: "\ue677", Color: ctpYellow},       // 
	".nix":            {Icon: "\U000f1105", Color: ctpTeal},     // 󱄅
	".node":           {Icon: "\U000f0399", Color: ctpRed},      // 󰎙
	".npmignore":      {Icon: "\ue71e", Color: ctpRed},          // 
	".nswag":          {Icon: "\ue60b", Color: ctpGreen},        // 
	".nu":             {Icon: "\ue691", Color: ctpMauve},        // 
	".o":              {Icon: "\uea8c", Color: ctpSapphire},     // 
	".obj":            {Icon: "\U000f01a7", Color: ctpText},     // 󰆧
	".odin":           {Icon: "\U000f0b94", Color: ctpBlue},     // 󰮔
	".odf":            {Icon: "\uf37b", Color: ctpRed},          // 
	".odg":            {Icon: "\uf379", Color: ctpPeach},        // 
	".odp":            {Icon: "\uf37a", Color: ctpPeach},        // 
	".ods":            {Icon: "\uf378", Color: ctpGreen},        // 
	".odt":            {Icon: "\uf37c", Color: ctpSapphire},     // 
	".ogg":            {Icon: "\U000f0223", Color: ctpText},     // 󰈣
	".ogv":            {Icon: "\U000f0381", Color: ctpPeach},    // 󰎁
	".opus":           {Icon: "\U000f0223", Color: ctpPeach},    // 󰈣
	".org":            {Icon: "\ue633", Color: ctpTeal},         // 
	".otf":            {Icon: "\ue659", Color: ctpRed},          // 
	".out":            {Icon: "\ueae8", Color: ctpRed},          // 
	".part":           {Icon: "\uf43a", Color: ctpSubtext0},     // 
	".patch":          {Icon: "\U000f0993", Color: ctpRed},      // 󰦓
	".pck":            {Icon: "\uf487", Color: ctpSubtext0},     // 
	".pdf":            {Icon: "\U000f0226", Color: ctpRed},      // 󰈦
	".php":            {Icon: "\U000f031f", Color: ctpMauve},    // 󰌟
	".pl":             {Icon: "\ue67e", Color: ctpTeal},         // 
	".pls":            {Icon: "\U000f01bc", Color: ctpPeach},    // 󰆼
	".ply":            {Icon: "\U000f01a7", Color: ctpSubtext0}, // 󰆧
	".pm":             {Icon: "\ue67e", Color: ctpTeal},         // 
	".png":            {Icon: "\U000f0e2d", Color: ctpMauve},    // 󰸭
	".po":             {Icon: "\U000f05ca", Color: ctpTeal},     // 󰗊
	".pot":            {Icon: "\U000f05ca", Color: ctpTeal},     // 󰗊
	".pp":             {Icon: "\ue631", Color: ctpPeach},        // 
	".ppt":            {Icon: "\U000f1390", Color: ctpRed},      // 󱎐
	".pptx":           {Icon: "\U000f1390", Color: ctpRed},      // 󱎐
	".prisma":         {Icon: "\ue684", Color: ctpBlue},         // 
	".pro":            {Icon: "\U000f15ff", Color: ctpTeal},     // 󱗿
	".procfile":       {Icon: "\ue607", Color: ctpSubtext0},     // 
	".properties":     {Icon: "\U000f0b37", Color: ctpGreen},    // 󰬷
	".ps1":            {Icon: "\U000f0a0a", Color: ctpBlue},     // 󰨊
	".psb":            {Icon: "\U000f021f", Color: ctpSapphire}, // 󰈟
	".psd":            {Icon: "\ue7b8", Color: ctpSapphire},     // 
	".psd1":           {Icon: "\U000f0a0a", Color: ctpBlue},     // 󰨊
	".psm1":           {Icon: "\U000f0a0a", Color: ctpBlue},     // 󰨊
	".pub":            {Icon: "\U000f0306", Color: ctpSapphire}, // 󰌆
	".pxd":            {Icon: "\U000f0afd", Color: ctpYellow},   // 󰫽
	".pxi":            {Icon: "\U000f0afd", Color: ctpYellow},   // 󰫽
	".pxm":            {Icon: "\uf1c5", Color: ctpSubtext0},     // 
	".py":             {Icon: "\U000f0320", Color: ctpYellow},   // 󰌠
	".pyc":            {Icon: "\ue606", Color: ctpPeach},        // 
	".pyd":            {Icon: "\ue606", Color: ctpPeach},        // 
	".pyi":            {Icon: "\U000f0320", Color: ctpYellow},   // 󰌠
	".pyo":            {Icon: "\ue606", Color: ctpPeach},        // 
	".pyw":            {Icon: "\U000f0320", Color: ctpYellow},   // 󰌠
	".pyx":            {Icon: "\U000f0afd", Color: ctpYellow},   // 󰫽
	".qm":             {Icon: "\U000f05ca", Color: ctpSapphire}, // 󰗊
	".qml":            {Icon: "\U000f0afe", Color: ctpTeal},     // 󰫾
	".qrc":            {Icon: "\uf375", Color: ctpGreen},        // 
	".qss":            {Icon: "\uf375", Color: ctpGreen},        // 
	".query":          {Icon: "\ue21c", Color: ctpGreen},        // 
	".r":              {Icon: "\U000f07d4", Color: ctpBlue},     // 󰟔
	".rake":           {Icon: "\U000f0d2d", Color: ctpRed},      // 󰴭
	".rakefile":       {Icon: "\ue21e", Color: ctpRed},          // 
	".rar":            {Icon: "\U000f05c4", Color: ctpGreen},    // 󰗄
	".razor":          {Icon: "\U000f0214", Color: ctpText},     // 󰈔
	".rb":             {Icon: "\U000f0d2d", Color: ctpRed},      // 󰴭
	".rdata":          {Icon: "\uf25d", Color: ctpSapphire},     // 
	".rdb":            {Icon: "\ue76d", Color: ctpRed},          // 
	".rdoc":           {Icon: "\uf48a", Color: ctpSapphire},     // 
	".rds":            {Icon: "\uf25d", Color: ctpSapphire},     // 
	".readme":         {Icon: "\uf05a", Color: ctpSapphire},     // 
	".res":            {Icon: "\U000f0aff", Color: ctpTeal},     // 󰫿
	".resi":           {Icon: "\U000f0aff", Color: ctpTeal},     // 󰫿
	".rlib":           {Icon: "\ue7a8", Color: ctpPeach},        // 
	".rmd":            {Icon: "\U000f0354", Color: ctpTeal},     // 󰍔
	".rpm":            {Icon: "\U000f05c4", Color: ctpRed},      // 󰗄
	".rproj":          {Icon: "\U000f05c6", Color: ctpGreen},    // 󰗆
	".rs":             {Icon: "\U000f1617", Color: ctpPeach},    // 󱘗
	".rspec":          {Icon: "\ue21e", Color: ctpRed},          // 
	".rspec_parallel": {Icon: "\ue21e", Color: ctpRed},          // 
	".rspec_status":   {Icon: "\ue21e", Color: ctpRed},          // 
	".rss":            {Icon: "\U000f05c0", Color: ctpPeach},    // 󰗀
	".rtf":            {Icon: "\U000f069e", Color: ctpTeal},     // 󰚞
	".ru":             {Icon: "\U000f0d2d", Color: ctpRed},      // 󰴭
	".rubydoc":        {Icon: "\ue73b", Color: ctpRed},          // 
	".s":              {Icon: "\ue637", Color: ctpMauve},        // 
	".sass":           {Icon: "\U000f07ec", Color: ctpRed},      // 󰟬
	".sbt":            {Icon: "\ue737", Color: ctpPeach},        // 
	".sc":             {Icon: "\ue737", Color: ctpRed},          // 
	".scad":           {Icon: "\uf34e", Color: ctpYellow},       // 
	".scala":          {Icon: "\ue737", Color: ctpRed},          // 
	".scm":            {Icon: "\ue6b1", Color: ctpText},         // 
	".scss":           {Icon: "\U000f07ec", Color: ctpRed},      // 󰟬
	".sh":             {Icon: "\ue691", Color: ctpText},         // 
	".sha1":           {Icon: "\U000f0565", Color: ctpSubtext0}, // 󰕥
	".sha224":         {Icon: "\U000f0565", Color: ctpSubtext0}, // 󰕥
	".sha256":         {Icon: "\U000f0565", Color: ctpSubtext0}, // 󰕥
	".sha384":         {Icon: "\U000f0565", Color: ctpSubtext0}, // 󰕥
	".sha512":         {Icon: "\U000f0565", Color: ctpSubtext0}, // 󰕥
	".shell":          {Icon: "\ue795", Color: ctpGreen},        // 
	".sig":            {Icon: "\u03bb", Color: ctpPeach},        // λ
	".signature":      {Icon: "\u03bb", Color: ctpPeach},        // λ
	".skp":            {Icon: "\uea8c", Color: ctpSapphire},     // 
	".sldasm":         {Icon: "\U000f0eeb", Color: ctpGreen},    // 󰻫
	".sldprt":         {Icon: "\U000f0eeb", Color: ctpGreen},    // 󰻫
	".slim":           {Icon: "\ue692", Color: ctpPeach},        // 
	".sln":            {Icon: "\U000f0610", Color: ctpBlue},     // 󰘐
	".slvs":           {Icon: "\U000f0eeb", Color: ctpGreen},    // 󰻫
	".sml":            {Icon: "\U000f0627", Color: ctpPeach},    // 󰘧
	".so":             {Icon: "\U000f107c", Color: ctpSapphire}, // 󱁼
	".sol":            {Icon: "\ue656", Color: ctpTeal},         // 
	".spec.js":        {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	".spec.jsx":       {Icon: "\ue625", Color: ctpTeal},         // 
	".spec.ts":        {Icon: "\U000f06e6", Color: ctpTeal},     // 󰛦
	".spec.tsx":       {Icon: "\ue7ba", Color: ctpBlue},         // 
	".sql":            {Icon: "\U000f01bc", Color: ctpYellow},   // 󰆼
	".sqlite":         {Icon: "\uf1c0", Color: ctpGreen},        // 
	".sqlite3":        {Icon: "\uf1c0", Color: ctpGreen},        // 
	".srt":            {Icon: "\U000f0a16", Color: ctpYellow},   // 󰨖
	".ssa":            {Icon: "\U000f0a16", Color: ctpPeach},    // 󰨖
	".ste":            {Icon: "\U000f0eeb", Color: ctpGreen},    // 󰻫
	".step":           {Icon: "\U000f0eeb", Color: ctpGreen},    // 󰻫
	".stl":            {Icon: "\uea8c", Color: ctpSapphire},     // 
	".stp":            {Icon: "\U000f0b00", Color: ctpYellow},   // 󰬀
	".strings":        {Icon: "\U000f05ca", Color: ctpSapphire}, // 󰗊
	".sty":            {Icon: "\ue69b", Color: ctpGreen},        // 
	".styl":           {Icon: "\U000f0d12", Color: ctpText},     // 󰴒
	".stylus":         {Icon: "\U000f0d12", Color: ctpText},     // 󰴒
	".sub":            {Icon: "\U000f06a9", Color: ctpText},     // 󰚩
	".sublime":        {Icon: "\ue7aa", Color: ctpPeach},        // 
	".suo":            {Icon: "\U000f0610", Color: ctpRed},      // 󰘐
	".sv":             {Icon: "\U000f035b", Color: ctpGreen},    // 󰍛
	".svelte":         {Icon: "\ue697", Color: ctpPeach},        // 
	".svg":            {Icon: "\U000f0721", Color: ctpYellow},   // 󰜡
	".svh":            {Icon: "\U000f035b", Color: ctpGreen},    // 󰍛
	".swift":          {Icon: "\U000f06e5", Color: ctpPeach},    // 󰛥
	".t":              {Icon: "\U000f1a7c", Color: ctpTeal},     // 󱩼
	".tar":            {Icon: "\U000f05c4", Color: ctpTeal},     // 󰗄
	".taz":            {Icon: "\uf410", Color: ctpPeach},        // 
	".tbc":            {Icon: "\U000f06d3", Color: ctpSapphire}, // 󰛓
	".tbz":            {Icon: "\uf410", Color: ctpPeach},        // 
	".tbz2":           {Icon: "\uf410", Color: ctpPeach},        // 
	".tcl":            {Icon: "\U000f06d3", Color: ctpRed},      // 󰛓
	".templ":          {Icon: "\U000f0b01", Color: ctpTeal},     // 󰬁
	".terminal":       {Icon: "\uf489", Color: ctpGreen},        // 
	".test.js":        {Icon: "\U000f031e", Color: ctpYellow},   // 󰌞
	".test.jsx":       {Icon: "\ue625", Color: ctpTeal},         // 
	".test.ts":        {Icon: "\U000f06e6", Color: ctpTeal},     // 󰛦
	".test.tsx":       {Icon: "\ue7ba", Color: ctpBlue},         // 
	".tex":            {Icon: "\ue69b", Color: ctpGreen},        // 
	".tf":             {Icon: "\U000f1062", Color: ctpBlue},     // 󱁢
	".tfvars":         {Icon: "\U000f1062", Color: ctpTeal},     // 󱁢
	".tgz":            {Icon: "\U000f05c4", Color: ctpText},     // 󰗄
	".tiff":           {Icon: "\U000f021f", Color: ctpYellow},   // 󰈟
	".tlz":            {Icon: "\uf410", Color: ctpPeach},        // 
	".tmux":           {Icon: "\uebc8", Color: ctpGreen},        // 
	".toml":           {Icon: "\ue6b2", Color: ctpPeach},        // 
	".torrent":        {Icon: "\ue275", Color: ctpSapphire},     // 
	".tres":           {Icon: "\ue65f", Color: ctpGreen},        // 
	".ts":             {Icon: "\U000f06e6", Color: ctpTeal},     // 󰛦
	".tscn":           {Icon: "\ue65f", Color: ctpGreen},        // 
	".tsconfig":       {Icon: "\ue772", Color: ctpPeach},        // 
	".tsv":            {Icon: "\ue64a", Color: ctpBlue},         // 
	".tsx":            {Icon: "\ue7ba", Color: ctpBlue},         // 
	".ttf":            {Icon: "\ue659", Color: ctpRed},          // 
	".twig":           {Icon: "\ue61c", Color: ctpGreen},        // 
	".txt":            {Icon: "\U000f09a8", Color: ctpText},     // 󰦨
	".txz":            {Icon: "\U000f05c4", Color: ctpMauve},    // 󰗄
	".typ":            {Icon: "\U000f0b1b", Color: ctpTeal},     // 󰬛
	".typoscript":     {Icon: "\ue772", Color: ctpPeach},        // 
	".tz":             {Icon: "\uf410", Color: ctpPeach},        // 
	".tzo":            {Icon: "\uf410", Color: ctpPeach},        // 
	".ui":             {Icon: "\U000f05c0", Color: ctpPeach},    // 󰗀
	".v":              {Icon: "\ue6ac", Color: ctpBlue},         // 
	".vala":           {Icon: "\U000f0b1d", Color: ctpMauve},    // 󰬝
	".vh":             {Icon: "\U000f035b", Color: ctpGreen},    // 󰍛
	".vhd":            {Icon: "\U000f035b", Color: ctpGreen},    // 󰍛
	".vhdl":           {Icon: "\U000f035b", Color: ctpGreen},    // 󰍛
	".video":          {Icon: "\uf03d", Color: ctpSubtext0},     // 
	".vi":             {Icon: "\ue81e", Color: ctpPeach},        // 
	".vim":            {Icon: "\ue7c5", Color: ctpGreen},        // 
	".vsh":            {Icon: "\ue6ac", Color: ctpBlue},         // 
	".vsix":           {Icon: "\U000f0a1e", Color: ctpSapphire}, // 󰨞
	".vue":            {Icon: "\U000f0844", Color: ctpGreen},    // 󰡄
	".war":            {Icon: "\ue256", Color: ctpRed},          // 
	".wasm":           {Icon: "\ue6a1", Color: ctpBlue},         // 
	".wav":            {Icon: "\U000f0223", Color: ctpGreen},    // 󰈣
	".webm":           {Icon: "\U000f022b", Color: ctpText},     // 󰈫
	".webmanifest":    {Icon: "\U000f0626", Color: ctpYellow},   // 󰘦
	".webp":           {Icon: "\U000f021f", Color: ctpBlue},     // 󰈟
	".webpack":        {Icon: "\U000f072b", Color: ctpSapphire}, // 󰜫
	".windows":        {Icon: "\uf17a", Color: ctpSapphire},     // 
	".wma":            {Icon: "\U000f0223", Color: ctpBlue},     // 󰈣
	".woff":           {Icon: "\ue659", Color: ctpRed},          // 
	".woff2":          {Icon: "\ue659", Color: ctpRed},          // 
	".wrl":            {Icon: "\U000f0b03", Color: ctpBlue},     // 󰬃
	".wrz":            {Icon: "\U000f01a7", Color: ctpSubtext0}, // 󰆧
	".wv":             {Icon: "\uf001", Color: ctpSapphire},     // 
	".wvc":            {Icon: "\uf001", Color: ctpSapphire},     // 
	".x":              {Icon: "\U000f0aff", Color: ctpTeal},     // 󰫿
	".xaml":           {Icon: "\U000f0673", Color: ctpSapphire}, // 󰙳
	".xcf":            {Icon: "\uf338", Color: ctpSubtext0},     // 
	".xcplayground":   {Icon: "\ue755", Color: ctpPeach},        // 
	".xcstrings":      {Icon: "\U000f05ca", Color: ctpSapphire}, // 󰗊
	".xhtml":          {Icon: "\U000f031d", Color: ctpPeach},    // 󰌝
	".xls":            {Icon: "\U000f138f", Color: ctpGreen},    // 󱎏
	".xlsx":           {Icon: "\U000f138f", Color: ctpGreen},    // 󱎏
	".xm":             {Icon: "\ue691", Color: ctpSapphire},     // 
	".xml":            {Icon: "\U000f05c0", Color: ctpPeach},    // 󰗀
	".xpi":            {Icon: "\ueae6", Color: ctpSubtext0},     // 
	".xul":            {Icon: "\U000f05c0", Color: ctpPeach},    // 󰗀
	".xz":             {Icon: "\U000f05c4", Color: ctpGreen},    // 󰗄
	".yaml":           {Icon: "\ue6a8", Color: ctpMauve},        // 
	".yml":            {Icon: "\ue6a8", Color: ctpMauve},        // 
	".zig":            {Icon: "\ue6a9", Color: ctpPeach},        // 
	".zip":            {Icon: "\U000f05c4", Color: ctpTeal},     // 󰗄
	".zsh":            {Icon: "\ue691", Color: ctpGreen},        // 
	".zsh-theme":      {Icon: "\ue691", Color: ctpGreen},        // 
	".zshrc":          {Icon: "\ue795", Color: ctpGreen},        // 
	".zst":            {Icon: "\U000f05c4", Color: ctpYellow},   // 󰗄
}

func patchFileIconsForNerdFontsV2() {
	extIconMap[".cs"] = IconProperties{Icon: "\uf81a", Color: ctpRosewater}     // 
	extIconMap[".csproj"] = IconProperties{Icon: "\uf81a", Color: ctpRed}       // 
	extIconMap[".csx"] = IconProperties{Icon: "\uf81a", Color: ctpSapphire}     // 
	extIconMap[".license"] = IconProperties{Icon: "\uf718", Color: ctpSubtext0} // 
	extIconMap[".node"] = IconProperties{Icon: "\uf898", Color: ctpRed}         // 
	extIconMap[".rtf"] = IconProperties{Icon: "\uf718", Color: ctpSubtext0}     // 
	extIconMap[".vue"] = IconProperties{Icon: "\ufd42", Color: ctpGreen}        // ﵂
}

func IconForFile(name string, isSubmodule bool, isLinkedWorktree bool, isDirectory bool, customIconsConfig *config.CustomIconsConfig) IconProperties {
	base := filepath.Base(name)
	if icon, ok := customIconsConfig.Filenames[base]; ok {
		return IconProperties{Color: icon.Color, Icon: icon.Icon}
	}
	if icon, ok := nameIconMap[base]; ok {
		return icon
	}

	ext := strings.ToLower(filepath.Ext(name))
	if icon, ok := customIconsConfig.Extensions[ext]; ok {
		return IconProperties{Color: icon.Color, Icon: icon.Icon}
	}
	if icon, ok := extIconMap[ext]; ok {
		return icon
	}

	if isSubmodule {
		return DEFAULT_SUBMODULE_ICON
	} else if isLinkedWorktree {
		return IconProperties{LINKED_WORKTREE_ICON, "#4E4E4E"}
	} else if isDirectory {
		return DEFAULT_DIRECTORY_ICON
	}
	return DEFAULT_FILE_ICON
}
