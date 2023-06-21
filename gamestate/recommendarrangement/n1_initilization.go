package recommendarrangement

import (
	"image/color"
	"math"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/arrangementselect"
)

func Initilization() []RecommendedArrangement {
	var I []RecommendedArrangement = []RecommendedArrangement{}
	for i := 0; i < 2; i++ {
		temp := RecommendedArrangement{}
		switch i {
		case 0:
			temp.Name = "預設布陣"
			temp.Select = arrangementselect.RECOMMEND
			temp.X = constant.RECOMMENDED_ARRANGEMENT_BUTTON_X
			temp.backGround = color.White
			temp.text = color.Black
		case 1:
			temp.Name = "自行布陣"
			temp.Select = arrangementselect.MANUAL
			temp.X = constant.MANUAL_ARRANGEMENT_BUTTON_X
			temp.backGround = color.Gray16{Y: math.MaxUint16 / 2}
			temp.text = color.White
		}
		temp.Y = constant.ARRANGEMENT_BUTTON_Y
		I = append(I, temp)
	}

	return I
}
