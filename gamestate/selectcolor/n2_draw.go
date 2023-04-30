package selectcolor

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

func (s *SelectColor) Draw(screen *ebiten.Image, font font.Face) {
	img := ebiten.NewImage(300, 300)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(3, 3)
	op.GeoM.Translate(372, 100)
	text.Draw(img, s.Title, font, 0, 50, color.Black)
	screen.DrawImage(img, op)
	for _, v := range s.KomaList {
		v.Draw(screen, font)
	}
}
