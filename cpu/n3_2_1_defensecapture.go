package cpu

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/action"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/levelholder"
	"github.com/littletrainee/GunGiBoardGameGUI/otherfunction"
)

// defenseCapture 俘獲防禦，可以導出那個位置的駒可以俘獲對方的駒從而避免被將軍的狀況，若可以透過俘獲解除將軍則回傳true，否則回傳false
//
// 參數b為棋盤物件
func (c *CPU) defenseCapture(b board.Board, levelHolder levelholder.LevelHolder) bool {
	// 迭代檢查每個Block
	for k, v := range b.Blocks {
		// 有包含駒，並且顏色不是自家顏色
		if len(v.KomaStack) > 0 && v.KomaStack[len(v.KomaStack)-1].Color == c.SelfColor {
			//宣告
			var (
				whereCanCheckmate    image.Point = image.Point{X: c.MoveToTarget[0], Y: c.MoveToTarget[1]}
				targetPBlockPosition image.Point
				hinder               bool
			)
			// 迭代目標駒的可移動範圍所有方向
			for _, direction := range v.KomaStack[len(v.KomaStack)-1].MoveableRange {
				hinder = false
				// 迭代每個方向的各個階段
				for i, layer := range direction {
					// 確認是否為當前駒可以移動的階段範圍
					if i < len(v.KomaStack) {
						// 迭代每個階段的所有位置
						for _, pos := range layer {
							// 定義位置
							targetPBlockPosition = image.Point{
								X: v.KomaStack[len(v.KomaStack)-1].CurrentPosition.X - pos.X,
								Y: v.KomaStack[len(v.KomaStack)-1].CurrentPosition.Y - pos.Y,
							}
							// 確認是否存在於棋盤內
							targetBlock, ok := b.Blocks[targetPBlockPosition]
							// 若存在於棋盤內
							if ok && !hinder {
								// 若目標的段數大於當前的段數，中斷迭代
								if len(targetBlock.KomaStack) > len(v.KomaStack) {
									break
								} else if len(targetBlock.KomaStack) == len(v.KomaStack) {
									hinder = true
								}
								switch otherfunction.Move(targetBlock.KomaStack, v.KomaStack, levelHolder) {
								case action.CAPTURE_ONLY, action.CAPTURE_OR_CONTROL:
									if whereCanCheckmate == targetPBlockPosition {
										// 目標位置的自家駒可以俘獲可以將軍自家帥的位置的駒
										c.MoveToTarget = []int{k.X, k.Y, c.MoveToTarget[0], c.MoveToTarget[1]}
										return true
									}

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
