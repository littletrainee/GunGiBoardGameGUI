package selectcolor

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"golang.org/x/image/font"
)

func Initilization(font font.Face) SelectColor {
	var rt SelectColor = SelectColor{
		Title: "請選擇顏色",
		KomaList: []koma.Koma{
			{
				Name:              "帥",
				Color:             color.White,
				CurrentCoordinate: image.Point{X: 472, Y: 384},
				Img:               ebiten.NewImage(int(constant.RADIUS)+1, int(constant.RADIUS)+1),
				Op:                &ebiten.DrawImageOptions{},
			},
			{
				Name:              "帥",
				Color:             color.Black,
				CurrentCoordinate: image.Point{X: 552, Y: 384},
				Img:               ebiten.NewImage(int(constant.RADIUS)+1, int(constant.RADIUS)+1),
				Op:                &ebiten.DrawImageOptions{},
			}},
	}
	for i := range rt.KomaList {
		rt.KomaList[i].SetGeoMetry(0)
		text.Draw(rt.KomaList[i].Img, rt.KomaList[i].Name, font, 0, 20, color.Black)

	}
	return rt
}
