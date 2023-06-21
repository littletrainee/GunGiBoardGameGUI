package cpu

import (
	"fmt"
	"image"
	"math"

	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/cpuselect"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// DuelingPhaseMoveKoma 對弈期間的移駒，俘獲、迴避、新、嘗試俘獲與其他移動的選擇判斷要如何動作
//
// 參數b為棋盤物件
func (c *CPU) DuelingPhaseMoveKoma(b *board.Board) {
	switch c.Select {
	case cpuselect.DEFENSE_CAPTURE:
		var (
			previousBlockPosition image.Point = image.Point{X: c.MoveToTarget[0], Y: c.MoveToTarget[1]}
			targetBlockPosition   image.Point = image.Point{X: c.MoveToTarget[2], Y: c.MoveToTarget[3]}
			previousBlock         block.Block = b.Blocks[previousBlockPosition]
			targetBlock           block.Block = b.Blocks[targetBlockPosition]
			shift                 int
			cloneKoma             koma.Koma
			tempKomaSlice         []koma.Koma
		)

		// 嘗試移除非自家的駒
		for _, v := range targetBlock.KomaStack {
			if v.Color == c.SelfColor {
				// 確認偏移量
				shift = block.Shift(tempKomaSlice)
				// 複製目標駒
				cloneKoma = v.Clone()
				// 設定目標駒座標與位置
				cloneKoma.SetCurrentCoordinate(targetBlock.Coordinate, shift)
				cloneKoma.SetCurrentPosition(targetBlock.Name)
				// 設定是否旋轉
				cloneKoma.SetGeoMetry(math.Pi)
				// 將目標駒加至堆疊
				tempKomaSlice = append(tempKomaSlice, cloneKoma)
			}
		}

		// 將先前位置堆疊的最上層駒增加至目標位置的堆疊
		// 確認偏移量
		shift = block.Shift(tempKomaSlice)
		cloneKoma = previousBlock.KomaStack[len(previousBlock.KomaStack)-1].Clone()
		cloneKoma.SetCurrentCoordinate(targetBlock.Coordinate, shift)
		cloneKoma.SetCurrentPosition(targetBlock.Name)
		cloneKoma.SetGeoMetry(math.Pi)
		tempKomaSlice = append(tempKomaSlice, cloneKoma)

		targetBlock.KomaStack = tempKomaSlice
		previousBlock.KomaStack = previousBlock.KomaStack[:len(previousBlock.KomaStack)-1]
		b.Blocks[previousBlockPosition] = previousBlock
		b.Blocks[targetBlockPosition] = targetBlock
		c.MoveToTarget = nil
		c.Select = cpuselect.None
	case cpuselect.DEFENSE_AVOID:
		var (
			previousBlockPosition image.Point = image.Point{X: c.MoveToTarget[0], Y: c.MoveToTarget[1]}
			targetBlockPosition   image.Point = image.Point{X: c.MoveToTarget[2], Y: c.MoveToTarget[3]}
			previousBlock         block.Block = b.Blocks[previousBlockPosition]
			targetBlock           block.Block = b.Blocks[targetBlockPosition]
			shift                 int
			cloneKoma             koma.Koma
			tempKomaSlice         []koma.Koma
		)

		// 確認偏移量
		shift = block.Shift(tempKomaSlice)
		cloneKoma = previousBlock.KomaStack[len(previousBlock.KomaStack)-1].Clone()
		cloneKoma.SetCurrentCoordinate(targetBlock.Coordinate, shift)
		cloneKoma.SetCurrentPosition(targetBlock.Name)
		cloneKoma.SetGeoMetry(math.Pi)
		tempKomaSlice = append(tempKomaSlice, cloneKoma)

		targetBlock.KomaStack = tempKomaSlice
		previousBlock.KomaStack = previousBlock.KomaStack[:len(previousBlock.KomaStack)-1]
		b.Blocks[previousBlockPosition] = previousBlock
		b.Blocks[targetBlockPosition] = targetBlock
		c.MoveToTarget = nil
		c.Select = cpuselect.None
	case cpuselect.DEFENSE_ARATA:
		var (
			targetBlockPosision image.Point = image.Point{X: c.MoveToTarget[1], Y: c.MoveToTarget[2]}
			targetBlock         block.Block = b.Blocks[targetBlockPosision]
			shift               int
			cloneKoma           koma.Koma
			tempKomaSlice       []koma.Koma = b.Blocks[targetBlockPosision].KomaStack
		)
		shift = block.Shift(tempKomaSlice)
		cloneKoma = c.KomaDai[c.MoveToTarget[0]].Item1.Clone()
		cloneKoma.SetCurrentCoordinate(targetBlock.Coordinate, shift)
		cloneKoma.SetCurrentPosition(targetBlock.Name)
		cloneKoma.SetGeoMetry(math.Pi)

		tempKomaSlice = append(tempKomaSlice, cloneKoma)
		c.KomaDai[c.MoveToTarget[0]].Item2--
		if c.KomaDai[c.MoveToTarget[0]].Item2 == 0 {
			c.KomaDai[c.MoveToTarget[0]].Item1 = koma.Koma{}
		}

		targetBlock.KomaStack = tempKomaSlice
		b.Blocks[targetBlockPosision] = targetBlock
		c.MoveToTarget = nil
		c.Select = cpuselect.None
	case cpuselect.TRY_CAPTURE:
		if len(c.MoveToTarget) > 0 {
			if len(c.MoveToTarget) != 4 {
				fmt.Println("check")
			}
			var (
				previousBlockPosition image.Point = image.Point{X: c.MoveToTarget[0], Y: c.MoveToTarget[1]}
				targetBlockPosition   image.Point = image.Point{X: c.MoveToTarget[2], Y: c.MoveToTarget[3]}
				previousBlock         block.Block = b.Blocks[previousBlockPosition]
				targetBlock           block.Block = b.Blocks[targetBlockPosition]
				shift                 int
				cloneKoma             koma.Koma
				tempKomaSlice         []koma.Koma
			)

			// 嘗試移除非自家的駒
			for _, v := range targetBlock.KomaStack {
				if v.Color == c.SelfColor {
					// 確認偏移量
					shift = block.Shift(tempKomaSlice)
					// 複製目標駒
					cloneKoma = v.Clone()
					// 設定目標駒座標與位置
					cloneKoma.SetCurrentCoordinate(targetBlock.Coordinate, shift)
					cloneKoma.SetCurrentPosition(targetBlock.Name)
					// 設定是否旋轉
					cloneKoma.SetGeoMetry(math.Pi)
					// 將目標駒加至堆疊
					tempKomaSlice = append(tempKomaSlice, cloneKoma)
				}
			}

			// 將先前位置堆疊的最上層駒增加至目標位置的堆疊
			// 確認偏移量
			shift = block.Shift(tempKomaSlice)
			cloneKoma = previousBlock.KomaStack[len(previousBlock.KomaStack)-1].Clone()
			cloneKoma.SetCurrentCoordinate(targetBlock.Coordinate, shift)
			cloneKoma.SetCurrentPosition(targetBlock.Name)
			cloneKoma.SetGeoMetry(math.Pi)
			tempKomaSlice = append(tempKomaSlice, cloneKoma)

			targetBlock.KomaStack = tempKomaSlice
			previousBlock.KomaStack = previousBlock.KomaStack[:len(previousBlock.KomaStack)-1]
			b.Blocks[previousBlockPosition] = previousBlock
			b.Blocks[targetBlockPosition] = targetBlock
			c.MoveToTarget = nil
		}
	case cpuselect.None:
		var (
			targetPosition image.Point // 設定目標位置的座標
			targetBlock    block.Block // 複製目標block
		)
		if len(c.MoveToTarget) == 3 {
			// 設定物位置與複製目標block
			targetPosition = image.Point{X: c.MoveToTarget[1], Y: c.MoveToTarget[2]}
			targetBlock = b.Blocks[targetPosition]
			// 從駒台上複製駒
			cloneKoma := c.KomaDai[c.MoveToTarget[0]].Item1.Clone()

			// 設定複製的駒位置
			cloneKoma.SetCurrentCoordinate(targetBlock.Coordinate, block.Shift(targetBlock.KomaStack))
			cloneKoma.SetCurrentPosition(targetBlock.Name)
			cloneKoma.SetGeoMetry(math.Pi)

			// 目標block堆疊增加複製的駒
			targetBlock.KomaStack = append(targetBlock.KomaStack, cloneKoma)
			// 將駒台上的目標駒數量減少
			c.KomaDai[c.MoveToTarget[0]].Item2--
			// 若目標駒的數量為0時，則重設為空struct已減少記憶體使用量
			if c.KomaDai[c.MoveToTarget[0]].Item2 == 0 {
				c.KomaDai[c.MoveToTarget[0]].Item1 = koma.Koma{}
			}
			// 將修改後的目標block賦予回去map
			b.Blocks[targetPosition] = targetBlock
			c.MoveToTarget = nil
		} else if len(c.MoveToTarget) == 4 {
			// 設定要移動的駒的原本block位置
			previousPosition := image.Point{X: c.MoveToTarget[0], Y: c.MoveToTarget[1]}
			// 複製要移動駒的原本block
			previousBlock := b.Blocks[previousPosition]
			targetPosition = image.Point{X: c.MoveToTarget[2], Y: c.MoveToTarget[3]}
			targetBlock = b.Blocks[targetPosition]

			// 複製最後一個駒並且修改其位置
			cloneLastKoma := previousBlock.KomaStack[len(previousBlock.KomaStack)-1].Clone()
			cloneLastKoma.SetCurrentCoordinate(targetBlock.Coordinate, block.Shift(targetBlock.KomaStack))
			cloneLastKoma.SetCurrentPosition(targetBlock.Name)
			cloneLastKoma.SetGeoMetry(math.Pi)

			// 目標block堆疊增加複製的駒
			targetBlock.KomaStack = append(targetBlock.KomaStack, cloneLastKoma)
			// 將原本block的KomaStack移除最後一個駒
			previousBlock.KomaStack = previousBlock.KomaStack[:len(previousBlock.KomaStack)-1]
			// 並賦予回去map
			b.Blocks[previousPosition] = previousBlock
			// 將修改後的目標block賦予回去map
			b.Blocks[targetPosition] = targetBlock
			c.MoveToTarget = nil
		}
	default:
	}
}
