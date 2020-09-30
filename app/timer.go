package app

import (
	"time"

	"github.com/wailsapp/wails"
)

// Timer holds the ticker for monitoring work time
// and chill time.
type Timer struct {
	runtime      *wails.Runtime
	log          *wails.CustomLogger
	waittime     time.Duration
	worktime     time.Duration
	ticker       *time.Ticker
	seconds      *time.Ticker
	waiting      bool
	done, reset  chan struct{}
	secondPassed int
}

// SkipBreak resets to work time without finishing
// the chill timer.
func (t *Timer) SkipBreak() {
	if t.waiting {
		t.runtime.Window.UnFullscreen()
		t.ticker.Reset(t.worktime)
		t.waiting = false
		t.runtime.Events.Emit("add-chill-time", t.secondPassed)
		t.secondPassed = 0
		t.runtime.Events.Emit("working")
	}
}

// StartBreak starts chill time immediately.
func (t *Timer) StartBreak() {
	if !t.waiting {
		t.runtime.Window.Fullscreen()
		t.waiting = true
		t.runtime.Events.Emit("add-work-time", t.secondPassed)
		t.secondPassed = 0
		t.runtime.Events.Emit("chilling")
	}
}

func (t *Timer) EndBreak() {
	if t.waiting {
		t.runtime.Events.Emit("add-chill-time", t.secondPassed)
		t.secondPassed = 0
		t.reset <- struct{}{}
	}
}

// NewTimer creates a new Ticker and starts the goroutine
// to monitor it.
func NewTimer(worktime, waittime int) *Timer {
	t := &Timer{
		waittime:     time.Duration(waittime) * time.Second,
		worktime:     time.Duration(worktime) * time.Second,
		ticker:       time.NewTicker(time.Duration(worktime) * time.Second),
		seconds:      time.NewTicker(time.Second),
		waiting:      false,
		done:         make(chan struct{}),
		reset:        make(chan struct{}),
		secondPassed: 0,
	}

	go func() {
		for {
			select {
			case <-t.ticker.C:
				t.StartBreak()
			case <-t.seconds.C:
				t.secondPassed++
				t.runtime.Events.Emit("tick", t.secondPassed)
				if t.waiting {
					if time.Duration(t.secondPassed)*time.Second == t.waittime {
						t.runtime.Events.Emit("endable")
					}
				}
			case <-t.reset:
				t.SkipBreak()
			case <-t.done:
				return
			}
		}
	}()

	return t
}

// WailsInit is used when binding to let Timer access the runtime.
func (t *Timer) WailsInit(runtime *wails.Runtime) error {
	t.runtime = runtime
	t.log = runtime.Log.New("Timer")
	return nil
}
