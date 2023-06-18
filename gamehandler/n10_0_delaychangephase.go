package gamehandler

import (
	"time"

	_phase "github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/gunginotationgenerator/enum/phase"
)

func (g *Game) delayedChangePhaseTo(p phase.Phase) {
	g.GameState.Phase = _phase.BUFFER_ZONE
	go func() {
		time.Sleep(time.Second / 2)
		g.GameState.Phase = p
	}()
}
