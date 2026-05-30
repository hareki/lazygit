package gui

func (self *GuiRepoState) GetStartupFile() string {
	return self.StartupFile
}

func (self *GuiRepoState) SetStartupFile(value string) {
	self.StartupFile = value
}
