package declaresumi

import "github.com/hajimehoshi/ebiten/v2"

type DeclareSuMi struct {
	Show     bool
	SuMiText string
	SuMiImg  *ebiten.Image
	SuMiOpt  *ebiten.DrawImageOptions
}
