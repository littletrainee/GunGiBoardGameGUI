package GameHolder

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/Const"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	// 繪製棋盤
	// Base
	vector.DrawFilledRect(screen,
		Const.BOARD_BASE_POSITION_X, Const.BOARD_BASE_POSITION_Y,
		Const.BOARD_BASE_SIZE_WIDTH, Const.BOARD_BASE_SIZE_WIDTH,
		color.Black, true)
	g.Board.Draw(screen)
	g.Own.DrawKomaDai(screen, g.Font)
	g.Opponent.DrawKomaDai(screen, g.Font)
}
