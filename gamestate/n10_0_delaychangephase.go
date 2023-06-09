package gamestate

import (
	"time"

	_phase "github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/gunginotationgenerator/enum/phase"
)

// DelayedChangePhaseTo 延遲切換階段，讓程式先進入延遲階段，待0.5秒後再進入欲進入的階段
//
// 參數p為目標延遲後欲進入的目標階段
func (g *GameState) DelayedChangePhaseTo(p phase.Phase) {
	switch p {
	case _phase.SET_COLOR, _phase.SET_LEVEL, _phase.SET_COUNTDOWN_FOR_GAMING:
		g.Phase = _phase.BEFORE_GAMING_BUFFER_ZONE
	default:
		g.Phase = _phase.GAMING_BUFFER_ZONE
	}
	go func() {
		time.Sleep(time.Second / 2)
		g.Phase = p
	}()
}
