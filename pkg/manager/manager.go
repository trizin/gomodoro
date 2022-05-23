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
)

type Manager struct {
	Timer         timer.Timer
	BreakDuration int
	WorkDuration  int
	PomodoroState PomodoroState
}
type TickMsg struct {
	Time int
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

func (m *Manager) TotalDuration() int {
	if m.PomodoroState == 0 {
		return m.WorkDuration
	} else {
		return m.BreakDuration
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

func (m Manager) UpdateState() (Manager, tea.Cmd) {
	if m.Timer.Running {
		m.Timer.Increase()
	}

	if m.Timer.Count >= m.BreakDuration && m.PomodoroState == Break {
		m.Timer.Reset()
		m.PomodoroState = Work
	}

	if m.Timer.Count >= m.WorkDuration && m.PomodoroState == Work {
		m.Timer.Reset()
		m.PomodoroState = Break
	}

	return m, m.tick()
}
