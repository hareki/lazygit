package helpers

import "path/filepath"

// selectStartupFileIfNeeded selects the file requested via the --file CLI flag
// in the Files panel, once that panel has been populated. It runs at most once:
// the stored path is cleared on the first attempt so that the periodic
// background files refresh doesn't keep snapping the cursor back. Returns true
// if a file was actually selected, so the caller can scroll it into view.
//
// The flag value may be absolute or relative to the current working directory;
// we normalize it to a repo-root-relative, forward-slash path (the form
// models.File.Path uses) before matching. If the path is invalid or the file
// isn't in the panel, this is a no-op.
func (self *RefreshHelper) selectStartupFileIfNeeded() bool {
	repoState := self.c.State().GetRepoState()
	startupFile := repoState.GetStartupFile()
	if startupFile == "" {
		return false
	}
	repoState.SetStartupFile("")

	absPath, err := filepath.Abs(startupFile)
	if err != nil {
		return false
	}

	relPath, err := filepath.Rel(self.c.Git().RepoPaths.RepoPath(), absPath)
	if err != nil {
		return false
	}

	return self.c.Contexts().Files.SelectPath(
		filepath.ToSlash(relPath),
		self.c.UserConfig().Gui.ShowRootItemInFileTree,
	)
}

// scrollStartupFileIntoView scrolls the Files panel so the file selected via
// selectStartupFileIfNeeded is visible. The refresh's own postRefreshUpdate
// focuses the selected line without scrolling (ScrollSelectionIntoView is false
// so that a background refresh never yanks the viewport), which leaves the
// startup file selected but possibly off-screen. This is the one focus that
// scrolls it into view.
//
// It's queued on the UI thread so it runs after refreshView has rendered the
// Files view with the new selection; only then can FocusPoint compute the
// origin. Being queued last also makes its origin update the one that wins.
func (self *RefreshHelper) scrollStartupFileIntoView() {
	self.c.OnUIThread(func() error {
		self.c.Contexts().Files.FocusLine(true)
		return nil
	})
}
