package level

import (
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	_level "github.com/littletrainee/gunginotationgenerator/enum/level"
)

func Initilization() []Level {
	var L []Level = []Level{}
	for i := 0; i < 4; i++ {
		temp := Level{}
		switch i {
		case 0:
			temp.Name = "入門"
			temp.Code = _level.BEGINNER
			temp.X = constant.LEVEL_BEGINNER_X
			temp.backGround = color.RGBA{R: 255, G: 255, B: 255, A: 255}
			temp.text = color.Black
		case 1:
			temp.Name = "初級"
			temp.Code = _level.ELEMENTARY
			temp.X = constant.LEVEL_ELEMENTYARY_X
			temp.backGround = color.RGBA{R: 170, G: 170, B: 170, A: 170}
			temp.text = color.Black
		case 2:
			temp.Name = "中級"
			temp.Code = _level.INTERMEDIATE
			temp.X = constant.LEVEL_INTERMEDIATE_X
			temp.backGround = color.RGBA{R: 85, G: 85, B: 85, A: 85}
			temp.text = color.White
		case 3:
			temp.Name = "高級"
			temp.Code = _level.ADVANCED
			temp.X = constant.LEVEL_ADVANCED_X
			temp.backGround = color.RGBA{}
			temp.text = color.White
		}
		temp.Y = constant.LEVEL_Y
		L = append(L, temp)
	}
	return L
}
