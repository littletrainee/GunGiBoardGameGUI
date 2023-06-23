package cpu

import (
	"image"
	"math/rand"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
)

// defenseAvoid 閃躲防禦函式，可以導出帥移動到哪個位置可以避免被將軍的狀況，若可以避免則回傳true，否則回傳false
//
// 參數b為棋盤物件
func (c *CPU) defenseAvoid(b board.Board) bool {
	var (
		suIPosition            image.Point = image.Point{X: c.MoveToTarget[2], Y: c.MoveToTarget[3]}
		confirmSuIProbablyMove []image.Point
		confirmPosition        []image.Point
	)

	// 確認帥可以移動的位置
	confirmSuIProbablyMove = checkSuIConfirmMove(b, suIPosition)

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
		c.MoveToTarget = []int{suIPosition.X, suIPosition.Y, rnd.X, rnd.Y}
		return true
	}
	return false
}
