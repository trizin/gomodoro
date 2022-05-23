package main

import (
	"fmt"
	"gomodoro/pkg/manager"
	"gomodoro/pkg/model"
	"gomodoro/pkg/progress"
	"gomodoro/pkg/timer"

	tea "github.com/charmbracelet/bubbletea"
)

var breakDuration = 5 * 60
var longBreakDuration = 30 * 60
var workDuration = 25 * 60

var M manager.Manager

func main() {
	M = manager.Manager{
		BreakDuration: breakDuration,
		WorkDuration:  workDuration,
		PomodoroState: manager.Work,
		Count:         0,
		Timer: timer.Timer{
			Running: true,
			Count:   0,
		},
		LongBreakDuration: longBreakDuration,
	}

	p := tea.NewProgram(
		&model.TeaModel{
			Manager: M,
			Progress: progress.New(progress.WithGradient(
				"#7D56F4",
				"#f02961",
			), progress.WithoutPercentage()),
		},
		tea.WithAltScreen(),
	)
	if err := p.Start(); err != nil {
		fmt.Printf("An error occured while starting the app %v\n", err)
		panic("i'm pancikingnkgingng")
	}
}
