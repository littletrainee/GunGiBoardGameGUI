package gamehandler

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	switch g.GameState.Phase {
	// 選擇顏色
	case phase.SELECT_COLOR:
		g.GameState.Draw(screen, g.Font)

	// 選擇階級
	case phase.SELECT_LEVEL:
		g.GameState.Draw(screen, g.Font)

	// 若階級是入門或初級詢問是否用推薦的配置
	case phase.RECOMMENDED_ARRANGEMENT:
		g.GameState.Draw(screen, g.Font)

	// 對弈前的倒數計時
	case phase.PREPARE_FOR_GAMING:
		g.PrepareForGameing.Draw(screen, g.Font, g.GameState.Phase)

	case phase.ARRANGEMENT_PHASE,
		phase.DUELING_PHASE_SELECT_KOMA,
		phase.DUELING_PHASE_MOVE_KOMA,
		phase.DUELING_PHASE_CLICK_CLOCK,
		phase.DUELING_PHASE_CAPTURE_OR_CONTROL_ASK,
		phase.CPU_SELECT_MOVE,
		phase.CPU_MOVE_KOMA,
		phase.CPU_CLICK_CLOCK,
		phase.DELAY:

		// 繪製棋盤
		// Base
		vector.DrawFilledRect(screen,
			constant.BOARD_BASE_POSITION_X, constant.BOARD_BASE_POSITION_Y,
			constant.BOARD_BASE_SIZE_WIDTH, constant.BOARD_BASE_SIZE_WIDTH,
			color.Black, true)
		g.Board.Draw(screen, g.Font)
		g.Player1.DrawKomaTai(screen, g.Font)
		g.Player2.DrawKomaTai(screen, g.Font)
		vector.DrawFilledRect(screen,
			float32(constant.TIMER_POSITION_X), float32(constant.TIMER_POSITION_Y),
			float32(constant.TIMER_BASE_WIDTH), float32(constant.TIMER_BASE_HEIGHT),
			color.Black, true)
		g.Player1Timer.Draw(screen, g.Font, g.GameState.Phase)
		g.Player2Timer.Draw(screen, g.Font, g.GameState.Phase)
		g.Capture.Draw(screen)
	}
}
