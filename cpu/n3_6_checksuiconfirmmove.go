package cpu

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// checkSuIConfirmMove 檢查帥可移動的位置，並回傳切片
//
// 參數b為棋盤物件，suiPosition為帥的位置
func checkSuIConfirmMove(b board.Board, suiPosition image.Point) []image.Point {
	var (
		c               []image.Point
		komaStackLength int       = len(b.Blocks[suiPosition].KomaStack)
		suIKoma         koma.Koma = b.Blocks[suiPosition].KomaStack[komaStackLength-1]
	)
	for _, direction := range suIKoma.MoveableRange {
		for i, layer := range direction {
			if i < komaStackLength {
				for _, pos := range layer {
					targetPosition := image.Point{
						X: suIKoma.CurrentPosition.X - pos.X,
						Y: suIKoma.CurrentPosition.Y - pos.Y,
					}
					targetBlock, ok := b.Blocks[targetPosition]
					if ok {
						if len(targetBlock.KomaStack) == 0 {
							c = append(c, targetPosition)
						} else {
							break
						}
					} else {
						break
					}
				}
			}
		}
	}
	return c
}
