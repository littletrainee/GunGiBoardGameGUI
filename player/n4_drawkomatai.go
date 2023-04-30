package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

func (p *Player) DrawKomaTai(screen *ebiten.Image, font font.Face) {
	for _, v := range p.KomaTai {
		if v.Item2 > 0 {
			v.Item1.Draw(screen, font)
		}
	}
}
