// 駒物件
package koma

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// 駒物件
type Koma struct {
	Name              string                   // 駒所代表的名稱，類似代號
	Color             color.Gray16             // 駒的顏色
	CurrentCoordinate image.Point              // 駒在畫面上的座標
	CurrentPosition   image.Point              // 駒在棋盤上的座標
	Img               *ebiten.Image            // 駒文字的透明圖像
	Op                *ebiten.DrawImageOptions // 駒文字透明圖像的參數
	MoveableRange     [][][]image.Point        // 駒所能移動的範圍
}
