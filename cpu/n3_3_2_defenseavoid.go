package cpu

import (
	"image"
	"log"
	"math/rand"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/otherfunction"
)

// defenseAvoid 閃躲防禦函式，可以導出帥移動到哪個位置可以避免被將軍的狀況，若可以避免則回傳true，否則回傳false
//
// 參數b為棋盤物件
func (c *CPU) defenseAvoid(g gamestate.GameState, b board.Board) bool {
	var (
		suIPosition            image.Point = image.Point{X: c.MoveToTarget[2], Y: c.MoveToTarget[3]}
		confirmSuIProbablyMove []image.Point
		confirmPosition        []image.Point
		err                    error
	)

	// 確認帥可以移動的位置
	confirmSuIProbablyMove, err = suICheck(g.LevelHolder, b.Blocks[suIPosition].KomaStack, b)
	if err != nil {
		log.Println(err)
	}

	// 迭代所有核可的移動範圍
	for _, targetPos := range confirmSuIProbablyMove {
		if !stillInCaptureRange(b, targetPos, c.SelfColor, g) {
			if !otherfunction.Contain(confirmPosition, targetPos) {
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
