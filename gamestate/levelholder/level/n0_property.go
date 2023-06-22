// 供玩家選擇階級的物件
package level

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	_level "github.com/littletrainee/gunginotationgenerator/enum/level"
)

// 階級物件
type Level struct {
	Code       _level.Level             // 階級代號
	Coor       image.Point              // 階級按鈕在畫面上的座標
	Background color.RGBA               // 階級按鈕的背景顏色
	Img        *ebiten.Image            // 階級按鈕文字的圖像
	Opt        *ebiten.DrawImageOptions // 階級按鈕文字圖像的參數
}
