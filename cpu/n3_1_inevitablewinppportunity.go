package cpu

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/action"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/levelholder"
	"github.com/littletrainee/GunGiBoardGameGUI/otherfunction"
)

// inevitableWinOpportunity 必勝的機會，判斷是否可以將軍對方的帥，若可以將軍對方的帥則回傳true，否則false
//
// 參數b為棋盤物件
func (c *CPU) inevitableWinOpportunity(b board.Board, l levelholder.LevelHolder) bool {
	// 迭代每個block
	for assumePosition, assumedBlock := range b.Blocks {
		var (
			assumedBlockLength int = len(assumedBlock.KomaStack)
		)
		// 若有駒且最上面的駒是己方的
		if assumedBlockLength > 0 && assumedBlock.KomaStack[assumedBlockLength-1].Color == c.SelfColor {
			assumedKoma := assumedBlock.KomaStack[assumedBlockLength-1]
			// 迭代可能的方向
			for _, direction := range assumedKoma.MoveableRange {
				hinder := false
				// 基於可能的方向去進行迭代每個方向的段
				for i, layer := range direction {
					if i < assumedBlockLength {
						// 基於每個段迭代位置
						for _, pos := range layer {
							// 迭代的每個位置
							targetPosition := image.Point{
								X: assumedKoma.CurrentPosition.X - pos.X,
								Y: assumedKoma.CurrentPosition.Y - pos.Y,
							}
							// 確認目標位置是否在棋盤內
							targetBlock, ok := b.Blocks[targetPosition]
							targetBlockLen := len(targetBlock.KomaStack)
							if ok && !hinder && targetBlockLen > 0 {
								if targetBlockLen > assumedBlockLength {
									break
								} else if targetBlockLen == assumedBlockLength {
									hinder = true
								}

								switch otherfunction.Move(targetBlock.KomaStack, assumedBlock.KomaStack, l) {
								case action.CAPTURE_ONLY, action.CAPTURE_OR_CONTROL:
									if targetBlock.HasSuI() {
										c.MoveToTarget = []int{assumePosition.X, assumePosition.Y, targetPosition.X, targetPosition.Y}
										return true
									}
								}
								// if targetKoma.Name == "帥" && targetKoma.Color != c.SelfColor {
								// 	c.MoveToTarget = []int{assumePosition.X, assumePosition.Y, targetPosition.X, targetPosition.Y}
								// 	return true
								// }
							} else { // 不存在則中斷
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
