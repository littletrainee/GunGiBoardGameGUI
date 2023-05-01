package player

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/GunGiBoardGameGUI/koma/move"
	_constant "github.com/littletrainee/gunginotationgenerator/constant"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/pair"
)

func (p *Player) SetKomaTai(l level.Level) {
	var (
		source       string
		row          float32 = 0
		column       float32 = 0
		distance     float32 = constant.RADIUS*2 + 1
		numberOfKoma int
	)
	switch l {
	case level.BEGINNER:
		source = _constant.ALLKOMA[:30]
	case level.ELEMENTARY:
		source = _constant.ALLKOMA[:33]
	case level.INTERMEDIATE:
	case level.ADVANCED:
		source = _constant.ALLKOMA
	}
	if !p.IsOwn {
		distance *= -1
		p.MaxRange = 3
	} else {
		p.MaxRange = 7
	}

	for _, v := range source {
		var (
			c        float32   = float32(p.KomaOnKomaTaiCoordinate.X) + distance*column
			r        float32   = float32(p.KomaOnKomaTaiCoordinate.Y) + distance*row
			tempKoma koma.Koma = koma.Koma{
				Name:  string(v),
				Color: p.SelfColor,
				CurrentCoordinate: image.Point{
					X: int(c),
					Y: int(r)},
				Op: &ebiten.DrawImageOptions{},
			}
			repeat int
		)

		if p.IsOwn {
			tempKoma.SetGeoMetry(0)
		} else {
			tempKoma.SetGeoMetry(math.Pi)
		}

		switch string(v) {
		case "帥": // 0
			tempKoma.ProbablyMoveing = move.PM帥()
			repeat = 1
		case "大": // 1
			tempKoma.ProbablyMoveing = move.PM大()
			repeat = 1
		case "中": // 2
			tempKoma.ProbablyMoveing = move.PM中()
			repeat = 1
		case "小": // 3
			tempKoma.ProbablyMoveing = move.PM小()
			repeat = 2
		case "侍": // 4
			tempKoma.ProbablyMoveing = move.PM侍()
			repeat = 2

		case "忍": // 5
			tempKoma.ProbablyMoveing = move.PM忍()
			repeat = 2
		case "槍": // 6
			tempKoma.ProbablyMoveing = move.PM槍()
			repeat = 3
		case "砦": // 7
			tempKoma.ProbablyMoveing = move.PM砦()
			repeat = 2
		case "馬": // 8
			tempKoma.ProbablyMoveing = move.PM馬()
			repeat = 2
		case "兵": // 9
			tempKoma.ProbablyMoveing = move.PM兵()
			repeat = 4

		case "弓":
			tempKoma.ProbablyMoveing = move.PM弓()
			repeat = 2
		case "砲":
			tempKoma.ProbablyMoveing = move.PM砲()
			repeat = 1
		case "筒":
			tempKoma.ProbablyMoveing = move.PM筒()
			repeat = 1
		case "謀":
			tempKoma.ProbablyMoveing = move.PM謀()
			repeat = 1
		}
		for i := 0; i < repeat; i++ {
			index := Contains(p.KomaTai, tempKoma)
			if index > 0 {
				p.KomaTai[index].Item2++
			} else {
				p.KomaTai = append(p.KomaTai, pair.Pair[koma.Koma, int]{
					Item1: tempKoma, Item2: 1})
			}
		}
		numberOfKoma += repeat
		if column < 3 {
			column++
		} else {
			row++
			column = 0
		}
	}
}

func Contains(slice []pair.Pair[koma.Koma, int], target koma.Koma) int {
	for i, v := range slice {
		if v.Item1.Name == target.Name {
			return i
		}
	}
	return -1
}
