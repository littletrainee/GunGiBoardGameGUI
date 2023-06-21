package cpu

import (
	"image"
	"math"

	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// ArrangementPhaseMoveKoma 布陣階段放置駒到棋盤上
//
// 參數b為棋盤物件
func (c *CPU) ArrangementPhaseMoveKoma(b board.Board) {
	var (
		position image.Point = image.Point{X: c.MoveToTarget[1], Y: c.MoveToTarget[2]}
	)
	targetblock := b.Blocks[position]
	tempKoma := c.Player.KomaDai[c.MoveToTarget[0]].Item1.Clone()
	tempKoma.SetCurrentCoordinate(targetblock.Coordinate, block.Shift(targetblock.KomaStack))
	tempKoma.SetCurrentPosition(targetblock.Name)
	tempKoma.Op.GeoM.Reset()
	tempKoma.SetGeoMetry(math.Pi)

	targetblock.KomaStack = append(targetblock.KomaStack, tempKoma)
	b.Blocks[position] = targetblock
	c.Player.KomaDai[c.MoveToTarget[0]].Item2--
	if c.Player.KomaDai[c.MoveToTarget[0]].Item2 == 0 {
		c.Player.KomaDai[c.MoveToTarget[0]].Item1 = koma.Koma{}
	}
	c.MoveToTarget = nil
}
