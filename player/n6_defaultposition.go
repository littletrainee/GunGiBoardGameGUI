package player

import (
	"image"
	"math"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/GunGiBoardGameGUI/koma/defaultposition"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
	"github.com/littletrainee/pair"
)

func (p *Player) DefaultPosition(g gamestate.GameState, b board.Board) {
	var (
		target []pair.Pair[int, image.Point]
	)
	if g.Level == level.BEGINNER && p.IsOwn {
		target = defaultposition.OwnBeginner()
	} else if g.Level == level.BEGINNER && !p.IsOwn {
		target = defaultposition.OpponentBeginner()
	}
	for _, v := range target {
		pos := image.Point{X: v.Item2.X, Y: v.Item2.Y}
		tempblock := b.Blocks[pos]
		tempKoma := p.KomaTai[v.Item1].Item1.Clone()
		tempKoma.CurrentCoordinate.X = int(tempblock.Coordinate.X + constant.BLOCK_SIZE/2)
		tempKoma.CurrentCoordinate.Y = int(tempblock.Coordinate.Y + constant.BLOCK_SIZE/2)
		tempKoma.SetCurrentPosition(tempblock.Name.Item1, tempblock.Name.Item2)
		if p.IsOwn {
			tempKoma.SetGeoMetry(0)
		} else {
			tempKoma.SetGeoMetry(math.Pi)
		}
		tempblock.KomaStack = append(tempblock.KomaStack, tempKoma)
		p.KomaTai[v.Item1].Item2--
		if p.KomaTai[v.Item1].Item2 == 0 {
			p.KomaTai[v.Item1].Item1 = koma.Koma{}
		}
		b.Blocks[pos] = tempblock
	}
}
