package anotherroundorend

import "github.com/hajimehoshi/ebiten/v2"

type AnotherRoundOrEnd struct {
	Show             bool
	TitleText        string
	AnotherRoundText string
	EndText          string
	TitleImg         *ebiten.Image
	AnotherRoundImg  *ebiten.Image
	EndImg           *ebiten.Image
	TitleOpt         *ebiten.DrawImageOptions
	AnotherRoundOpt  *ebiten.DrawImageOptions
	EndOpt           *ebiten.DrawImageOptions
}
