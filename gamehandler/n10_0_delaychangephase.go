package gamehandler

import (
	"time"

	_phase "github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/gunginotationgenerator/enum/phase"
)

// delayedChangePhaseTo 延遲切換階段，讓程式先進入延遲階段，待0.5秒後再進入欲進入的階段
//
// 參數p為目標延遲後欲進入的目標階段
func (g *Game) delayedChangePhaseTo(p phase.Phase) {
	g.GameState.Phase = _phase.BUFFER_ZONE
	go func() {
		time.Sleep(time.Second / 2)
		g.GameState.Phase = p
	}()
}
