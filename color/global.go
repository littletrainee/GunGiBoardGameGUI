// 軍儀棋會用的顏色全域變數
package color

import "image/color"

var (
	BoardColor         color.RGBA = color.RGBA{R: 200, G: 139, B: 71, A: 255}  // 棋盤顏色
	DisableButtonColor color.RGBA = color.RGBA{R: 255, G: 179, B: 95, A: 255}  // 未啟動顏色
	ConfirmColor       color.RGBA = color.RGBA{R: 157, G: 250, B: 153, A: 255} // 確認顏色
	DenyColor          color.RGBA = color.RGBA{R: 254, G: 167, B: 160, A: 255} // 禁止顏色
	CaptureColor       color.RGBA = color.RGBA{R: 51, G: 200, B: 255, A: 64}   // 俘獲顏色

	TimerPauseColor color.RGBA = color.RGBA{R: 128, G: 128, B: 128, A: 255} // 棋鐘暫停的顏色
)
