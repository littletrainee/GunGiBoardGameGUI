package levelholder

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/levelholder/level"
	_level "github.com/littletrainee/gunginotationgenerator/enum/level"
	"golang.org/x/image/font"
)

// Initilization 初始化LevelHolder物件，並設定其所包含的標題圖像、階級切片後回傳
//
// 參數font為文字所需要的字型來源
func Initilization(font font.Face) LevelHolder {
	var (
		sl LevelHolder = LevelHolder{
			TitleImg:  ebiten.NewImage(300, 300),
			TitleOpt:  &ebiten.DrawImageOptions{},
			LevelList: []level.Level{},
		}
		templevel level.Level
	)
	for i := 0; i < 4; i++ {
		templevel = level.Level{}
		switch i {
		case 0:
			templevel.Code = _level.BEGINNER
			templevel.Coor = image.Point{
				X: constant.LEVEL_BEGINNER_X,
				Y: constant.LEVEL_Y,
			}
			templevel.Background = color.RGBA{R: 255, G: 255, B: 255, A: 255}
			templevel.Img = ebiten.NewImage(constant.LEVEL_BLOCK_SIZE, constant.LEVEL_BLOCK_SIZE)
			templevel.Opt = &ebiten.DrawImageOptions{}
			templevel.Opt.GeoM.Translate(float64(templevel.Coor.X), float64(templevel.Coor.Y))
			text.Draw(templevel.Img, "入門", font, 30, 55, color.Black)
		case 1:
			templevel.Code = _level.ELEMENTARY
			templevel.Coor = image.Point{
				X: constant.LEVEL_ELEMENTYARY_X,
				Y: constant.LEVEL_Y,
			}
			templevel.Background = color.RGBA{R: 170, G: 170, B: 170, A: 170}
			templevel.Img = ebiten.NewImage(constant.LEVEL_BLOCK_SIZE, constant.LEVEL_BLOCK_SIZE)
			templevel.Opt = &ebiten.DrawImageOptions{}
			templevel.Opt.GeoM.Translate(float64(templevel.Coor.X), float64(templevel.Coor.Y))
			text.Draw(templevel.Img, "初級", font, 30, 55, color.Black)
		case 2:
			templevel.Code = _level.INTERMEDIATE
			templevel.Coor = image.Point{
				X: constant.LEVEL_INTERMEDIATE_X,
				Y: constant.LEVEL_Y,
			}
			templevel.Background = color.RGBA{R: 85, G: 85, B: 85, A: 85}
			templevel.Img = ebiten.NewImage(constant.LEVEL_BLOCK_SIZE, constant.LEVEL_BLOCK_SIZE)
			templevel.Opt = &ebiten.DrawImageOptions{}
			templevel.Opt.GeoM.Translate(float64(templevel.Coor.X), float64(templevel.Coor.Y))
			text.Draw(templevel.Img, "中級", font, 30, 55, color.White)
		case 3:
			templevel.Code = _level.ADVANCED
			templevel.Coor = image.Point{
				X: constant.LEVEL_ADVANCED_X,
				Y: constant.LEVEL_Y,
			}
			templevel.Background = color.RGBA{}
			templevel.Img = ebiten.NewImage(constant.LEVEL_BLOCK_SIZE, constant.LEVEL_BLOCK_SIZE)
			templevel.Opt = &ebiten.DrawImageOptions{}
			templevel.Opt.GeoM.Translate(float64(templevel.Coor.X), float64(templevel.Coor.Y))
			text.Draw(templevel.Img, "高級", font, 30, +55, color.White)
		}
		sl.LevelList = append(sl.LevelList, templevel)
	}
	return sl
}
