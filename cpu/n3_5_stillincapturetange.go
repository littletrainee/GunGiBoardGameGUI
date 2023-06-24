package cpu

import (
	"image"
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/action"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/GunGiBoardGameGUI/otherfunction"
)

// stillInCaptureRange 確認若帥移動到這個位置是否可以使俘獲不成立，若還是成立回傳true，否則回傳false
//
// 參數b為棋盤物件，suIPosition為帥可以移動的位置，selfColor為自家的顏色
func stillInCaptureRange(b board.Board, suIPosition image.Point, selfColor color.Gray16, g gamestate.GameState) bool {
	// 逐項尋找每個block最上面的駒可移動的位置是否不包含在帥
	for _, v := range b.Blocks {
		// 確認最上面的駒是否為對家的駒
		if len(v.KomaStack) > 0 && v.KomaStack[len(v.KomaStack)-1].Color != selfColor {
			var (
				// 定義目標block的駒數量
				stackLength int = len(v.KomaStack)
				// 定義目標block的最上方駒
				theTopOne koma.Koma = v.KomaStack[stackLength-1]
				hinder    bool
			)
			// 迭代最上方駒可能的方向
			for _, direction := range theTopOne.MoveableRange {
				hinder = false
				// 迭代允許的階層
				for i, layer := range direction {
					if i < stackLength {
						// 迭代允許的位置
						for _, pos := range layer {
							// 設定目標位置
							targetPosition := image.Point{
								X: theTopOne.CurrentPosition.X - pos.X,
								Y: theTopOne.CurrentPosition.Y + pos.Y,
							}
							// 確認目標位置是否在棋盤內
							targetBlock, ok := b.Blocks[targetPosition]
							// 若目標位置在棋盤內
							if ok && !hinder {
								// 若在棋盤內並且沒有牆阻礙
								if len(targetBlock.KomaStack) == stackLength {
									hinder = true
								} else if len(targetBlock.KomaStack) > stackLength {
									break
								}

								switch otherfunction.Move(targetBlock.KomaStack, v.KomaStack, g.LevelHolder) {
								case action.CAPTURE_ONLY, action.CAPTURE_OR_CONTROL:
									if targetPosition == suIPosition {
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
