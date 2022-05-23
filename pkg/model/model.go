package model

import (
	"gomodoro/pkg/manager"
	"gomodoro/pkg/progress"
	"gomodoro/pkg/styles"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

type TeaModel struct {
	Manager  manager.Manager
	Progress progress.Model
}

func (m TeaModel) Init() tea.Cmd {
	m.updateGradient()
	return m.Manager.Init()
}

func (m *TeaModel) updateGradient() {
	m.Progress.SetPercent(0)
	if m.Manager.PomodoroState == manager.Work {
		m.Progress.SetRamp(
			styles.WorkGradient[0],
			styles.WorkGradient[1],
			false,
		)
	} else {
		m.Progress.SetRamp(
			styles.BreakGradient[0],
			styles.BreakGradient[1],
			false,
		)
	}
}

func (m TeaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.Progress.Width = msg.Width
		return m, nil

	case progress.FrameMsg:
		progressModel, cmd := m.Progress.Update(msg)
		m.Progress = progressModel.(progress.Model)
		return m, cmd

	case manager.StateChange:
		m.updateGradient()

	case manager.TickMsg:
		var cmd tea.Cmd
		m.Manager, cmd = m.Manager.UpdateState()
		incr := 0.0
		if m.Manager.Timer.Running {
			incr = 1 / float64(m.Manager.TotalDuration())
		}
		snd := m.Progress.SetPercent(incr * float64(m.Manager.Timer.Count))

		return m, tea.Batch(snd, cmd)

	case tea.KeyMsg:

		switch msg.String() {

		case "s":
			m.Manager.Skip()
			m.updateGradient()

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
	if m.Manager.PomodoroState == manager.Work {
		state = "WORK"
	} else if m.Manager.PomodoroState == manager.Break {
		state = "BREAK"
	} else {
		state = "LONG BREAK"
	}
	state += " - " + strconv.Itoa(m.Manager.Count)
	// The header
	s := "\n\n"
	stpp := ""
	if !m.Manager.Timer.Running {
		stpp += "STOPPED "
	}

	s += m.Progress.View() + "\n"

	s += styles.GetTextStyle(int(m.Manager.PomodoroState),
		m.Manager.Timer.Running,
	).Width(
		m.Progress.Width,
	).Render(
		stpp + state + "\n" + m.Manager.GetHumanTime(),
	)
	// The footer
	s += styles.HelpStyle.Width(
		m.Progress.Width,
	).Render("q - quit, enter - start/stop, s - skip")

	// Send the UI for rendering
	return s
}
