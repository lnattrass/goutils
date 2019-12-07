package watchdog

import (
	"time"
)

// The Watchdog type
type Watchdog struct {
	interval time.Duration
	timer    *time.Timer
	C        <-chan time.Time
}

// New watchdog timer
func New(interval time.Duration) *Watchdog {
	w := Watchdog{
		interval: interval,
		timer:    time.NewTimer(interval),
	}
	w.C = w.timer.C
	return &w
}

// Stop the Watchdog
func (w *Watchdog) Stop() {
	w.timer.Stop()
}

// Kick the Watchdog timer
func (w *Watchdog) Kick() {
	w.timer.Stop()
	w.timer.Reset(w.interval)
}
