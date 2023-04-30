package recommendarrangement

import (
	"image/color"
	"math"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func Initilization() []RecommendedArrangement {
	var I []RecommendedArrangement = []RecommendedArrangement{}
	for i := 0; i < 2; i++ {
		temp := RecommendedArrangement{}
		switch i {
		case 0:
			temp.Name = "預設布陣"
			temp.Code = true
			temp.X = constant.RECOMMENDED_ARRANGEMENT
			temp.backGround = color.White
			temp.text = color.Black
		case 1:
			temp.Name = "自行布陣"
			temp.Code = false
			temp.X = constant.NO_RECOMMENDED_ARRANGEMENT
			temp.backGround = color.Gray16{Y: math.MaxUint16 / 2}
			temp.text = color.White
		}
		temp.Y = constant.RECOMMENDED_ARRANGEMENT_Y
		I = append(I, temp)
	}

	return I
}
