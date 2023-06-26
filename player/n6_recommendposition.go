package player

import (
	"image"
	"math"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/koma/recommendposition"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/pair"
)

// RecommendPosition 在入門與初級階段的推薦布陣配置
//
// 參數g為遊戲的階級來源，b為棋盤物件
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
		tempKoma := p.KomaDai[v.Item1][0].Clone()
		tempLen := len(p.KomaDai[v.Item1])
		tempKoma.SetCurrentCoordinate(tempblock.Coordinate, 0)
		tempKoma.SetCurrentPosition(tempblock.Name)
		if p.IsOwn {
			tempKoma.SetGeoMetry(0)
		} else {
			tempKoma.SetGeoMetry(math.Pi)
		}
		tempblock.KomaStack = append(tempblock.KomaStack, tempKoma)
		// p.KomaDai[v.Item1].Item2--
		p.KomaDai[v.Item1] = p.KomaDai[v.Item1][:tempLen-1]
		b.Blocks[pos] = tempblock
	}
}
