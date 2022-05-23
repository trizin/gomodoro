package styles

import "github.com/charmbracelet/lipgloss"

var workColor = "#7D56F4"
var breakColor = "#f02961"

var WorkGradient = [2]string{"#7D56F4", "#6638f2"}
var BreakGradient = [2]string{"#f02961", "#ed0748"}

func GetTextStyle(pomodoroState int) lipgloss.Style {
	textStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Padding(1).
		Width(60).Align(lipgloss.Center)

	if pomodoroState == 0 {
		textStyle.Background(lipgloss.Color(workColor))
	} else {
		textStyle.Background(lipgloss.Color(breakColor))
	}

	return textStyle
}

var HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Margin(1, 0).Align(lipgloss.Center)
