package colorholder

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"golang.org/x/image/font"
)

// Initilization 初始化ColorHolder物件，並設定其所包含的標題圖像、駒切片後回傳
//
// 參數font為文字所需要的字型來源
func Initilization(font font.Face) ColorHolder {
	var sc ColorHolder = ColorHolder{
		TitleImg: ebiten.NewImage(300, 300),
		TitleOpt: &ebiten.DrawImageOptions{},
		KomaList: []koma.Koma{
			{
				Name:              "帥",
				Color:             color.White,
				CurrentCoordinate: image.Point{X: 472, Y: 384},
				Img:               ebiten.NewImage(int(constant.KOMA_SIZE)+1, int(constant.KOMA_SIZE)+1),
				Op:                &ebiten.DrawImageOptions{},
			},
			{
				Name:              "帥",
				Color:             color.Black,
				CurrentCoordinate: image.Point{X: 552, Y: 384},
				Img:               ebiten.NewImage(int(constant.KOMA_SIZE)+1, int(constant.KOMA_SIZE)+1),
				Op:                &ebiten.DrawImageOptions{},
			}},
	}
	sc.TitleOpt.GeoM.Scale(3, 3)
	sc.TitleOpt.GeoM.Translate(372, 100)
	text.Draw(sc.TitleImg, "請選擇顏色", font, 0, 50, color.Black)

	for i := range sc.KomaList {
		sc.KomaList[i].SetGeoMetry(0)
		text.Draw(sc.KomaList[i].Img, sc.KomaList[i].Name, font, 0, 20, color.Black)

	}
	return sc
}
