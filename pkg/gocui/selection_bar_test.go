package gocui

import (
	"testing"

	"github.com/gdamore/tcell/v3"
	"github.com/stretchr/testify/assert"
)

// A strong bar color needs a text color of its own to keep the line numbers under
// it readable. Nothing past the bar may pick that up: those columns still show the
// diff's own colors.
func TestSelectionBarColorsStayOnTheBar(t *testing.T) {
	WithSimulationScreen(t, 14, 6)

	v := NewView("name", 0, 0, 11, 5, OutputTrue) // InnerWidth 10
	v.Highlight = true
	v.SelBgColor = NewRGBColor(0x89, 0xb4, 0xfa)
	v.SelTextColor = NewRGBColor(0x1e, 0x1e, 0x2e)
	v.SelectedLineColorWidth = 2

	// Red text on a dark red line, across the whole width of the view.
	v.writeString("\x1b[38;2;243;139;168;48;2;68;50;68m0123456789\n")
	v.draw(true)

	barBg := tcell.NewRGBColor(0x89, 0xb4, 0xfa)
	barFg := tcell.NewRGBColor(0x1e, 0x1e, 0x2e)
	lineBg := tcell.NewRGBColor(68, 50, 68)
	lineFg := tcell.NewRGBColor(243, 139, 168)

	// The frame puts view column x at screen column x+1.
	colors := func(x int) (tcell.Color, tcell.Color) {
		_, style, _ := Screen.Get(x+1, 1)
		return style.GetForeground(), style.GetBackground()
	}

	for _, x := range []int{0, 1, 8, 9} {
		fg, bg := colors(x)
		assert.Equal(t, barFg, fg, "text color in the bar at column %d", x)
		assert.Equal(t, barBg, bg, "background in the bar at column %d", x)
	}
	for x := 2; x < 8; x++ {
		fg, bg := colors(x)
		assert.Equal(t, lineFg, fg, "text color past the bar at column %d", x)
		assert.Equal(t, lineBg, bg, "background past the bar at column %d", x)
	}
}
