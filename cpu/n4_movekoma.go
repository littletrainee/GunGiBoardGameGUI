package cpu

import (
	"image"
	"math"

	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

func (c *CPU) MoveKoma(b board.Board) {
	if len(c.TargetMove) != 0 {
		var (
			previousBlock block.Block = b.Blocks[c.TargetMove[0]]
			targetBlock   block.Block = b.Blocks[c.TargetMove[1]]
			shift         int
			cloneKoma     koma.Koma
			tempKomaSlice []koma.Koma
		)

		for _, v := range targetBlock.KomaStack {
			if v.Color == c.SelfColor {
				shift = block.Shift(tempKomaSlice)
				cloneKoma = v.Clone()
				cloneKoma.SetCurrentCoordinate(targetBlock.Coordinate, shift)
				cloneKoma.SetCurrentPosition(targetBlock.Name)
				cloneKoma.SetGeoMetry(math.Pi)
				tempKomaSlice = append(tempKomaSlice, cloneKoma)
			}
		}
		shift = block.Shift(tempKomaSlice)
		cloneKoma = previousBlock.KomaStack[len(previousBlock.KomaStack)-1].Clone()
		cloneKoma.SetCurrentCoordinate(targetBlock.Coordinate, shift)
		cloneKoma.SetCurrentPosition(targetBlock.Name)
		cloneKoma.SetGeoMetry(math.Pi)
		tempKomaSlice = append(tempKomaSlice, cloneKoma)

		targetBlock.KomaStack = tempKomaSlice
		previousBlock.KomaStack = previousBlock.KomaStack[:len(previousBlock.KomaStack)-1]
		b.Blocks[c.TargetMove[0]] = previousBlock
		b.Blocks[c.TargetMove[1]] = targetBlock
		c.TargetMove = nil
	} else {
		var (
			// 複製目標block
			targetBlock block.Block = b.Blocks[c.targetPosition]
			// 設定目標偏移量
			shift int = block.Shift(targetBlock.KomaStack)
		)
		if len(c.targetKoma) == 1 {
			// 從駒台上複製駒
			cloneKoma := c.KomaTai[c.targetKoma[0]].Item1.Clone()

			// 設定複製的駒位置
			cloneKoma.SetCurrentCoordinate(targetBlock.Coordinate, shift)
			cloneKoma.SetCurrentPosition(targetBlock.Name)
			cloneKoma.SetGeoMetry(math.Pi)

			// 目標block堆疊增加複製的駒
			targetBlock.KomaStack = append(targetBlock.KomaStack, cloneKoma)
			// 將駒台上的目標駒數量減少
			c.KomaTai[c.targetKoma[0]].Item2--
			// 若目標駒的數量為0時，則重設為空struct已漸少記憶體使用量
			if c.KomaTai[c.targetKoma[0]].Item2 == 0 {
				c.KomaTai[c.targetKoma[0]].Item1 = koma.Koma{}
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
