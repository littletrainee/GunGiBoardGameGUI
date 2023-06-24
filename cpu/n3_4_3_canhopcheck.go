package cpu

import (
	"errors"
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/action"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/levelholder"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/GunGiBoardGameGUI/otherfunction"
)

// canHopCheck 其他駒可以移動的範圍，回傳值為可以移動的選項切片
//
// 參數levelholder為階級控制物件，whereHop為可跳躍居的位置區塊切片，b為棋盤物件
func canHopCheck(levelholder levelholder.LevelHolder, whereHop []koma.Koma, b board.Board) ([]image.Point, error) {
	var (
		hinder          bool // 設定是否已經有阻礙
		confirmPosition []image.Point
		currentDan      int       = len(whereHop)
		hop             koma.Koma = whereHop[currentDan-1]
		err             error
	)
	switch hop.Name {
	case "弓":
		for i, direction := range hop.MoveableRange {
			switch i {
			case 0:
				temp := image.Point{X: hop.CurrentPosition.X, Y: hop.CurrentPosition.Y - 1}
				if len(b.Blocks[temp].KomaStack) > currentDan {
					tempblock := b.Blocks[temp]
					tempblock.CurrentColor = color.DenyColor
					b.Blocks[temp] = tempblock
					continue
				}
			case 1:
				temp := image.Point{X: hop.CurrentPosition.X - 1, Y: hop.CurrentPosition.Y - 1}
				if len(b.Blocks[temp].KomaStack) > currentDan {
					tempblock := b.Blocks[temp]
					tempblock.CurrentColor = color.DenyColor
					b.Blocks[temp] = tempblock
					continue
				}
			case 7:
				temp := image.Point{X: hop.CurrentPosition.X + 1, Y: hop.CurrentPosition.Y - 1}
				if len(b.Blocks[temp].KomaStack) > currentDan {
					tempblock := b.Blocks[temp]
					tempblock.CurrentColor = color.DenyColor
					b.Blocks[temp] = tempblock
					continue
				}
			}
			for danIndex, eachDanCanMove := range direction {
				hinder = false
				if danIndex < currentDan {
					// 從每個段的移動迭代每個位置
					for _, coor := range eachDanCanMove {
						// 設定目標位置
						targetBlockPosition := image.Point{
							X: hop.CurrentPosition.X - coor.X,
							Y: hop.CurrentPosition.Y - coor.Y,
						}

						// 確認目標位置是否在棋盤內
						targetBlock, exists := b.Blocks[targetBlockPosition]

						// 若在棋盤內且沒有阻隔
						if exists && !hinder {
							// 從blocks取出目標block
							targetBlockLen := len(targetBlock.KomaStack)

							if targetBlockLen > currentDan { // 目標段數高於本身段數則中斷
								break
							} else if targetBlockLen == currentDan { // 否則將hinder設為true
								hinder = true
							}

							switch otherfunction.Move(targetBlock.KomaStack, whereHop, levelholder) {
							case action.MOVE, action.CAPTURE_ONLY, action.CAPTURE_OR_CONTROL:
								if !targetBlock.HasSuI() {
									confirmPosition = append(confirmPosition, targetBlockPosition)
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
	case "砲":
		for i, direction := range hop.MoveableRange {
			if i == 0 {
				temp := image.Point{X: hop.CurrentPosition.X, Y: hop.CurrentPosition.Y - 1}
				if len(b.Blocks[temp].KomaStack) > currentDan {
					tempblock := b.Blocks[temp]
					tempblock.CurrentColor = color.DenyColor
					b.Blocks[temp] = tempblock
					continue
				}
				temp = image.Point{X: hop.CurrentPosition.X, Y: hop.CurrentPosition.Y - 2}
				if len(b.Blocks[temp].KomaStack) > currentDan {
					tempblock := b.Blocks[temp]
					tempblock.CurrentColor = color.DenyColor
					b.Blocks[temp] = tempblock
					continue
				}
			}
			for danIndex, eachDanCanMove := range direction {
				hinder = false
				if danIndex < currentDan {
					// 從每個段的移動迭代每個位置
					for _, coor := range eachDanCanMove {
						// 設定目標位置
						targetBlockPosition := image.Point{
							X: hop.CurrentPosition.X - coor.X,
							Y: hop.CurrentPosition.Y - coor.Y,
						}

						// 確認目標位置是否在棋盤內
						targetBlock, exists := b.Blocks[targetBlockPosition]

						// 若在棋盤內且沒有阻隔
						if exists && !hinder {
							// 從blocks取出目標block
							targetBlockLen := len(targetBlock.KomaStack)

							if targetBlockLen > currentDan { // 目標段數高於本身段數則中斷
								break
							} else if targetBlockLen == currentDan { // 否則將hinder設為true
								hinder = true
							}

							switch otherfunction.Move(targetBlock.KomaStack, whereHop, levelholder) {
							case action.MOVE, action.CAPTURE_ONLY, action.CAPTURE_OR_CONTROL:
								if !targetBlock.HasSuI() {
									confirmPosition = append(confirmPosition, targetBlockPosition)
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

	case "筒":
		for i, direction := range hop.MoveableRange {
			if i == 0 {
				temp := image.Point{X: hop.CurrentPosition.X, Y: hop.CurrentPosition.Y - 1}
				if len(b.Blocks[temp].KomaStack) > currentDan {
					tempblock := b.Blocks[temp]
					tempblock.CurrentColor = color.DenyColor
					b.Blocks[temp] = tempblock
					continue
				}
			}
			for danIndex, eachDanCanMove := range direction {
				hinder = false
				if danIndex < currentDan {
					// 從每個段的移動迭代每個位置
					for _, coor := range eachDanCanMove {
						// 設定目標位置
						targetBlockPosition := image.Point{
							X: hop.CurrentPosition.X - coor.X,
							Y: hop.CurrentPosition.Y - coor.Y,
						}

						// 確認目標位置是否在棋盤內
						targetBlock, exists := b.Blocks[targetBlockPosition]

						// 若在棋盤內且沒有阻隔
						if exists && !hinder {
							// 從blocks取出目標block
							targetBlockLen := len(targetBlock.KomaStack)

							if targetBlockLen > currentDan { // 目標段數高於本身段數則中斷
								break
							} else if targetBlockLen == currentDan { // 否則將hinder設為true
								hinder = true
							}

							switch otherfunction.Move(targetBlock.KomaStack, whereHop, levelholder) {
							case action.MOVE, action.CAPTURE_ONLY, action.CAPTURE_OR_CONTROL:
								if !targetBlock.HasSuI() {
									confirmPosition = append(confirmPosition, targetBlockPosition)
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
	if len(confirmPosition) == 0 {
		err = errors.New("canHopCheck的confirmPosition是空的")
	}
	return confirmPosition, err
}
