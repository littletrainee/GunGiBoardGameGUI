package cpu

import (
	"image"
	"math"
	"math/rand"

	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
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
			tempKoma.SetCurrentCoordinate(targetblock.Coordinate, 0)
			tempKoma.SetCurrentPosition(targetblock.Name)
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
				shift               int
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
			shift = block.Shift(targetblock.KomaStack)
			//複製出來
			tempKoma = c.Player.KomaTai[whichOneCanBeChosen[whichOne]].Item1.Clone()
			// 設定目標格的位置
			tempKoma.SetCurrentCoordinate(targetblock.Coordinate, shift)
			//設定目標代號
			tempKoma.SetCurrentPosition(targetblock.Name)
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
