package manager

import (
	"gomodoro/pkg/timer"
	"math"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type PomodoroState int

const (
	Work PomodoroState = iota
	Break
	LongBreak
)

type Manager struct {
	Timer             timer.Timer
	BreakDuration     int
	WorkDuration      int
	LongBreakDuration int
	PomodoroState     PomodoroState
	Count             int
}
type TickMsg struct {
	Time int
}

type StateChange struct {
	State int
}

func (m *Manager) GetHumanTime() string {
	n := m.TotalDuration() - m.Timer.Count
	minute := math.Ceil(float64(n / 60))
	second := n % 60

	minuteString := strconv.Itoa(int(minute))
	secondString := strconv.Itoa(second)

	if minute < 10 {
		minuteString = "0" + minuteString
	}
	if second < 10 {
		secondString = "0" + secondString
	}

	return minuteString + ":" + secondString

}

func (m *Manager) Skip() {
	m.Timer.Count = 1e10
	m.UpdateState()
}

func (m *Manager) TotalDuration() int {
	if m.PomodoroState == Work {
		return m.WorkDuration
	} else if m.PomodoroState == Break {
		return m.BreakDuration
	} else {
		return m.LongBreakDuration
	}
}

func (m Manager) Init() tea.Cmd {
	return m.tick()
}

func (m Manager) tick() tea.Cmd {
	return tea.Tick(time.Second, func(_ time.Time) tea.Msg {
		return TickMsg{Time: m.Timer.Count}
	})
}

func (m Manager) state() tea.Cmd {
	return func() tea.Msg {
		return StateChange{State: int(m.PomodoroState)}
	}
}

func (m Manager) UpdateState() (Manager, tea.Cmd) {
	if m.Timer.Running {
		m.Timer.Increase()
	}

	if m.Timer.Count >= m.BreakDuration && m.PomodoroState == Break {
		m.Timer.Reset()
		m.PomodoroState = Work
		return m, tea.Batch(m.tick(), m.state())
	}

	if m.Timer.Count >= m.LongBreakDuration && m.PomodoroState == LongBreak {
		m.Timer.Reset()
		m.PomodoroState = Work
		return m, tea.Batch(m.tick(), m.state())
	}

	if m.Timer.Count >= m.WorkDuration && m.PomodoroState == Work {
		m.Count++
		m.Timer.Reset()
		if m.Count%4 == 0 && m.Count > 0 {
			m.PomodoroState = LongBreak
		} else {
			m.PomodoroState = Break
		}
		return m, tea.Batch(m.tick(), m.state())
	}

	return m, m.tick()
}
