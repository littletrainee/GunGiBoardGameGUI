package cpu

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// defenseAvoid 閃躲防禦函式，可以導出帥移動到哪個位置可以避免被將軍的狀況，若可以避免則回傳true，否則回傳false
//
// 參數b為棋盤物件
func (c *CPU) defenseAvoid(b board.Board) bool {
	var (
		suiPosition            image.Point = image.Point{X: c.MoveToTarget[2], Y: c.MoveToTarget[3]}
		confirmSuIProbablyMove []image.Point
		confirmPosition        []image.Point
	)

	// 確認帥可以移動的位置
	confirmSuIProbablyMove = confirmMove(b, suiPosition)

	// 迭代所有核可的移動範圍
	for _, targetPos := range confirmSuIProbablyMove {
		if !stillInCaptureRange(b, targetPos, c.SelfColor) {
			if !contain(confirmPosition, targetPos) {
				confirmPosition = append(confirmPosition, targetPos)
			}
		}
	}

	if len(confirmPosition) > 0 {
		rnd := confirmPosition[rand.Intn(len(confirmPosition))]
		c.MoveToTarget = []int{suiPosition.X, suiPosition.Y, rnd.X, rnd.Y}
		return true
	}
	return false
}

func contain(slice []image.Point, target image.Point) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

func stillInCaptureRange(b board.Board, confirmPosition image.Point, selfColor color.Gray16) bool {
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
			for _, direction := range theTopOne.ProbablyMoveing {
				hinder = false
				// 迭代允許的階層
				for i, layer := range direction {
					if i < stackLength {
						// 迭代允許的位置
						for _, pos := range layer {
							if !hinder {
								// 設定目標位置
								targetPosition := image.Point{
									X: theTopOne.CurrentPosition.X - pos.X,
									Y: theTopOne.CurrentPosition.Y + pos.Y,
								}
								// 確認目標位置是否在棋盤內
								targetBlock, ok := b.Blocks[targetPosition]
								// 若目標位置在棋盤內
								if ok {
									// 若在棋盤內並且沒有牆阻礙
									if len(targetBlock.KomaStack) == stackLength {
										hinder = true
									} else if len(targetBlock.KomaStack) > stackLength {
										break
									}

									if targetPosition == confirmPosition {
										return true
									}
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
			}
		}
	}
	return false
}

// 帥可以移動的範圍
func confirmMove(b board.Board, suiPosition image.Point) []image.Point {
	var (
		confirmmove     []image.Point
		komaStackLength int       = len(b.Blocks[suiPosition].KomaStack)
		suIKoma         koma.Koma = b.Blocks[suiPosition].KomaStack[komaStackLength-1]
	)
	for _, direction := range suIKoma.ProbablyMoveing {
		for i, layer := range direction {
			if i < komaStackLength {
				for _, pos := range layer {
					targetPosition := image.Point{
						X: suIKoma.CurrentPosition.X - pos.X,
						Y: suIKoma.CurrentPosition.Y - pos.Y,
					}
					targetBlock, ok := b.Blocks[targetPosition]
					if ok {
						if len(targetBlock.KomaStack) == 0 {
							confirmmove = append(confirmmove, targetPosition)
						} else {
							break
						}
					} else {
						break
					}
				}
			}
		}
	}

	return confirmmove
}