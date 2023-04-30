package timer

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Timer struct {
	RemainingTime     int
	BackgroundColor   color.RGBA
	StopCountDown     chan bool
	CurrentCoordinate image.Point
	Img               *ebiten.Image
	Op                *ebiten.DrawImageOptions
}
