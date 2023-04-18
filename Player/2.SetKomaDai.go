package Player

import (
	"github.com/littletrainee/GunGiBoardGameGUI/Const"
	"github.com/littletrainee/GunGiBoardGameGUI/Coordinate"
	"github.com/littletrainee/GunGiBoardGameGUI/Koma"
	"github.com/littletrainee/GunGiBoardGameGUI/Koma/Move"
	"github.com/littletrainee/GunGiBoardGameGUI/Player/Pair"
)

func (p *Player) SetKomaDai(level Const.Level) {
	var (
		source   string
		row      float32 = 0
		column   float32 = 0
		distance float32 = Const.RADIUS*2 + 1
	)
	switch level {
	case Const.Beginner:
		source = Const.ALLKOMA[:30]
	case Const.Elementary:
		source = Const.ALLKOMA[:33]
	case Const.Intermediate:
	case Const.Advanced:
		source = Const.ALLKOMA
	}

	for _, v := range source {
		var (
			tempKoma Koma.Koma = Koma.Koma{
				Name:  string(v),
				Color: p.Color,
				CurrentCoordinate: Coordinate.Coordinate{
					X: p.coordinate.Item1 + distance*column,
					Y: p.coordinate.Item2 + distance*row}}
			repeat int
		)

		switch string(v) {
		case "帥":
			tempKoma.ProbablyMoveing = Move.PM帥()
			repeat = 1
		case "大":
			tempKoma.ProbablyMoveing = Move.PM大()
			repeat = 1
		case "中":
			tempKoma.ProbablyMoveing = Move.PM中()
			repeat = 1
		case "小":
			tempKoma.ProbablyMoveing = Move.PM小()
			repeat = 2
		case "侍":
			tempKoma.ProbablyMoveing = Move.PM侍()
			repeat = 2
		case "忍":
			tempKoma.ProbablyMoveing = Move.PM忍()
			repeat = 2
		case "槍":
			tempKoma.ProbablyMoveing = Move.PM槍()
			repeat = 3
		case "砦":
			tempKoma.ProbablyMoveing = Move.PM砦()
			repeat = 2
		case "馬":
			tempKoma.ProbablyMoveing = Move.PM馬()
			repeat = 2
		case "兵":
			tempKoma.ProbablyMoveing = Move.PM兵()
			repeat = 4
		case "弓":
			tempKoma.ProbablyMoveing = Move.PM弓()
			repeat = 2
		case "砲":
			tempKoma.ProbablyMoveing = Move.PM砲()
			repeat = 1
		case "筒":
			tempKoma.ProbablyMoveing = Move.PM筒()
			repeat = 1
		case "謀":
			tempKoma.ProbablyMoveing = Move.PM謀()
			repeat = 1
		}
		for i := 0; i < repeat; i++ {
			index := Contains(p.KomaDai, tempKoma)
			if index > 0 {
				p.KomaDai[index].Item2++
			} else {
				p.KomaDai = append(p.KomaDai, Pair.Pair[Koma.Koma, int]{
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

func Contains(slice []Pair.Pair[Koma.Koma, int], target Koma.Koma) int {
	for i, v := range slice {
		if v.Item1.Name == target.Name {
			return i
		}
	}
	return -1
}
