package gamehandler

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/action"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/levelholder"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/GunGiBoardGameGUI/otherfunction"
)

// otherCheck 其他駒可以移動的範圍，回傳值為可以移動的選項切片
//
// 參數direction為每個可以移動方向中的當前的段，lastKoma為當前駒物件，currentDan為當前的段，b為棋盤物件
func otherCheck(levelholder levelholder.LevelHolder, whereOther []koma.Koma, b board.Board) []image.Point {

	// 設定是否已經有阻礙
	var (
		hinder          bool
		confirmPosition []image.Point
		currentDan      int       = len(whereOther)
		other           koma.Koma = whereOther[currentDan-1]
	)
	// 迭代當前的方向
	for _, direction := range other.MoveableRange {
		for danIndex, eachDanCanMove := range direction {
			hinder = false
			if danIndex < currentDan {
				// 從每個段的移動迭代每個位置
				for _, coor := range eachDanCanMove {
					// 設定目標位置
					targetBlockPosition := image.Point{
						X: other.CurrentPosition.X - coor.X,
						Y: other.CurrentPosition.Y + coor.Y,
					}

					// 確認目標位置是否在棋盤內
					targetBlock, exists := b.Blocks[targetBlockPosition]

					// 若在棋盤內且沒有阻隔
					if exists && !hinder {
						// 從blocks取出目標block
						var targetlen int = len(targetBlock.KomaStack)

						if targetlen > currentDan { // 目標段數高於本身段數則中斷
							break
						} else if targetlen == currentDan { // 否則將hinder設為true
							hinder = true
						}

						switch otherfunction.Move(targetBlock.KomaStack, whereOther, levelholder) {
						case action.MOVE, action.CAPTURE_OR_CONTROL:
							confirmPosition = append(confirmPosition, targetBlockPosition)
							targetBlock.CurrentColor = color.ConfirmColor
						case action.CAPTURE_ONLY:
							confirmPosition = append(confirmPosition, targetBlockPosition)
							targetBlock.CurrentColor = color.CaptureColor
						default:
							targetBlock.CurrentColor = color.DenyColor
						}
						b.Blocks[targetBlockPosition] = targetBlock
					} else {
						break
					}
				}
			} else {
				break
			}
		}
	}
	return confirmPosition
}
