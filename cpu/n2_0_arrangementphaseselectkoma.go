package cpu

import (
	"image"
	"math/rand"
	"time"

	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
)

func (c *CPU) ArrangementPhaseSelectKoma(g gamestate.GameState, b *board.Board) {
	if !c.Player.IsSuMi {
		var (
			coor        image.Point // 座標
			targetblock block.Block // 目標塊
		)
		if !c.Player.IsSetSuI {
			rand.Seed(time.Now().UnixNano())
			c.MoveToTarget = []int{0, rand.Intn(9) + 1, rand.Intn(3) + 1}
			c.Player.IsSetSuI = true
		} else {
			var (
				whichOneCanBeChosen []int = []int{}
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
				if !targetblock.HasSuI() && len(targetblock.KomaStack) < g.MaxLayer {
					c.MoveToTarget = []int{whichOne, coor.X, coor.Y}
					return
				}
			}
		}
	}
}
