package helpers

import "path/filepath"

// selectStartupFileIfNeeded selects the file requested via the --file CLI flag
// in the Files panel, once that panel has been populated. It runs at most once:
// the stored path is cleared on the first attempt so that the periodic
// background files refresh doesn't keep snapping the cursor back.
//
// The flag value may be absolute or relative to the current working directory;
// we normalize it to a repo-root-relative, forward-slash path (the form
// models.File.Path uses) before matching. If the path is invalid or the file
// isn't in the panel, this is a no-op.
func (self *RefreshHelper) selectStartupFileIfNeeded() {
	repoState := self.c.State().GetRepoState()
	startupFile := repoState.GetStartupFile()
	if startupFile == "" {
		return
	}
	repoState.SetStartupFile("")

	absPath, err := filepath.Abs(startupFile)
	if err != nil {
		return
	}

	relPath, err := filepath.Rel(self.c.Git().RepoPaths.RepoPath(), absPath)
	if err != nil {
		return
	}

	self.c.Contexts().Files.SelectPath(
		filepath.ToSlash(relPath),
		self.c.UserConfig().Gui.ShowRootItemInFileTree,
	)
}
