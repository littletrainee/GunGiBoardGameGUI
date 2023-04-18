package Player

import (
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/Const"
	"github.com/littletrainee/GunGiBoardGameGUI/Koma"
	"github.com/littletrainee/GunGiBoardGameGUI/Player/Pair"
)

func (p *Player) Initilization(s string, c Const.Color, level Const.Level) {
	p.Name = s
	if c == Const.Black {
		p.Color = color.Black
		p.coordinate = Pair.Pair[float32, float32]{
			Item1: Const.KOMA_ON_OWN_KOMADAI_X,
			Item2: Const.KOMA_ON_OWN_KOMADAI_Y}
	} else {
		p.Color = color.White
		p.coordinate = Pair.Pair[float32, float32]{
			Item1: Const.KOMA_ON_OPPONENT_KOMADAI_X,
			Item2: Const.KOMA_ON_OPPONENT_KOMADAI_Y}
	}
	p.KomaDai = []Pair.Pair[Koma.Koma, int]{}
	p.SetKomaDai(level)
}
