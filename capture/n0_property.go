package capture

import "github.com/hajimehoshi/ebiten/v2"

type Capture struct {
	Show              bool // 是否顯示俘獲畫面
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
