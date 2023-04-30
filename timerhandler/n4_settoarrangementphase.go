package timerhandler

import (
	"time"

	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func (t *TimerHandler) SetToArrangementPhase() {
	// 對StopCount傳送訊號使得這個執行續離開
	t.Player1Timer.StopCountDown <- true
	t.Player2Timer.StopCountDown <- true
	t.Player1Timer.RemainingTime = constant.REMAINING_TIME
	t.Player2Timer.RemainingTime = constant.CPU_REMAINING_TIME
	t.Player1Timer.BackgroundColor = color.ConfirmColor
	t.Player1Timer.CountDown()
	t.Player2Timer.CountDown()
	// 等待0.5秒後關閉另一個執行續，使其留下僅有一個繼續
	time.Sleep(time.Second)
	t.Player2Timer.StopCountDown <- true
}
