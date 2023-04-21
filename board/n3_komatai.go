package board

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	_color "github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func KomaTai(screen *ebiten.Image) {
	// own koma dai
	// base
	vector.DrawFilledRect(screen, constant.OWN_KOMADAI_POSITION_X,
		constant.OWN_KOMADAI_POSITION_Y, constant.KOMADAI_SIZE, constant.KOMADAI_SIZE,
		color.Black, true)
	// area
	vector.DrawFilledRect(screen, constant.OWN_KOMADAI_POSITION_X+1,
		constant.OWN_KOMADAI_POSITION_Y+1, constant.KOMADAI_SIZE-2,
		constant.KOMADAI_SIZE-2, _color.BoardColor, true)

	// opponent koma dai
	// base
	vector.DrawFilledRect(screen, constant.OPPONENT_POSITION_X,
		constant.OPPONENT_POSITION_Y, constant.KOMADAI_SIZE, constant.KOMADAI_SIZE,
		color.Black, true)

	vector.DrawFilledRect(screen, constant.OPPONENT_POSITION_X+1,
		constant.OPPONENT_POSITION_Y+1, constant.KOMADAI_SIZE-2,
		constant.KOMADAI_SIZE-2, _color.BoardColor, true)

}
