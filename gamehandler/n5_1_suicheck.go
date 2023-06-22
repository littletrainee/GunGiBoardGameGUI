package gamehandler

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// suICheck 在入門與初級階級下，率可以移動的範圍，必須不包含任何駒，回傳值為可以移動的選項切片
//
// 參數eachCanCanMove為每個可以移動方向中的當前段可以移動的範圍，suI為帥物件，b為棋盤物件，currentDan為當前的段
func suICheck(eachDanCanMove []image.Point, suI koma.Koma, b board.Board, currentDan int) []image.Point {
	// 設定是否已經有阻礙
	var (
		hinder          bool
		confirmPosition []image.Point
	)
	// 從每個段的移動迭代每個位置
	for _, coor := range eachDanCanMove {
		// 設定目標位置
		targetBlockPosition := image.Point{
			X: 10 - suI.CurrentPosition.X + coor.X,
			Y: suI.CurrentPosition.Y + coor.Y,
		}

		// 確認目標位置是否在棋盤內
		_, exists := b.Blocks[targetBlockPosition]

		// 若在棋盤內
		if exists {
			// 若沒有阻隔
			if !hinder {
				// 從blocks取出目標block
				tempblock := b.Blocks[targetBlockPosition]
				targetlen := len(tempblock.KomaStack)
				// 目標點有駒或者目標的段數大於或等於當前的段數
				if currentDan == targetlen {
					if tempblock.KomaStack[targetlen-1].Color != suI.Color {
						confirmPosition = append(confirmPosition, targetBlockPosition)
						tempblock.CurrentColor = color.ConfirmColor
					} else {
						tempblock.CurrentColor = color.DenyColor
					}
				} else {
					tempblock.CurrentColor = color.ConfirmColor
					confirmPosition = append(confirmPosition, targetBlockPosition)
				}
				b.Blocks[targetBlockPosition] = tempblock
			} else {
				break
			}
		}
	}
	return confirmPosition

}
