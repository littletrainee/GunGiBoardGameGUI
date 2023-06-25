package cpu

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/action"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/levelholder"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/GunGiBoardGameGUI/otherfunction"
)

// invitableLoseOpportunity 必輸的機會，判斷是否被對方將軍自己的帥，若被將軍則回傳true，否則回傳false
//
// 參數b為棋盤物件
func (c *CPU) inevitableLoseOpportunity(b board.Board, levelholder levelholder.LevelHolder) bool {
	// 迭代每個block
	for k, v := range b.Blocks {
		var currentLen int = len(v.KomaStack)
		// 若有駒且最上面的駒不是己方的
		if currentLen > 0 && v.KomaStack[currentLen-1].Color != c.SelfColor {
			// 當前迭代的block最上方的駒
			var theTopOne koma.Koma = v.KomaStack[currentLen-1]
			// 迭代可能的方向
			for _, direction := range theTopOne.MoveableRange {
				var hinder bool = false
				// 基於可能的方向去進行迭代每個方向的段
				for i, layer := range direction {
					if currentLen > i {
						// 基於每個段迭代位置
						for _, pos := range layer {
							// 目標駒可以移動的位置
							var targetBlockPosition image.Point = image.Point{
								X: theTopOne.CurrentPosition.X - pos.X,
								Y: theTopOne.CurrentPosition.Y + pos.Y,
							}
							// 確認目標位置是否在棋盤內
							targetBlock, ok := b.Blocks[targetBlockPosition]
							if ok && !hinder {
								var targetBlockLen int = len(targetBlock.KomaStack)
								// 若目標的段數大於當前的段數，中斷迭代
								if targetBlockLen > currentLen {
									break
								} else if targetBlockLen == currentLen {
									hinder = true
								}

								switch otherfunction.Move(targetBlock.KomaStack, v.KomaStack, levelholder) {
								case action.CAPTURE_ONLY, action.CAPTURE_OR_CONTROL:
									// 定義最後一個駒
									var lastOne koma.Koma = targetBlock.KomaStack[targetBlockLen-1]
									if lastOne.Name == "帥" && lastOne.Color == c.SelfColor {
										// 前兩個代表哪邊可以將軍的座標，後兩個代表帥的座標
										c.MoveToTarget = []int{k.X, k.Y, targetBlockPosition.X, targetBlockPosition.Y}
										return true
									}
								}
							} else { // 不存在則中斷
								break
							}
						}
					}
				}
			}

		}
	}
	return false
}
