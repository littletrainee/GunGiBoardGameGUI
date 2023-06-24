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

func taiChuUCheck(levelholder levelholder.LevelHolder, whereTaiChuU []koma.Koma, b board.Board) ([]image.Point, error) {
	var (
		hinder          bool // 設定是否已經有阻礙
		confirmPosition []image.Point
		currentDan      int       = len(whereTaiChuU)
		taiChuU         koma.Koma = whereTaiChuU[currentDan-1]
		err             error
	)
	if taiChuU.Name == "大" {
		for d, v := range taiChuU.MoveableRange {
			if d == 0 || d == 2 || d == 4 || d == 6 { // 方向
				hinder = false
				for i := 0; i < len(v[0]); i++ {
					targetBlockPosition := image.Point{
						X: taiChuU.CurrentPosition.X - v[0][i].X,
						Y: taiChuU.CurrentPosition.Y + v[0][i].Y,
					}
					targetBlock, exists := b.Blocks[targetBlockPosition]
					if exists && !hinder {
						targetBlockLength := len(targetBlock.KomaStack)
						if targetBlockLength != 0 || targetBlockLength >= currentDan {
							hinder = true
						}

						switch otherfunction.Move(targetBlock.KomaStack, whereTaiChuU, levelholder) {
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
				for danIndex, eachDanCanMove := range v {
					if danIndex < currentDan {
						// 從每個段的移動迭代每個位置
						for _, coor := range eachDanCanMove {
							// 設定目標位置
							targetBlockPosition := image.Point{
								X: taiChuU.CurrentPosition.X - coor.X,
								Y: taiChuU.CurrentPosition.Y - coor.Y,
							}

							// 從blocks取出目標block與確認目標位置是否在棋盤內
							targetBlock, exists := b.Blocks[targetBlockPosition]

							// 若在棋盤內且沒有阻隔
							if exists && !hinder {
								// 從目標block取出目標的段數
								targetBlockLen := len(targetBlock.KomaStack)
								if targetBlockLen > currentDan {
									break
								} else if targetBlockLen == currentDan {
									hinder = true
								}

								switch otherfunction.Move(targetBlock.KomaStack, whereTaiChuU, levelholder) {
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
		}
	} else {
		for d, v := range taiChuU.MoveableRange {
			if d == 1 || d == 3 || d == 5 || d == 7 { // 方向
				hinder = false
				for i := 0; i < len(v[0]); i++ {
					targetblockPosition := image.Point{
						X: taiChuU.CurrentPosition.X - v[0][i].X,
						Y: taiChuU.CurrentPosition.Y + v[0][i].Y,
					}
					targetBlock, exists := b.Blocks[targetblockPosition]
					if exists && !hinder {
						targetBlockLen := len(targetBlock.KomaStack)
						if targetBlockLen > currentDan {
							break
						} else if targetBlockLen == currentDan {
							hinder = true
						}

						switch otherfunction.Move(targetBlock.KomaStack, whereTaiChuU, levelholder) {
						case action.MOVE, action.CAPTURE_ONLY, action.CAPTURE_OR_CONTROL:
							if !targetBlock.HasSuI() {
								confirmPosition = append(confirmPosition, targetblockPosition)
							}
						}
					} else {
						break
					}
				}
			} else {
				for danIndex, eachDanCanMove := range v {
					if danIndex < currentDan {
						// 從每個段的移動迭代每個位置
						for _, coor := range eachDanCanMove {
							// 設定目標位置
							targetBlockPosition := image.Point{
								X: taiChuU.CurrentPosition.X - coor.X,
								Y: taiChuU.CurrentPosition.Y - coor.Y,
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

								switch otherfunction.Move(targetBlock.KomaStack, whereTaiChuU, levelholder) {
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
		}
	}
	if len(confirmPosition) == 0 {
		err = errors.New("taiChuUCheck的confirmPosition 是空的")
	}
	return confirmPosition, err
}
