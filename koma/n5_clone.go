package koma

import "github.com/hajimehoshi/ebiten/v2"

// Clone 複製除了除了繪製位置以外的駒
func (k *Koma) Clone() Koma {
	return Koma{
		Name:              k.Name,
		Color:             k.Color,
		CurrentCoordinate: k.CurrentCoordinate,
		CurrentPosition:   k.CurrentPosition,
		Img:               k.Img,
		Op:                &ebiten.DrawImageOptions{},
		MoveableRange:     k.MoveableRange,
	}
}
