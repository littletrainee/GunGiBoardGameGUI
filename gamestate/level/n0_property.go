package level

import (
	"image/color"

	_level "github.com/littletrainee/gunginotationgenerator/enum/level"
)

// 階級物件
type Level struct {
	Name       string
	Code       _level.Level
	X          float32
	Y          float32
	text       color.Gray16
	backGround color.RGBA
}
