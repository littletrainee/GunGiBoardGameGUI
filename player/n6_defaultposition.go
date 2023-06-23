package player

import (
	"image"
	"math"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/GunGiBoardGameGUI/koma/recommendposition"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/pair"
)

// 推薦位置
func (p *Player) RecommendPosition(g gamestate.GameState, b board.Board) {
	var (
		target []pair.Pair[int, image.Point]
	)
	if g.LevelHolder.CurrentLevel == level.BEGINNER {
		if p.IsOwn {
			target = recommendposition.OwnBeginner()
		} else {
			target = recommendposition.OpponentBeginner()
		}
	} else if g.LevelHolder.CurrentLevel == level.ELEMENTARY {
		if p.IsOwn {
			target = recommendposition.OwnElementary()
		} else {
			target = recommendposition.OpponentElementary()
		}
	}
	for _, v := range target {
		pos := image.Point{X: v.Item2.X, Y: v.Item2.Y}
		tempblock := b.Blocks[pos]
		tempKoma := p.KomaDai[v.Item1].Item1.Clone()
		tempKoma.SetCurrentCoordinate(tempblock.Coordinate, 0)
		tempKoma.SetCurrentPosition(tempblock.Name)
		if p.IsOwn {
			tempKoma.SetGeoMetry(0)
		} else {
			tempKoma.SetGeoMetry(math.Pi)
		}
		tempblock.KomaStack = append(tempblock.KomaStack, tempKoma)
		p.KomaDai[v.Item1].Item2--
		if p.KomaDai[v.Item1].Item2 == 0 {
			p.KomaDai[v.Item1].Item1 = koma.Koma{}
		}
		b.Blocks[pos] = tempblock
	}
}
