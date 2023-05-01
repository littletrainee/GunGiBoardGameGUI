package gamehandler

import (
	"image"

	_color "github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

func checkTargetBlock(eachDanCanMove []image.Point, lastKoma koma.Koma, g *Game,
	currentDan int) []image.Point {
	// 設定是否已經有阻礙
	var (
		wall            bool
		confirmPosition []image.Point
	)
	// 從每個段的移動迭代每個位置
	for _, coor := range eachDanCanMove {
		// 設定目標位置
		targetBlockPosition := image.Point{
			X: lastKoma.CurrentPosition.Item1 +
				int(coor.X),
			Y: lastKoma.CurrentPosition.Item2 +
				int(coor.Y),
		}

		// 確認目標位置是否在棋盤內
		_, exists := g.Board.Blocks[targetBlockPosition]

		// 若在棋盤內
		if exists {
			// 若沒有阻隔
			if !wall {
				// 從blocks取出目標block
				tempblock := g.Board.Blocks[targetBlockPosition]
				targetlen := len(tempblock.KomaStack)
				// 目標點有駒或者目標的段數大於或等於當前的段數
				if targetlen != 0 || targetlen >= currentDan {
					wall = true
				}

				// if tempblock.HasSuI() { // 目標block有帥
				// 	tempblock.CurrentColor = _color.DenyColor
				// } else if targetlen < currentDan { //未超過未超過標block的段數小於當前的段數
				// 	tempblock.CurrentColor = _color.ConfirmColor
				// 	confirmPosition = append(confirmPosition, targetBlockPosition)
				// } else if targetlen == currentDan+1 {
				// 	tempblock.CurrentColor = _color.ConfirmColor
				// 	confirmPosition = append(confirmPosition, targetBlockPosition)

				// 目標block的段數小魚或等於當前段數，並且未超過 最高段的限制
				// } else if targetlen <= currentDan && targetlen < g.GameState.MaxLayer && !tempblock.HasSuI(){
				if targetlen <= currentDan && targetlen < g.GameState.MaxLayer && !tempblock.HasSuI() {
					tempblock.CurrentColor = _color.ConfirmColor
					confirmPosition = append(confirmPosition, targetBlockPosition)
				} else {
					tempblock.CurrentColor = _color.DenyColor
				}
				g.Board.Blocks[targetBlockPosition] = tempblock
			} else {
				break
			}
		}
	}
	return confirmPosition
}
