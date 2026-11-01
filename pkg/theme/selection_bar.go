package theme

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	"github.com/jesseduffield/lazygit/pkg/gocui"
	"github.com/samber/lo"
)

var (
	// GocuiSelectionBarBgColor is the color of the bar that marks the selection in the
	// main views (see gocui.View.SelectedLineColorWidth)
	GocuiSelectionBarBgColor gocui.Attribute
	// GocuiSelectionBarFgColor is applied to the text under that bar
	GocuiSelectionBarFgColor gocui.Attribute
)

func updateSelectionBarColors(themeConfig config.ThemeConfig) {
	GocuiSelectionBarBgColor = GetGocuiStyle(
		lo.CoalesceSliceOrEmpty(themeConfig.SelectionBarBgColor, themeConfig.SelectedLineBgColor))
	GocuiSelectionBarFgColor = GetGocuiStyle(
		lo.CoalesceSliceOrEmpty(themeConfig.SelectionBarFgColor, themeConfig.SelectedLineFgColor))
}
