package gamehandler

import (
	"image"

	_color "github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// otherCheck 其他駒可以移動的範圍，回傳值為可以移動的選項切片
//
// 參數direction為每個可以移動方向中的當前的段，lastKoma為當前駒物件，currentDan為當前的段，b為棋盤物件
func otherCheck(direction [][]image.Point, lastKoma koma.Koma, currentDan int, g *Game,
) []image.Point {
	// 設定是否已經有阻礙
	var (
		hinder          bool
		confirmPosition []image.Point
	)
	for danIndex, eachDanCanMove := range direction {
		if danIndex < currentDan {
			// 從每個段的移動迭代每個位置
			for _, coor := range eachDanCanMove {
				// 設定目標位置
				targetBlockPosition := image.Point{
					X: lastKoma.CurrentPosition.X - coor.X,
					Y: lastKoma.CurrentPosition.Y + coor.Y,
				}

				// 確認目標位置是否在棋盤內
				tempblock, exists := g.Board.Blocks[targetBlockPosition]

				// 若在棋盤內
				if exists {
					// 若沒有阻隔
					if !hinder {
						// 從blocks取出目標block
						var targetlen int = len(tempblock.KomaStack)

						if targetlen > currentDan { // 目標段數高於本身段數則中斷
							break
						} else if targetlen == currentDan { // 否則將hinder設為true
							hinder = true
						}

						// 目標段等於當前段則設為CaptureColorl
						if targetlen == currentDan && targetlen == g.GameState.LevelHolder.MaxLayer {
							tempblock.CurrentColor = _color.CaptureColor
						} else if targetlen <= currentDan {
							tempblock.CurrentColor = _color.ConfirmColor
						} else if targetlen > currentDan {
							tempblock.CurrentColor = _color.DenyColor
						}
						confirmPosition = append(confirmPosition, targetBlockPosition)
						g.Board.Blocks[targetBlockPosition] = tempblock
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
	return confirmPosition
}
