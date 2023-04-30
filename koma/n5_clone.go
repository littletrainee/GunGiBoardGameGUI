package koma

import "github.com/hajimehoshi/ebiten/v2"

func (k *Koma) Clone() Koma {
	return Koma{
		Name:              k.Name,
		Color:             k.Color,
		CurrentCoordinate: k.CurrentCoordinate,
		CurrentPosition:   k.CurrentPosition,
		Img:               k.Img,
		Op:                &ebiten.DrawImageOptions{},
		ProbablyMoveing:   k.ProbablyMoveing,
	}
}
