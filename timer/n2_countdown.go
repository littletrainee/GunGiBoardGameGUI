package timer

import (
	"time"
)

// CountDown 用另一個線程控制棋鐘的剩餘時間
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
