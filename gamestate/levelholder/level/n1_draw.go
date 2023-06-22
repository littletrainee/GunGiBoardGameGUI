package level

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// Draw 繪製Level物件
//
// 參數screen為繪製的畫面
func (l *Level) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(l.Coor.X), float32(l.Coor.Y),
		constant.LEVEL_BLOCK_SIZE, constant.LEVEL_BLOCK_SIZE, color.Black, true)
	vector.DrawFilledRect(screen, float32(l.Coor.X)+1, float32(l.Coor.Y)+1,
		constant.LEVEL_BLOCK_SIZE-2, constant.LEVEL_BLOCK_SIZE-2, l.Background, true)
	screen.DrawImage(l.Img, l.Opt)
}
