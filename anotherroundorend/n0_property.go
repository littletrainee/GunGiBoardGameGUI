package anotherroundorend

import "github.com/hajimehoshi/ebiten/v2"

type AnotherRoundOrEnd struct {
	Show             bool
	AnotherRoundText string
	EndText          string
	AnotherRoundImg  *ebiten.Image
	EndImg           *ebiten.Image
	AnotherRoundOpt  *ebiten.DrawImageOptions
	EndOpt           *ebiten.DrawImageOptions
}
