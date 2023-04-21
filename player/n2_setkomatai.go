package player

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/coordinate"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/GunGiBoardGameGUI/koma/move"
	_constant "github.com/littletrainee/gunginotationgenerator/constant"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/pair"
)

func (p *Player) SetKomaTai(l level.Level) {
	var (
		source   string
		row      float32 = 0
		column   float32 = 0
		distance float32 = constant.RADIUS*2 + 1
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

	for _, v := range source {
		var (
			c        float32   = p.coordinate.Item1 + distance*column
			r        float32   = p.coordinate.Item2 + distance*row
			tempKoma koma.Koma = koma.Koma{
				Name:  string(v),
				Color: p.Color,
				Img:   ebiten.NewImage(int(constant.RADIUS)+1, int(constant.RADIUS)+1),
				CurrentCoordinate: coordinate.Coordinate{
					X: c,
					Y: r},
				Op: &ebiten.DrawImageOptions{},
			}
			repeat int
		)
		if p.Color == color.Black {
			tempKoma.Op.GeoM.Rotate(0)
			tempKoma.Op.GeoM.Translate(float64(c-10), float64(r-12))
		} else {
			tempKoma.Op.GeoM.Rotate(math.Pi)
			tempKoma.Op.GeoM.Translate(float64(c+10), float64(r+12))
		}

		switch string(v) {
		case "帥":
			tempKoma.ProbablyMoveing = move.PM帥()
			repeat = 1
		case "大":
			tempKoma.ProbablyMoveing = move.PM大()
			repeat = 1
		case "中":
			tempKoma.ProbablyMoveing = move.PM中()
			repeat = 1
		case "小":
			tempKoma.ProbablyMoveing = move.PM小()
			repeat = 2
		case "侍":
			tempKoma.ProbablyMoveing = move.PM侍()
			repeat = 2
		case "忍":
			tempKoma.ProbablyMoveing = move.PM忍()
			repeat = 2
		case "槍":
			tempKoma.ProbablyMoveing = move.PM槍()
			repeat = 3
		case "砦":
			tempKoma.ProbablyMoveing = move.PM砦()
			repeat = 2
		case "馬":
			tempKoma.ProbablyMoveing = move.PM馬()
			repeat = 2
		case "兵":
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
			index := Contains(p.KomaDai, tempKoma)
			if index > 0 {
				p.KomaDai[index].Item2++
			} else {
				p.KomaDai = append(p.KomaDai, pair.Pair[koma.Koma, int]{
					Item1: tempKoma, Item2: 1})
			}
		}
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
