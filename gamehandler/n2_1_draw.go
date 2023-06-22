package gamehandler

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
)

// Draw 實例化Game介面下的Draw，用於每禎繪製一次所需的畫面
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	switch g.GameState.Phase {
	case phase.SELECT_COLOR,
		phase.SELECT_LEVEL,
		phase.SELECT_RECOMMEND_OR_MANUAL_ARRANGEMENT:
		g.GameState.Draw(screen)

	// 對弈前的倒數計時
	case phase.BEGIN_COUNTDOWN_FOR_GAMING:
		g.PrepareForGameing.Draw(screen, g.Font, g.GameState.Phase)

	case phase.SELECT_KOMA,
		phase.MOVE_KOMA,
		phase.CLICK_CLOCK,
		phase.PLAYER_CAPTURE_OR_CONTROL_ASK,
		phase.ANOTHER_ROUND_OR_END,
		phase.BUFFER_ZONE:

		// 繪製棋盤
		// Base
		vector.DrawFilledRect(screen,
			constant.BOARD_X, constant.BOARD_Y,
			constant.BOARD_WIDTH, constant.BOARD_WIDTH,
			color.Black, true)
		g.Board.Draw(screen)
		g.Player1.DrawKomaTai(screen, g.Font)
		g.Player2.DrawKomaTai(screen, g.Font)
		vector.DrawFilledRect(screen,
			float32(constant.TIMER_X), float32(constant.TIMER_Y),
			float32(constant.TIMER_WIDTH), float32(constant.TIMER_HEIGHT),
			color.Black, true)
		g.Player1Timer.Draw(screen, g.Font, g.GameState.Phase)
		g.Player2Timer.Draw(screen, g.Font, g.GameState.Phase)
		g.Capture.Draw(screen)
		g.AnotherRoundOrEnd.Draw(screen)
		g.DeclareSuMi.Draw(screen)
	}
}
