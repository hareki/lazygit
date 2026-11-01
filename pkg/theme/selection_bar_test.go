package theme

import (
	"testing"

	"github.com/jesseduffield/lazygit/pkg/config"
	"github.com/jesseduffield/lazygit/pkg/gocui"
	"github.com/stretchr/testify/assert"
)

func TestSelectionBarColorsFallBackToTheSelectedLine(t *testing.T) {
	UpdateTheme(config.ThemeConfig{
		SelectedLineBgColor: []string{"blue"},
		SelectedLineFgColor: []string{"bold"},
	})

	assert.Equal(t, gocui.ColorBlue, GocuiSelectionBarBgColor)
	assert.Equal(t, gocui.AttrBold, GocuiSelectionBarFgColor)
}

func TestSelectionBarColorsOverrideTheSelectedLineOnTheBarOnly(t *testing.T) {
	UpdateTheme(config.ThemeConfig{
		SelectedLineBgColor: []string{"blue"},
		SelectedLineFgColor: []string{"bold"},
		SelectionBarBgColor: []string{"yellow"},
		SelectionBarFgColor: []string{"black", "underline"},
	})

	assert.Equal(t, gocui.ColorYellow, GocuiSelectionBarBgColor)
	assert.Equal(t, gocui.ColorBlack|gocui.AttrUnderline, GocuiSelectionBarFgColor)

	assert.Equal(t, gocui.ColorBlue, GocuiSelectedLineBgColor)
	assert.Equal(t, gocui.AttrBold, GocuiSelectedLineFgColor)
}
