package Board

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/Color"
	"github.com/littletrainee/GunGiBoardGameGUI/Const"
)

func KomaDai(screen *ebiten.Image) {
	// own koma dai
	// base
	vector.DrawFilledRect(screen, Const.OWN_KOMADAI_POSITION_X,
		Const.OWN_KOMADAI_POSITION_Y, Const.KOMADAI_SIZE, Const.KOMADAI_SIZE,
		color.Black, true)
	// area
	vector.DrawFilledRect(screen, Const.OWN_KOMADAI_POSITION_X+1,
		Const.OWN_KOMADAI_POSITION_Y+1, Const.KOMADAI_SIZE-2,
		Const.KOMADAI_SIZE-2, Color.BoardColor, true)

	// opponent koma dai
	// base
	vector.DrawFilledRect(screen, Const.OPPONENT_POSITION_X,
		Const.OPPONENT_POSITION_Y, Const.KOMADAI_SIZE, Const.KOMADAI_SIZE,
		color.Black, true)

	vector.DrawFilledRect(screen, Const.OPPONENT_POSITION_X+1,
		Const.OPPONENT_POSITION_Y+1, Const.KOMADAI_SIZE-2,
		Const.KOMADAI_SIZE-2, Color.BoardColor, true)

}
