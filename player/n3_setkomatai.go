package player

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/GunGiBoardGameGUI/koma/moveablerange"
	_constant "github.com/littletrainee/gunginotationgenerator/constant"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/pair"
	"golang.org/x/image/font"
)

func (p *Player) SetKomaTai(l level.Level, font font.Face) {
	var (
		source       string
		row          float32 = 0
		column       float32 = 0
		distance     float32 = constant.KOMA_SIZE*2 + 1
		numberOfKoma int
	)
	switch l {
	case level.BEGINNER:
		source = _constant.ALLKOMA[:30]
	case level.ELEMENTARY:
		source = _constant.ALLKOMA[:33]
	case level.INTERMEDIATE,
		level.ADVANCED:
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
			c        float32   = float32(p.KomaOnKomaDaiCoordinate.X) + distance*column
			r        float32   = float32(p.KomaOnKomaDaiCoordinate.Y) + distance*row
			tempKoma koma.Koma = koma.Koma{
				Name:  string(v),
				Color: p.SelfColor,
				CurrentCoordinate: image.Point{
					X: int(c),
					Y: int(r)},
				Op:  &ebiten.DrawImageOptions{},
				Img: ebiten.NewImage(int(constant.KOMA_SIZE)+1, int(constant.KOMA_SIZE)+1),
			}
			repeat int
		)
		text.Draw(tempKoma.Img, tempKoma.Name, font, 0, 20, color.Black)

		if p.IsOwn {
			tempKoma.SetGeoMetry(0)
		} else {
			tempKoma.SetGeoMetry(math.Pi)
		}

		switch string(v) {
		case "帥": // 0
			tempKoma.MoveableRange = moveablerange.MoveableRange帥()
			repeat = 1
		case "大": // 1
			tempKoma.MoveableRange = moveablerange.MoveableRange大()
			repeat = 1
		case "中": // 2
			tempKoma.MoveableRange = moveablerange.MoveableRange中()
			repeat = 1
		case "小": // 3
			tempKoma.MoveableRange = moveablerange.MoveableRange小()
			repeat = 2
		case "侍": // 4
			tempKoma.MoveableRange = moveablerange.MoveableRange侍()
			repeat = 2

		case "忍": // 5
			tempKoma.MoveableRange = moveablerange.MoveableRange忍()
			repeat = 2
		case "槍": // 6
			tempKoma.MoveableRange = moveablerange.MoveableRange槍()
			repeat = 3
		case "砦": // 7
			tempKoma.MoveableRange = moveablerange.MoveableRange砦()
			repeat = 2
		case "馬": // 8
			tempKoma.MoveableRange = moveablerange.MoveableRange馬()
			repeat = 2
		case "兵": // 9
			tempKoma.MoveableRange = moveablerange.MoveableRange兵()
			repeat = 4

		case "弓":
			tempKoma.MoveableRange = moveablerange.MoveableRange弓()
			repeat = 2
		case "砲":
			tempKoma.MoveableRange = moveablerange.MoveableRange砲()
			repeat = 1
		case "筒":
			tempKoma.MoveableRange = moveablerange.MoveableRange筒()
			repeat = 1
		case "謀":
			tempKoma.MoveableRange = moveablerange.MoveableRange謀()
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
