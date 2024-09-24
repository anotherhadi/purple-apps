package purple

import (
	"github.com/charmbracelet/lipgloss"
)

type colors struct {
	Accent    lipgloss.Color
	Gray      lipgloss.Color
	LightGray lipgloss.Color
	Muted     lipgloss.Color
}

var Colors = colors{
	Accent:    lipgloss.Color(Config.Global.Colors.Accent),
	Gray:      lipgloss.Color(Config.Global.Colors.Gray),
	LightGray: lipgloss.Color(Config.Global.Colors.LightGray),
	Muted:     lipgloss.Color(Config.Global.Colors.Muted),
}

// func relativeLuminance(r, g, b uint32) float64 {
// 	rf := float64(r) / 255.0
// 	gf := float64(g) / 255.0
// 	bf := float64(b) / 255.0
// 	rf = adjustColor(rf)
// 	gf = adjustColor(gf)
// 	bf = adjustColor(bf)
// 	return 0.2126*rf + 0.7151*gf + 0.0721*bf
// }
//
// func adjustColor(c float64) float64 {
// 	return math.Pow(c, 2.2)
// }

func GetFgColor(backgroundColor lipgloss.Color) lipgloss.Color {
	r, g, b, _ := backgroundColor.RGBA()
	// luminance := relativeLuminance(r, g, b)
	mean := int((r + g + b) / 3)
	if mean < 32768 {
		return lipgloss.Color("#FFFFFF")
	} else {
		return lipgloss.Color("#000000")
	}

	// if luminance >= 0.5 {
	// 	return lipgloss.Color("#000000")
	// } else {
	// 	return lipgloss.Color("#FFFFFF")
	// }
}
