package player

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/playerstate"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"golang.org/x/image/font"
)

func (p *Player) SetKomaTaiPosition(l level.Level, font font.Face) {
	if p.IsOwn {
		p.KomaTaiCoordinate = image.Point{
			X: constant.OWN_KOMATAI_POSITION_X,
			Y: constant.OWN_KOMATAI_POSITION_Y,
		}
		p.KomaOnKomaTaiCoordinate = image.Point{
			X: constant.KOMA_ON_OWN_KOMATAI_X,
			Y: constant.KOMA_ON_OWN_KOMATAI_Y,
		}
	} else {
		p.KomaTaiCoordinate = image.Point{
			X: constant.OPPONENT_KOMATAI_POSITION_X,
			Y: constant.OPPONENT_KOMATAI_POSITION_Y,
		}
		p.KomaOnKomaTaiCoordinate = image.Point{
			X: constant.KOMA_ON_OPPONENT_KOMATAI_X,
			Y: constant.KOMA_ON_OPPONENT_KOMATAI_Y,
		}
	}
	p.State = playerstate.WAIT_FOR_SELECT_KOMA
	p.SetKomaTai(l, font)
}
