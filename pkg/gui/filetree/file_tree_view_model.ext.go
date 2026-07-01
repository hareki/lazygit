package filetree

// Try to select the given path if present, expanding any collapsed parent
// directories along the way so the file is reachable. filepath must be a
// repo-root-relative path using forward slashes. Does nothing and returns false
// if the path isn't in the tree.
func (self *FileTreeViewModel) SelectPath(filepath string, showRootItem bool) bool {
	internalPath := InternalTreePathForFilePath(filepath, showRootItem)
	self.ExpandToPath(internalPath)
	index, found := self.GetIndexForPath(internalPath)
	if found {
		self.SetSelection(index)
	}
	return found
}
