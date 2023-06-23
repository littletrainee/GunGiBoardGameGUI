// 棋盤上放置駒的區塊物件
package block

import (
	"image"
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// 放置駒的區塊
type Block struct {
	Name         image.Point // Block的名稱，可以幫助駒定位在棋盤上的位置
	CurrentColor color.RGBA  // Block的顏色，用於顯示可以移動或是不可移動的範圍
	Coordinate   image.Point // Block在畫面上的座標，可以輔助定位駒在畫面上的確切座標
	KomaStack    []koma.Koma // Block上的駒堆疊切片
}
