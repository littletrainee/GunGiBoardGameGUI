package cpu

import (
	"image"
	"math"
	"math/rand"

	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/playerstate"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

func (c *CPU) AutoSetKoma(g gamestate.GameState, b *board.Board) {
	if !c.Player.IsSuMi {
		var (
			rndRow      int
			rndColumn   int
			tempKoma    koma.Koma
			coor        image.Point
			targetblock block.Block
		)
		if g.ArrangementPhaseRoundCount == 1 &&
			c.Player.State == playerstate.IS_SELECT_KOMA {
			rndColumn = rand.Intn(9)
			rndRow = rand.Intn(3)
			coor = image.Point{X: rndColumn, Y: rndRow}
			targetblock = b.Blocks[coor]
			switch len(targetblock.KomaStack) {
			}
			tempKoma = c.Player.KomaTai[0].Item1
			tempKoma.CurrentCoordinate.X = int(targetblock.Coordinate.X + constant.BLOCK_SIZE/2)
			tempKoma.CurrentCoordinate.Y = int(targetblock.Coordinate.Y + constant.BLOCK_SIZE/2)
			tempKoma.SetCurrentPosition(targetblock.Name.Item1, targetblock.Name.Item2)
			tempKoma.Op.GeoM.Reset()
			tempKoma.SetGeoMetry(math.Pi)
			targetblock.KomaStack = append(targetblock.KomaStack, tempKoma)
			c.Player.KomaTai[0].Item2--
			c.Player.KomaTai[0].Item1 = koma.Koma{}
			b.Blocks[coor] = targetblock
			c.Player.State = playerstate.WAIT_TO_CLICK_CLOCK
		} else if g.ArrangementPhaseRoundCount != 1 &&
			c.Player.State == playerstate.IS_SELECT_KOMA {
			var (
				whichOneCanBeChosen []int = []int{}
				whichOne            int
				shift               float32
			)
			// random select koma
			for i, v := range c.Player.KomaTai {
				if v.Item2 > 0 {
					whichOneCanBeChosen = append(whichOneCanBeChosen, i)
				}
			}
			whichOne = rand.Intn(len(whichOneCanBeChosen))
			for {
				rndColumn = rand.Intn(9)
				rndRow = rand.Intn(3)
				coor = image.Point{X: rndColumn, Y: rndRow}
				targetblock = b.Blocks[coor]
				if !targetblock.HasSuI() || len(targetblock.KomaStack) < g.MaxLayer {
					break
				}
			}
			switch len(targetblock.KomaStack) {
			case 0:
				shift = 0
			case 1:
				shift = constant.SECOND_DAN
			case 2:
				shift = constant.THIRD_DAN
			}
			//複製出來
			tempKoma = c.Player.KomaTai[whichOneCanBeChosen[whichOne]].Item1.Clone()
			// 設定目標格的位置
			tempKoma.CurrentCoordinate.X = int(targetblock.Coordinate.X + constant.BLOCK_SIZE/2 - shift)
			tempKoma.CurrentCoordinate.Y = int(targetblock.Coordinate.Y + constant.BLOCK_SIZE/2 - shift)
			//設定目標代號
			tempKoma.SetCurrentPosition(targetblock.Name.Item1, targetblock.Name.Item2)
			// 設定文字顯示的位置
			tempKoma.SetGeoMetry(math.Pi)
			targetblock.KomaStack = append(targetblock.KomaStack, tempKoma)
			c.Player.KomaTai[whichOneCanBeChosen[whichOne]].Item2--
			if c.Player.KomaTai[whichOneCanBeChosen[whichOne]].Item2 == 0 {
				c.Player.KomaTai[whichOneCanBeChosen[whichOne]].Item1 = koma.Koma{}
			}
			b.Blocks[coor] = targetblock
			c.Player.State = playerstate.WAIT_TO_CLICK_CLOCK
		}
	}
}
