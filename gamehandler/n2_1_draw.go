package gamehandler

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	// 繪製棋盤
	// Base
	vector.DrawFilledRect(screen,
		constant.BOARD_BASE_POSITION_X, constant.BOARD_BASE_POSITION_Y,
		constant.BOARD_BASE_SIZE_WIDTH, constant.BOARD_BASE_SIZE_WIDTH,
		color.Black, true)
	g.Board.Draw(screen)
	g.Own.DrawKomaTai(screen, g.Font)
	g.Opponent.DrawKomaTai(screen, g.Font)
}
