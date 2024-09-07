package purple

import (
	"math"

	"github.com/charmbracelet/lipgloss"
)

type colors struct {
	Accent    lipgloss.Color
	Gray      lipgloss.Color
	LightGray lipgloss.Color
	Muted     lipgloss.Color
}

var (
	Colors = colors{
		Accent:    lipgloss.Color(Config.Global.Colors.Accent),
		Gray:      lipgloss.Color(Config.Global.Colors.Gray),
		LightGray: lipgloss.Color(Config.Global.Colors.LightGray),
		Muted:     lipgloss.Color(Config.Global.Colors.Muted),
	}
)

func relativeLuminance(r, g, b uint32) float64 {
	rf := float64(r) / 255.0
	gf := float64(g) / 255.0
	bf := float64(b) / 255.0
	rf = adjustColor(rf)
	gf = adjustColor(gf)
	bf = adjustColor(bf)
	return 0.2126*rf + 0.7152*gf + 0.0722*bf
}

func adjustColor(c float64) float64 {
	if c <= 0.03928 {
		return c / 12.92
	}
	return math.Pow((c+0.055)/1.055, 2.4)
}

func GetFgColor(backgroundColor lipgloss.Color) lipgloss.Color {
	r, g, b, _ := backgroundColor.RGBA()
	luminance := relativeLuminance(r, g, b)

	if luminance > 0.5 {
		return lipgloss.Color("#000000")
	} else {
		return lipgloss.Color("#FFFFFF")
	}
}
