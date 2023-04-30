package timerhandler

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/gunginotationgenerator/enum/phase"
	"golang.org/x/image/font"
)

func (t *TimerHandler) Draw(screen *ebiten.Image, font font.Face, phase phase.Phase) {
	vector.DrawFilledRect(screen,
		float32(constant.TIMER_POSITION_X), float32(constant.TIMER_POSITION_Y),
		float32(constant.TIMER_BASE_WIDTH), float32(constant.TIMER_BASE_HEIGHT),
		color.Black, true)
	// opponent
	t.Player2Timer.Draw(screen, font, phase)
	// own
	t.Player1Timer.Draw(screen, font, phase)
}
