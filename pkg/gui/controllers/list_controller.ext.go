package controllers

func (self *ListController) wrappedDelta(change int) int {
	list := self.context.GetList()
	if list.IsSelectingRange() {
		return change
	}
	length := list.Len()
	if length <= 0 {
		return change
	}
	currentIdx := list.GetSelectedLineIdx()
	if change < 0 && currentIdx == 0 {
		return length - 1
	}
	if change > 0 && currentIdx == length-1 {
		return -(length - 1)
	}
	return change
}
