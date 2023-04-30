package level

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"golang.org/x/image/font"
)

func (l *Level) Draw(screen *ebiten.Image, font font.Face) {
	vector.DrawFilledRect(screen, l.X, l.Y,
		constant.LEVEL_BLOCK_SIZE, constant.LEVEL_BLOCK_SIZE, color.Black, true)
	vector.DrawFilledRect(screen, l.X+1, l.Y+1,
		constant.LEVEL_BLOCK_SIZE-2, constant.LEVEL_BLOCK_SIZE-2, l.backGround, true)
	text.Draw(screen, l.Name, font, int(l.X)+30, int(l.Y)+55, l.text)
}
