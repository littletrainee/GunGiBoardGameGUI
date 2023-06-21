package cpu

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// inevitableWinOpportunity 必勝的機會，判斷是否可以將軍對方的帥，若可以將軍對方的帥則回傳true，否則false
//
// 參數b為棋盤物件
func (c *CPU) inevitableWinOpportunity(b board.Board) bool {
	// 迭代每個block
	for _, v := range b.Blocks {
		var (
			komaStackLength int = len(v.KomaStack)
			// 當前迭代的block最上方的駒
			theTopOneOfKomaStack koma.Koma
			// 目標駒可以移動的位置
			confirmPosition image.Point
			// 可移動位置的目標駒
			lastOne koma.Koma
			hinder  bool
		)
		// 若有駒且最上面的駒是己方的
		if komaStackLength > 0 && v.KomaStack[komaStackLength-1].Color == c.SelfColor {
			theTopOneOfKomaStack = v.KomaStack[komaStackLength-1]
			// 迭代可能的方向
			for _, direction := range theTopOneOfKomaStack.ProbablyMoveing {
				hinder = false
				// 基於可能的方向去進行迭代每個方向的段
				for _, layer := range direction {
					// 基於每個段迭代位置
					for _, pos := range layer {
						// 迭代的每個位置
						confirmPosition = image.Point{
							X: theTopOneOfKomaStack.CurrentPosition.X - pos.X,
							Y: theTopOneOfKomaStack.CurrentPosition.Y - pos.Y,
						}
						// 確認目標位置是否在棋盤內
						targetBlock, ok := b.Blocks[confirmPosition]
						if ok {
							// 在棋盤內確認是否有阻擋已中斷這個方向的迭代
							if !hinder {
								// 目標block有駒
								if len(targetBlock.KomaStack) != 0 {
									// 有駒並且高度大於當前block
									if len(targetBlock.KomaStack) >= komaStackLength {
										hinder = true
									}
									// 最後一個駒是否是對家的帥
									lastOne = targetBlock.KomaStack[len(targetBlock.KomaStack)-1]
									if lastOne.Name == "帥" && lastOne.Color != c.SelfColor {
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
	return false
}
