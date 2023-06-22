// 顏色控制物件
package colorholder

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// 選擇玩家顏色
type ColorHolder struct {
	TitleImg *ebiten.Image            // 玩家選擇顏色的標題圖像
	TitleOpt *ebiten.DrawImageOptions // 玩家選擇顏色標題圖像的參數
	KomaList []koma.Koma              // 兩個選擇放在一個切片裡面，以迭代的方式做處理
	Own      color.Gray16             // 本家的顏色
	Turn     color.Gray16             // 執棋權
}
