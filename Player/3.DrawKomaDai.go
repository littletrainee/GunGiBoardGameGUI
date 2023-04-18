package Player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

func (p *Player) DrawKomaDai(screen *ebiten.Image, font font.Face) {
	for _, v := range p.KomaDai {
		if v.Item2 > 0 {
			v.Item1.Draw(screen, font)
		}
	}
}
