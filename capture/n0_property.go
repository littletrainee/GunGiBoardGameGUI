package capture

import "github.com/hajimehoshi/ebiten/v2"

type Capture struct {
	Show        bool
	CaptureBool bool
	ControlBool bool
	CaptureText string
	ControlText string
	CaptureImg  *ebiten.Image
	ControlImg  *ebiten.Image
	CaptureOpt  *ebiten.DrawImageOptions
	ControlOpt  *ebiten.DrawImageOptions
}
