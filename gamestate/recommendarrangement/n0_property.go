package recommendarrangement

import (
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/enum/arrangementselect"
)

type RecommendedArrangement struct {
	Name       string
	Select     arrangementselect.ArrangementSelect
	X          float32
	Y          float32
	text       color.Gray16
	backGround color.Gray16
}
