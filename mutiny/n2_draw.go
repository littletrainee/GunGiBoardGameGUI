package mutiny

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	_color "github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// Draw 繪製謀的叛變視窗在畫面上
//
// 參數screen為要繪製的畫面
func (m *Mutiny) Draw(screen *ebiten.Image) {
	if m.Show {
		// base
		vector.DrawFilledRect(screen,
			constant.MUTINY_X, constant.MUTINY_Y, constant.MUTINY_WIDTH, constant.MUTINY_HEIGHT, color.Black, true)
		// background
		vector.DrawFilledRect(screen,
			constant.MUTINY_X+1, constant.MUTINY_Y+1, constant.MUTINY_WIDTH-2, constant.MUTINY_HEIGHT-2,
			_color.BoardColor, true)
		screen.DrawImage(m.TitleImg, m.TitleOpt)

		// mutiny button base
		vector.DrawFilledRect(screen,
			constant.MUTINY_BUTTON_X, constant.MUTINY_BUTTON_Y,
			constant.MUTINY_BUTTON_WIDTH, constant.MUTINY_BUTTON_HEIGHT, color.Black, true)
		// mutiny button base
		vector.DrawFilledRect(screen, constant.MUTINY_BUTTON_X+1, constant.MUTINY_BUTTON_Y+1,
			constant.MUTINY_BUTTON_WIDTH-2, constant.MUTINY_BUTTON_HEIGHT-2, _color.CaptureColor, true)
		screen.DrawImage(m.MutinyImg, m.MutinyOpt)

		// cancel button base
		vector.DrawFilledRect(screen,
			constant.MUTINY_BUTTON_X*1.5, constant.MUTINY_BUTTON_Y,
			constant.MUTINY_BUTTON_WIDTH, constant.MUTINY_BUTTON_HEIGHT, color.Black, true)
		// cancel button base
		vector.DrawFilledRect(screen, constant.MUTINY_BUTTON_X*1.5+1, constant.MUTINY_BUTTON_Y+1,
			constant.MUTINY_BUTTON_WIDTH-2, constant.MUTINY_BUTTON_HEIGHT-2, _color.DenyColor, true)
		screen.DrawImage(m.CancelImg, m.CancelOpt)

	}
}
