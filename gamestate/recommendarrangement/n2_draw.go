package recommendarrangement

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"golang.org/x/image/font"
)

func (r *RecommendedArrangement) Draw(screen *ebiten.Image, font font.Face) {
	vector.DrawFilledRect(screen, r.X, r.Y,
		constant.RECOMMENDED_ARRANGEMENT_BLOCK_SIZE,
		constant.RECOMMENDED_ARRANGEMENT_BLOCK_SIZE, color.Black, true)
	vector.DrawFilledRect(screen, r.X+1, r.Y+1,
		constant.RECOMMENDED_ARRANGEMENT_BLOCK_SIZE-2,
		constant.RECOMMENDED_ARRANGEMENT_BLOCK_SIZE-2, r.backGround, true)
	text.Draw(screen, r.Name, font, int(r.X)+35, int(r.Y)+83, r.text)
}
