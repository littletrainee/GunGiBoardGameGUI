package cpu

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
)

// tryCapture 嘗試俘獲，嘗試是否可以俘獲對方的駒，若可以則回傳true，否則回傳false
//
// 參數b為棋盤物件
func (c *CPU) tryCapture(b board.Board) bool {
	// 迭代檢查每個Block
	for k, v := range b.Blocks {
		// 有包含駒，並且顏色不是自家顏色
		if len(v.KomaStack) > 0 && v.KomaStack[len(v.KomaStack)-1].Color == c.SelfColor {
			//宣告
			var (
				tempPos image.Point
				hinder  bool
			)
			// 迭代目標駒的可移動範圍所有方向
			for _, direction := range v.KomaStack[len(v.KomaStack)-1].ProbablyMoveing {
				hinder = false
				// 迭代每個方向的各個階段
				for i, layer := range direction {
					// 確認是否為當前駒可以移動的階段範圍
					if i < len(v.KomaStack) {
						// 迭代每個階段的所有位置
						for _, pos := range layer {
							if !hinder {
								// 定義位置
								tempPos = image.Point{
									X: v.KomaStack[len(v.KomaStack)-1].CurrentPosition.X - pos.X,
									Y: v.KomaStack[len(v.KomaStack)-1].CurrentPosition.Y - pos.Y,
								}
								// 確認是否存在於棋盤內
								targetBlock, ok := b.Blocks[tempPos]
								// 若存在於棋盤內
								if ok {
									// 若目標的段數大於當前的段數，中斷迭代
									if len(targetBlock.KomaStack) > len(v.KomaStack) {
										break
									} else if len(targetBlock.KomaStack) == len(v.KomaStack) {
										hinder = true
									}
									if len(targetBlock.KomaStack) > 0 && targetBlock.KomaStack[len(targetBlock.KomaStack)-1].Color != c.SelfColor {
										c.MoveToTarget = []int{k.X, k.Y, tempPos.X, tempPos.Y}
										return true
									}
								} else {
									break
								}
							} else {
								break
							}
						}
					} else {
						break
					}
				}
			}
		}
	}
	return false
}
