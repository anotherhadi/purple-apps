package purple

import "github.com/charmbracelet/lipgloss"

var (
	HeaderStyle     = lipgloss.NewStyle().Foreground(Accent).Bold(true).Align(lipgloss.Center)
	CellStyle       = lipgloss.NewStyle().Padding(0, 1)
	CurrentDayStyle = lipgloss.NewStyle().Foreground(Accent)
	OddRowStyle     = CellStyle.Foreground(Gray).Align(lipgloss.Center)
	EvenRowStyle    = CellStyle.Foreground(Lightgray).Align(lipgloss.Center)
	BorderStyle     = lipgloss.NewStyle().Foreground(Accent)
	CellHoverStyle  = lipgloss.NewStyle().Foreground(Accent).Bold(true).Underline(true)
)
