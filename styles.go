package purple

import "github.com/charmbracelet/lipgloss"

var (
	HeaderStyle     = lipgloss.NewStyle().Foreground(accent).Bold(true).Align(lipgloss.Center)
	CellStyle       = lipgloss.NewStyle().Padding(0, 1)
	CurrentDayStyle = lipgloss.NewStyle().Foreground(accent)
	OddRowStyle     = CellStyle.Foreground(gray).Align(lipgloss.Center)
	EvenRowStyle    = CellStyle.Foreground(lightgray).Align(lipgloss.Center)
	BorderStyle     = lipgloss.NewStyle().Foreground(accent)
	CellHoverStyle  = lipgloss.NewStyle().Foreground(accent).Bold(true).Underline(true)
)
