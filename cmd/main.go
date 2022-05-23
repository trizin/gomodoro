package main

import (
	"fmt"
	"gomodoro/pkg/manager"
	"gomodoro/pkg/model"
	"gomodoro/pkg/timer"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

var breakDuration = 5 * 60
var workDuration = 25 * 60

var M manager.Manager

func main() {
	M = manager.Manager{
		BreakDuration: breakDuration,
		WorkDuration:  workDuration,
		PomodoroState: manager.Work,
		Timer: timer.Timer{
			Running: true,
			Count:   0,
		},
	}

	p := tea.NewProgram(
		&model.TeaModel{
			Manager:  M,
			Progress: progress.New(progress.WithDefaultGradient()),
		},
	)
	if err := p.Start(); err != nil {
		fmt.Printf("An error occured while starting the app %v\n", err)
		panic("i'm pancikingnkgingng")
	}
}
