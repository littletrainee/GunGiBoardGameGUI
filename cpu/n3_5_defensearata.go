package cpu

import (
	"image"
	"math/rand"

	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

func (c *CPU) defenseARata(b board.Board, maxlayer int) {
	var (
		suiPosition   image.Point = positionSuI(b, c)
		tempKoma      koma.Koma
		komaTaiList   []int
		rndSelectKoma int
	)
	for i, v := range c.KomaDai {
		if v.Item2 > 0 {
			komaTaiList = append(komaTaiList, i)
		}
	}
	if len(komaTaiList) > 0 {
		rndSelectKoma = komaTaiList[rand.Intn(len(komaTaiList))]
		// 設定假定的駒
		tempKoma = c.KomaDai[rndSelectKoma].Item1.Clone()
		// 逐項迭代block
		for k := range b.Blocks {
			var cloneBoard board.Board = board.Board{Blocks: make(map[image.Point]block.Block)}
			// 複製board
			for k, v := range b.Blocks {
				cloneBoard.Blocks[k] = v
			}
			if len(cloneBoard.Blocks[k].KomaStack) < maxlayer && k != suiPosition {
				block := cloneBoard.Blocks[k]
				block.KomaStack = append(block.KomaStack, tempKoma)
				cloneBoard.Blocks[k] = block
				if !stillInCaptureRange(cloneBoard, suiPosition, c.SelfColor) {
					c.MoveToTarget = []int{rndSelectKoma, k.X, k.Y}
					return
				}
			}
		}
	}
}
