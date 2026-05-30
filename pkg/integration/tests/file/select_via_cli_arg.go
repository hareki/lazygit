package file

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var SelectViaCliArg = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Select a file in the Files panel on startup using the --file CLI arg",
	ExtraCmdArgs: []string{"--file=dir/file-two"},
	Skip:         false,
	SetupConfig: func(config *config.AppConfig) {
	},
	SetupRepo: func(shell *Shell) {
		shell.CreateFile("file-one", "original content\n")
		shell.CreateDir("dir")
		shell.CreateFile("dir/file-two", "original content\n")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		// The file passed via --file should be selected on startup, even though
		// it lives in a subdirectory and isn't the first item in the panel.
		t.Views().Files().
			IsFocused().
			SelectedLine(Contains("file-two"))
	},
})
