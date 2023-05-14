package cpu

import (
	"image"
	"math"

	"github.com/littletrainee/GunGiBoardGameGUI/block"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

func (c *CPU) MoveKoma(b board.Board) {
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
		cloneKoma.CurrentCoordinate.X = int(targetBlock.Coordinate.X+constant.BLOCK_SIZE/2) - shift
		cloneKoma.CurrentCoordinate.Y = int(targetBlock.Coordinate.Y+constant.BLOCK_SIZE/2) - shift
		cloneKoma.SetCurrentPosition(targetBlock.Name.Item1, targetBlock.Name.Item2)
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
		cloneLastKoma.CurrentCoordinate.X = int(targetBlock.Coordinate.X+constant.BLOCK_SIZE/2) - shift
		cloneLastKoma.CurrentCoordinate.Y = int(targetBlock.Coordinate.Y+constant.BLOCK_SIZE/2) - shift
		cloneLastKoma.SetCurrentPosition(targetBlock.Name.Item1, targetBlock.Name.Item2)
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
