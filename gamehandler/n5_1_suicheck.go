package gamehandler

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/action"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/levelholder"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/GunGiBoardGameGUI/otherfunction"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
)

// suICheck 在入門與初級階級下，率可以移動的範圍，必須不包含任何駒，回傳值為可以移動的選項切片
//
// 參數eachCanCanMove為每個可以移動方向中的當前段可以移動的範圍，suI為帥物件，b為棋盤物件，currentDan為當前的段
func suICheck(levelholder levelholder.LevelHolder, whereSuI []koma.Koma, b board.Board) []image.Point {
	// 設定是否已經有阻礙
	var (
		hinder          bool
		confirmPosition []image.Point
		currentDan      int       = len(whereSuI)
		suI             koma.Koma = whereSuI[currentDan-1]
	)
	// 若階級是入門或是初階
	if levelholder.CurrentLevel == level.BEGINNER || levelholder.CurrentLevel == level.ELEMENTARY {
		for _, direction := range suI.MoveableRange {
			for suIIndex, eachDanCanMove := range direction {
				if suIIndex < currentDan {
					// 從每個段的移動迭代每個位置
					for _, coor := range eachDanCanMove {
						// 設定目標位置
						targetBlockPosition := image.Point{
							X: suI.CurrentPosition.X + coor.X,
							Y: suI.CurrentPosition.Y + coor.Y,
						}

						// 從blocks取出目標block，並確認目標位置是否在棋盤內
						targetBlock, exists := b.Blocks[targetBlockPosition]

						// 若在棋盤內
						if exists {

							a := otherfunction.Move(targetBlock.KomaStack, whereSuI, levelholder)
							switch a {
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
	} else if levelholder.CurrentLevel == level.INTERMEDIATE || levelholder.CurrentLevel == level.ADVANCED {
		for _, direction := range suI.MoveableRange {
			hinder = false
			for suIIndex, eachDanCanMove := range direction {
				if suIIndex < currentDan {
					// 從每個段的移動迭代每個位置
					for _, coor := range eachDanCanMove {
						// 設定目標位置
						targetBlockPosition := image.Point{
							X: suI.CurrentPosition.X + coor.X,
							Y: suI.CurrentPosition.Y + coor.Y,
						}

						// 從blocks取出目標block，確認目標位置是否在棋盤內
						targetBlock, exists := b.Blocks[targetBlockPosition]

						// 若在棋盤內且沒有阻隔
						if exists && !hinder {
							targetlen := len(targetBlock.KomaStack)

							if targetlen > currentDan {
								break
							} else if targetlen == currentDan {
								hinder = true
							}

							a := otherfunction.Move(targetBlock.KomaStack, whereSuI, levelholder)
							switch a {
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
	}
	return confirmPosition

}
