package cpu

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// invitableLoseOpportunity 必輸的機會，判斷是否被對方將軍自己的帥，若被將軍則回傳true，否則回傳false
//
// 參數b為棋盤物件
func (c *CPU) inevitableLoseOpportunity(b board.Board) bool {
	// 迭代每個block
	for k, v := range b.Blocks {
		var (
			// 可移動位置的目標駒
			lastOne koma.Koma
			hinder  bool
		)
		// 若有駒且最上面的駒不是己方的
		if len(v.KomaStack) > 0 && v.KomaStack[len(v.KomaStack)-1].Color != c.SelfColor {
			// 當前迭代的block最上方的駒
			var theTopOne koma.Koma = v.KomaStack[len(v.KomaStack)-1]
			// 迭代可能的方向
			for _, direction := range theTopOne.MoveableRange {
				hinder = false
				// 基於可能的方向去進行迭代每個方向的段
				for i, layer := range direction {
					if len(v.KomaStack) > i {
						// 基於每個段迭代位置
						for _, pos := range layer {
							// 目標駒可以移動的位置
							var confirmPosition image.Point = image.Point{
								X: theTopOne.CurrentPosition.X - pos.X,
								Y: theTopOne.CurrentPosition.Y + pos.Y,
							}
							// 確認目標位置是否在棋盤內
							targetBlock, ok := b.Blocks[confirmPosition]
							if ok {
								if !hinder {
									// 在棋盤內確認是否有阻擋已中斷這個方向的迭代
									if len(targetBlock.KomaStack) > 0 {
										if len(targetBlock.KomaStack) >= len(v.KomaStack) {
											hinder = true
										}
										// 最後一個駒是否是對家的帥
										lastOne = targetBlock.KomaStack[len(targetBlock.KomaStack)-1]
										if lastOne.Name == "帥" && lastOne.Color == c.SelfColor {
											// 前兩個代表哪邊可以將軍的座標，後兩個代表帥的座標
											c.MoveToTarget = []int{k.X, k.Y, confirmPosition.X, confirmPosition.Y}
											return true
										}
									}
								} else {
									break
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
