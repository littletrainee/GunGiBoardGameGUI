package player

import (
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	_color "github.com/littletrainee/gunginotationgenerator/enum/color"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/pair"
)

func (p *Player) Initilization(s string, c _color.Color, l level.Level) {
	p.Name = s
	if c == _color.BLACK {
		p.Color = color.Black
		p.coordinate = pair.Pair[float32, float32]{
			Item1: constant.KOMA_ON_OWN_KOMADAI_X,
			Item2: constant.KOMA_ON_OWN_KOMADAI_Y}
	} else {
		p.Color = color.White
		p.coordinate = pair.Pair[float32, float32]{
			Item1: constant.KOMA_ON_OPPONENT_KOMADAI_X,
			Item2: constant.KOMA_ON_OPPONENT_KOMADAI_Y}
	}
	p.KomaDai = []pair.Pair[koma.Koma, int]{}
	p.SetKomaTai(l)
}
