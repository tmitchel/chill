package app

import (
	"time"

	"github.com/wailsapp/wails"
)

// Timer holds the ticker for monitoring work time
// and chill time.
type Timer struct {
	runtime  *wails.Runtime
	log      *wails.CustomLogger
	waittime time.Duration
	worktime time.Duration
	ticker   *time.Ticker
	waiting  bool
	done     <-chan struct{}
}

// SkipBreak resets to work time without finishing
// the chill timer.
func (t *Timer) SkipBreak() {
	if t.waiting {
		t.runtime.Window.UnFullscreen()
		t.waiting = false
		t.ticker.Reset(t.worktime)
	}
}

// StartBreak starts chill time immediately.
func (t *Timer) StartBreak() {
	if !t.waiting {
		t.runtime.Window.Fullscreen()
		t.waiting = true
		t.ticker.Reset(t.waittime)
	}
}

// NewTimer creates a new Ticker and starts the goroutine
// to monitor it.
func NewTimer(waittime, worktime int) *Timer {
	t := &Timer{
		waittime: time.Duration(waittime) * time.Second,
		worktime: time.Duration(worktime) * time.Second,
		ticker:   time.NewTicker(time.Duration(worktime) * time.Second),
		waiting:  false,
	}

	go func() {
		for {
			select {
			case <-t.ticker.C:
				t.log.Info("Tick")
				if t.waiting {
					// t.runtime.Window.UnFullscreen()
					t.ticker.Reset(t.worktime)
					t.waiting = false
				} else {
					// t.runtime.Window.Fullscreen()
					t.ticker.Reset(t.waittime)
					t.waiting = true
				}
			case <-t.done:
				return
			}
		}
	}()

	return t
}

// WailsInit is used when binding to let Timer access the runtime.Log
func (t *Timer) WailsInit(runtime *wails.Runtime) error {
	t.runtime = runtime
	t.log = runtime.Log.New("Timer")
	return nil
}
