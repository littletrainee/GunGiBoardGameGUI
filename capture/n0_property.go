package capture

import "github.com/hajimehoshi/ebiten/v2"

type Capture struct {
	Show              bool
	CaptureBool       bool
	CancelBool        bool
	ControlBool       bool
	CaptureText       string
	CancelText        string
	ControlText       string
	CaptureImg        *ebiten.Image
	CancelImg         *ebiten.Image
	EnableControlImg  *ebiten.Image
	DisableControlImg *ebiten.Image
	CaptureOpt        *ebiten.DrawImageOptions
	CancelOpt         *ebiten.DrawImageOptions
	ControlOpt        *ebiten.DrawImageOptions
}
