package diff

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var DiffRendererWidthFollowsMainView = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "A diff renderer is re-run with the new width when the main view is resized",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig: func(cfg *config.AppConfig) {
		cfg.GetUserConfig().Git.DiffRenderers = []config.DiffRendererConfig{
			{Command: "echo width=$LAZYGIT_COLUMNS && cat"},
		}
	},
	SetupRepo: func(shell *Shell) {
		shell.CreateFileAndAdd("file", "content\n")
		shell.Commit("initial")
		shell.UpdateFile("file", "changed\n")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		// The main view takes two thirds of the 150 column screen.
		t.Views().Files().
			Focus().
			Lines(Contains("file").IsSelected())
		t.Views().Main().
			Content(Contains("width=98"))

		// Half screen mode enlarges the focused side panel vertically only, so
		// the main view keeps its width...
		t.Views().Files().
			Press(keys.Universal.NextScreenMode)
		t.Views().Main().
			Content(Contains("width=98"))

		// ...until the main view is focused, which hides the side panels.
		t.Views().Files().
			Press(keys.Universal.FocusMainView)
		t.Views().Main().
			IsFocused().
			Content(Contains("width=148"))

		// Leaving the main view brings the side panels back.
		t.Views().Main().
			PressEscape()
		t.Views().Files().
			IsFocused()
		t.Views().Main().
			Content(Contains("width=98"))
	},
})
