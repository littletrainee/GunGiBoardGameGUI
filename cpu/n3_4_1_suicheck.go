package cpu

import (
	"errors"
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/action"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/levelholder"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/GunGiBoardGameGUI/otherfunction"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
)

// suICheck 檢查帥可移動的位置，並回傳切片
//
// 參數b為棋盤物件，suiPosition為帥的位置
func suICheck(levelholder levelholder.LevelHolder, whereSuI []koma.Koma, b board.Board) ([]image.Point, error) {
	var (
		hinder          bool
		confirmPosition []image.Point
		currentDan      int       = len(whereSuI)
		suI             koma.Koma = whereSuI[currentDan-1]
		err             error
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
							X: suI.CurrentPosition.X - coor.X,
							Y: suI.CurrentPosition.Y - coor.Y,
						}

						// 從blocks取出目標block，並確認目標位置是否在棋盤內
						targetBlock, exists := b.Blocks[targetBlockPosition]

						// 若在棋盤內
						if exists {
							switch otherfunction.Move(targetBlock.KomaStack, whereSuI, levelholder) {
							case action.MOVE, action.CAPTURE_ONLY:
								confirmPosition = append(confirmPosition, targetBlockPosition)
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
	} else if levelholder.CurrentLevel == level.INTERMEDIATE || levelholder.CurrentLevel == level.ADVANCED {
		for _, direction := range suI.MoveableRange {
			hinder = false
			for suIIndex, eachDanCanMove := range direction {
				if suIIndex < currentDan {
					// 從每個段的移動迭代每個位置
					for _, coor := range eachDanCanMove {
						// 設定目標位置
						targetBlockPosition := image.Point{
							X: suI.CurrentPosition.X - coor.X,
							Y: suI.CurrentPosition.Y - coor.Y,
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

							switch otherfunction.Move(targetBlock.KomaStack, whereSuI, levelholder) {
							case action.MOVE, action.CAPTURE_OR_CONTROL, action.CAPTURE_ONLY:
								confirmPosition = append(confirmPosition, targetBlockPosition)
							default:
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
		err = errors.New("suiCheck的confirmPositin 是空的")
	}
	return confirmPosition, err
}
