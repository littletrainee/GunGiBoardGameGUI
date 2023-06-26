package player

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// Draw 繪製駒台與駒台列表
//
// 參數screen為要繪製的畫面
func (p *Player) Draw(screen *ebiten.Image) {
	if p.IsOwn {
		vector.DrawFilledRect(screen, constant.OWN_KOMADAI_X,
			constant.OWN_KOMADAI_Y, constant.KOMADAI_SIZE, constant.KOMADAI_SIZE,
			color.Black, true)
		// area
		vector.DrawFilledRect(screen, constant.OWN_KOMADAI_X+1,
			constant.OWN_KOMADAI_Y+1, constant.KOMADAI_SIZE-2,
			constant.KOMADAI_SIZE-2, p.KomaDaiBackground, true)
	} else {
		vector.DrawFilledRect(screen, constant.OPPONENT_KOMADAI_POSITION_X,
			constant.OPPONENT_KOMADAI_POSITION_Y, constant.KOMADAI_SIZE, constant.KOMADAI_SIZE,
			color.Black, true)

		vector.DrawFilledRect(screen, constant.OPPONENT_KOMADAI_POSITION_X+1,
			constant.OPPONENT_KOMADAI_POSITION_Y+1, constant.KOMADAI_SIZE-2,
			constant.KOMADAI_SIZE-2, p.KomaDaiBackground, true)
	}
	for _, v := range p.KomaDai {
		if len(v) > 0 {
			v[0].Draw(screen)
		}
	}
}
