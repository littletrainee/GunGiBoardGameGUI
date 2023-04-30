package timer

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func Initilization(amontOfTime int, x, y int) Timer {
	return Timer{
		RemainingTime:     amontOfTime,
		StopCountDown:     make(chan bool),
		CurrentCoordinate: image.Point{X: x, Y: y},
		Op:                &ebiten.DrawImageOptions{},
	}
}
