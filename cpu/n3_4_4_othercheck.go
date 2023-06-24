package cpu

import (
	"errors"
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/action"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/levelholder"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/GunGiBoardGameGUI/otherfunction"
)

// otherCheck 其他駒可以移動的範圍，回傳值為可以移動的選項切片
//
// 參數eachCanCanMove為每個可以移動方向中的當前段可以移動的範圍，lastKoma為當前駒物件，g為本局遊戲的狀態，b為棋盤物件，currentDan為當前的段
func otherCheck(levelholder levelholder.LevelHolder, whereOther []koma.Koma, b board.Board) ([]image.Point, error) {
	// 設定是否已經有阻礙
	var (
		hinder          bool
		confirmPosition []image.Point
		currentDan      int       = len(whereOther)
		other           koma.Koma = whereOther[currentDan-1]
		err             error
	)
	for _, direction := range other.MoveableRange {
		for danIndex, eachDanCanMove := range direction {
			// 確定沒有超過當前階級可以的最高段數
			if danIndex < currentDan {
				// 從每個段的移動迭代每個位置
				for _, coor := range eachDanCanMove {
					// 設定目標位置
					targetBlockPosition := image.Point{
						X: other.CurrentPosition.X - coor.X,
						Y: other.CurrentPosition.Y - coor.Y,
					}

					// 確認目標位置是否在棋盤內
					targetBlock, exists := b.Blocks[targetBlockPosition]

					// 若在棋盤內
					if exists && !hinder {
						targetBlockLen := len(targetBlock.KomaStack)

						if targetBlockLen > currentDan {
							break
						} else if targetBlockLen == currentDan {
							hinder = true
						}

						switch otherfunction.Move(targetBlock.KomaStack, whereOther, levelholder) {
						case action.MOVE, action.CAPTURE_ONLY, action.CAPTURE_OR_CONTROL:
							if !targetBlock.HasSuI() {
								confirmPosition = append(confirmPosition, targetBlockPosition)
							}
						}
					} else {
						break
					}

				}
			}
		}
	}
	if len(confirmPosition) == 0 {
		err = errors.New("otherCheck的confirmPosition是空的")
	}
	return confirmPosition, err
}
