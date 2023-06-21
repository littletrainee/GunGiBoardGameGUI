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
	case phase.ASK_ARRANGEMENT:
		g.GameState.Draw(screen, g.Font)

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
		g.Board.Draw(screen, g.Font)
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
