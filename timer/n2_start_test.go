package timer

import (
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	var T Timer = Timer{RemainingTime: 20}
	T.CountDown()
	time.Sleep(3 * time.Second)
	time.Sleep(3 * time.Second)
	for T.RemainingTime != 0 {
		time.Sleep(time.Second)
	}
}
