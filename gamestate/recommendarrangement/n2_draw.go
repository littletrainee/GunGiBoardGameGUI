package recommendarrangement

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"golang.org/x/image/font"
)

func (i *RecommendedArrangement) Draw(screen *ebiten.Image, font font.Face) {
	vector.DrawFilledRect(screen, i.X, i.Y,
		constant.RECOMMENDED_ARRANGEMENT_BLOCK_SIZE,
		constant.RECOMMENDED_ARRANGEMENT_BLOCK_SIZE, color.Black, true)
	vector.DrawFilledRect(screen, i.X+1, i.Y+1,
		constant.RECOMMENDED_ARRANGEMENT_BLOCK_SIZE-2,
		constant.RECOMMENDED_ARRANGEMENT_BLOCK_SIZE-2, i.backGround, true)
	text.Draw(screen, i.Name, font, int(i.X)+35, int(i.Y)+83, i.text)
}
