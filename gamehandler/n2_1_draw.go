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

	case phase.ARRANGEMENT_PHASE, phase.DUELING_PHASE:

		// 繪製棋盤
		// Base
		vector.DrawFilledRect(screen,
			constant.BOARD_BASE_POSITION_X, constant.BOARD_BASE_POSITION_Y,
			constant.BOARD_BASE_SIZE_WIDTH, constant.BOARD_BASE_SIZE_WIDTH,
			color.Black, true)
		g.Board.Draw(screen, g.Font)
		g.Player1.DrawKomaTai(screen, g.Font)
		g.Player2.DrawKomaTai(screen, g.Font)
		// g.TimerPair.Draw(screen, g.Font)
	}
}
