package timer

type Timer struct {
	Running bool
	Count   int
}

func (t *Timer) Stop() {
	t.Running = false
}

func (t *Timer) Start() {
	t.Running = true
}

func (t *Timer) Toggle() {
	if t.Running {
		t.Stop()
	} else {
		t.Start()
	}
}

func (t *Timer) Increase() {
	t.Count++
}

func (t *Timer) Reset() {
	t.Count = 0
}
