package model

import (
	"gomodoro/pkg/manager"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var textStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	Padding(1).
	Width(80).Align(lipgloss.Center)

const (
	padding  = 2
	maxWidth = 80
)

type TeaModel struct {
	Manager  manager.Manager
	Progress progress.Model
}

func (m TeaModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	m.Progress.ShowPercentage = false
	return m.Manager.Init()
}

func (m TeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Progress.Width = msg.Width - padding*2 - 4
		if m.Progress.Width > maxWidth {
			m.Progress.Width = maxWidth
		}
		return m, nil

	case progress.FrameMsg:
		progressModel, cmd := m.Progress.Update(msg)
		m.Progress = progressModel.(progress.Model)
		return m, cmd

	case manager.TickMsg:
		var cmd tea.Cmd
		m.Manager, cmd = m.Manager.UpdateState()
		snd := m.Progress.IncrPercent(1 / float64(m.Manager.TotalDuration()))

		return m, tea.Batch(snd, cmd)

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter", " ":
			m.Manager.Timer.Toggle()
		}
	}

	return m, nil
}

func (m TeaModel) View() string {
	state := ""
	if m.Manager.PomodoroState == 0 {
		state = "WORK"
	} else {
		state = "BREAK"
	}
	// The header
	s := "\n\n"
	if !m.Manager.Timer.Running {
		s += "STOPPED\n"
	}

	s += m.Progress.View() + "\n"

	s += textStyle.Render(
		state + "\n" + m.Manager.GetHumanTime(),
	)
	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
