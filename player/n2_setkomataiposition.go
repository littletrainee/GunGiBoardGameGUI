package player

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"golang.org/x/image/font"
)

func (p *Player) SetKomaTaiPosition(l level.Level, font font.Face) {
	if p.IsOwn {
		p.KomaDaiCoordinate = image.Point{
			X: constant.OWN_KOMADAI_POSITION_X,
			Y: constant.OWN_KOMADAI_POSITION_Y,
		}
		p.KomaOnKomaDaiCoordinate = image.Point{
			X: constant.KOMA_ON_OWN_KOMADAI_X,
			Y: constant.KOMA_ON_OWN_KOMADAI_Y,
		}
	} else {
		p.KomaDaiCoordinate = image.Point{
			X: constant.OPPONENT_KOMADAI_POSITION_X,
			Y: constant.OPPONENT_KOMADAI_POSITION_Y,
		}
		p.KomaOnKomaDaiCoordinate = image.Point{
			X: constant.KOMA_ON_OPPONENT_KOMADAI_X,
			Y: constant.KOMA_ON_OPPONENT_KOMADAI_Y,
		}
	}
	p.SetKomaTai(l, font)
	// 設定駒台的背景顏色
	p.KomaDaiBackGroundColor = color.BoardColor
}
