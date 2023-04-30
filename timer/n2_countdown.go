package timer

import (
	"time"
)

func (t *Timer) CountDown() {
	go func() {
		for {
			select {
			case <-t.StopCountDown:
				return
			case <-time.After(time.Second):
				if t.RemainingTime > 0 {
					t.RemainingTime--
				}
			}
		}
	}()
}
