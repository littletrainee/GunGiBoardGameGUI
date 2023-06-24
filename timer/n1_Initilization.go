package timer

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Initilization 初始化棋鐘，並回傳
//
// 參數amontOfTime為剩餘思考時間，x,y為棋鐘在畫面上的座標
func Initilization(amontOfTime int, x, y int) Timer {
	return Timer{
		RemainingTime:     amontOfTime,
		StopCountDown:     make(chan bool),
		CurrentCoordinate: image.Point{X: x, Y: y},
		Op:                &ebiten.DrawImageOptions{},
	}
}
