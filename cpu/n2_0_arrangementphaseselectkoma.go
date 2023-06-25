package cpu

import (
	"image"
	"math/rand"

	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
)

// ArrangementPhaseSelectKoma 布陣階段的選駒，並且決定要將駒放在棋盤上的哪個位置
//
// 參數g為本局遊戲當前的狀態，b為棋盤物件
func (c *CPU) ArrangementPhaseSelectKoma(g gamestate.GameState, b *board.Board) {
	if !c.Player.IsSuMi {
		if !c.Player.IsSetSuI {
			c.MoveToTarget = []int{0, rand.Intn(9) + 1, rand.Intn(3) + 1}
			c.Player.IsSetSuI = true
			return
		}

		var (
			coor                image.Point // 座標
			targetblock         block.Block // 目標塊
			whichOneCanBeChosen []int       = []int{}
			whichOne            int
		)

		// 逐項迭代駒台上有的駒，加至待選的陣列中
		for i, v := range c.Player.KomaDai {
			if v.Item2 > 0 {
				whichOneCanBeChosen = append(whichOneCanBeChosen, i)
			}
		}
		// 設定選擇的目標
		whichOne = whichOneCanBeChosen[rand.Intn(len(whichOneCanBeChosen))]
		for {
			coor = image.Point{X: rand.Intn(9) + 1, Y: rand.Intn(3) + 1}
			targetblock = b.Blocks[coor]
			if !targetblock.HasSuI() && len(targetblock.KomaStack) < g.LevelHolder.MaxLayer {
				c.MoveToTarget = []int{whichOne, coor.X, coor.Y}
				return
			}
		}

	}
}
