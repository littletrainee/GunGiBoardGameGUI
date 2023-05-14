package gamehandler

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) CaptureOrControl() {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if g.Capture.CaptureButton(x, y) {

		}
		if g.Capture.ControlButton(x, y) {

		}
	}
}
