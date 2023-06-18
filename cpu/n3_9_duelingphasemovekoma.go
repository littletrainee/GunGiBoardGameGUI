package cpu

import (
	"image"
	"math"

	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/cpuselect"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

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
		c.checkmateBy = image.Point{}
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
		c.checkmateBy = image.Point{}
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

		targetBlock.KomaStack = tempKomaSlice
		b.Blocks[targetBlockPosision] = targetBlock
		c.MoveToTarget = nil
		c.Select = cpuselect.None
		c.checkmateBy = image.Point{}
	case cpuselect.TRY_CAPTURE:
		if len(c.MoveToTarget) > 0 {
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
			c.checkmateBy = image.Point{}
		}
	default:
		var (
			// 複製目標block
			targetBlock block.Block = b.Blocks[c.targetPosition]
			// 設定目標偏移量
			shift int = block.Shift(targetBlock.KomaStack)
		)
		if len(c.targetKoma) == 1 {
			// 從駒台上複製駒
			cloneKoma := c.KomaDai[c.targetKoma[0]].Item1.Clone()

			// 設定複製的駒位置
			cloneKoma.SetCurrentCoordinate(targetBlock.Coordinate, shift)
			cloneKoma.SetCurrentPosition(targetBlock.Name)
			cloneKoma.SetGeoMetry(math.Pi)

			// 目標block堆疊增加複製的駒
			targetBlock.KomaStack = append(targetBlock.KomaStack, cloneKoma)
			// 將駒台上的目標駒數量減少
			c.KomaDai[c.targetKoma[0]].Item2--
			// 若目標駒的數量為0時，則重設為空struct已減少記憶體使用量
			if c.KomaDai[c.targetKoma[0]].Item2 == 0 {
				c.KomaDai[c.targetKoma[0]].Item1 = koma.Koma{}
			}
		} else {
			// 複製要移動駒的原本block
			orginalBlock := b.Blocks[image.Point{X: c.targetKoma[0], Y: c.targetKoma[1]}]

			// 複製最後一個駒並且修改其位置
			cloneLastKoma := orginalBlock.KomaStack[len(orginalBlock.KomaStack)-1].Clone()
			cloneLastKoma.SetCurrentCoordinate(targetBlock.Coordinate, shift)
			cloneLastKoma.SetCurrentPosition(targetBlock.Name)
			cloneLastKoma.SetGeoMetry(math.Pi)

			// 目標block堆疊增加複製的駒
			targetBlock.KomaStack = append(targetBlock.KomaStack, cloneLastKoma)
			// 將原本block的KomaStack移除最後一個駒
			orginalBlock.KomaStack = orginalBlock.KomaStack[:len(orginalBlock.KomaStack)-1]
			// 並賦予回去map
			b.Blocks[image.Point{X: c.targetKoma[0], Y: c.targetKoma[1]}] = orginalBlock
		}
		// 將修改後的目標block賦予回去map
		b.Blocks[c.targetPosition] = targetBlock
		c.targetKoma = nil
		c.targetPosition = image.Point{}

	}
}
