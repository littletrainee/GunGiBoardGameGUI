package arrangementholder

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/arrangementselect"
	roma "github.com/littletrainee/GunGiBoardGameGUI/gamestate/arrangementholder/recommendormanualarrangement"
	"golang.org/x/image/font"
)

// Initilization 初始化ArrangementHolder物件，設定好標題、推薦或自訂布陣的選擇後回傳。
//
// 參數font為顯示文字的字型
func Initilization(font font.Face) ArrangementHolder {
	var (
		a ArrangementHolder = ArrangementHolder{
			TitleImg:        ebiten.NewImage(300, 300),
			TitleOpt:        &ebiten.DrawImageOptions{},
			ArrangementList: []roma.RecommendOrManualArrangement{},
		}
	)
	for i := 0; i < 2; i++ {
		temp := roma.RecommendOrManualArrangement{}
		switch i {
		case 0:
			temp.ArrangementSelect = arrangementselect.RECOMMEND
			temp.Coor = image.Point{
				X: constant.RECOMMENDED_ARRANGEMENT_BUTTON_X,
				Y: constant.ARRANGEMENT_BUTTON_Y,
			}
			temp.Background = color.White
			temp.Img = ebiten.NewImage(constant.ARRANGEMENT_BUTTON_SIZE, constant.ARRANGEMENT_BUTTON_SIZE)
			temp.Opt = &ebiten.DrawImageOptions{}
			temp.Opt.GeoM.Translate(float64(temp.Coor.X), float64(temp.Coor.Y))
			text.Draw(temp.Img, "推薦布陣", font, 35, 83, color.Black)
		case 1:
			temp.ArrangementSelect = arrangementselect.MANUAL
			temp.Coor = image.Point{
				X: constant.MANUAL_ARRANGEMENT_BUTTON_X,
				Y: constant.ARRANGEMENT_BUTTON_Y,
			}
			temp.Background = color.Gray16{Y: math.MaxUint16 / 2}
			temp.Img = ebiten.NewImage(constant.ARRANGEMENT_BUTTON_SIZE, constant.ARRANGEMENT_BUTTON_SIZE)
			temp.Opt = &ebiten.DrawImageOptions{}
			temp.Opt.GeoM.Translate(float64(temp.Coor.X), float64(temp.Coor.Y))
			text.Draw(temp.Img, "自行布陣", font, 35, 83, color.White)
		}
		a.ArrangementList = append(a.ArrangementList, temp)
	}
	a.TitleOpt.GeoM.Scale(3, 3)
	a.TitleOpt.GeoM.Translate(constant.ARRANGEMENT_TITLE_X, constant.TITLE_Y)
	text.Draw(a.TitleImg, "請選擇推薦或自訂布陣", font, 0, 50, color.Black)

	return a
}
