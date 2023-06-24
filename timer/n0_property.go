// 棋鐘物件
package timer

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// 棋鐘
type Timer struct {
	RemainingTime     int                      // 剩餘時間
	Background        color.RGBA               // 棋鐘的當前背景顏色
	StopCountDown     chan bool                // 棋鐘停止的通道
	CurrentCoordinate image.Point              // 棋鐘的座標
	Img               *ebiten.Image            // 棋鐘文字的透明圖像
	Op                *ebiten.DrawImageOptions // 棋鐘文字透明圖像的參數
}
