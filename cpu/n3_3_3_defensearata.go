package cpu

import (
	"image"
	"math/rand"

	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// defenseARata 新駒防禦函式，可以導出哪個位置新可以避免被將軍的狀況，若成立則回傳true，否則回傳false
//
// 參數b為棋盤物件，maxlayer為可以堆疊的最高段數
func (c *CPU) defenseARata(b board.Board, g gamestate.GameState) bool {
	var (
		suiPosition   image.Point = image.Point{X: c.MoveToTarget[2], Y: c.MoveToTarget[3]}
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
		for orginalkey := range b.Blocks {
			if orginalkey.Y > c.MaxRange {
				continue
			}
			// 複製board 用以比較核對位置
			var cloneBoard board.Board = board.Board{Blocks: make(map[image.Point]block.Block)}
			for k, v := range b.Blocks {
				cloneBoard.Blocks[k] = v
			}
			// 確認目標的位置段數小魚最大段數 ，並且不是帥的位置
			if len(cloneBoard.Blocks[orginalkey].KomaStack) < g.LevelHolder.MaxLayer && orginalkey != suiPosition {
				// 複製目標位置並且將居放到目標位置上用以核對是否還可以使俘獲成立
				block := cloneBoard.Blocks[orginalkey]
				block.KomaStack = append(block.KomaStack, tempKoma)
				cloneBoard.Blocks[orginalkey] = block
				if !stillInCaptureRange(cloneBoard, suiPosition, c.SelfColor, g) {
					c.MoveToTarget = []int{rndSelectKoma, orginalkey.X, orginalkey.Y}
					return true
				}
			}
		}
	}
	return false
}
