// 詢問玩家是否要俘獲對方駒的物件
package capture

import "github.com/hajimehoshi/ebiten/v2"

// 詢問俘獲物件
type Capture struct {
	CancelImg                 *ebiten.Image            // 顯示取消文字的圖像
	CancelOpt                 *ebiten.DrawImageOptions // 取消文字的繪製參數
	CaptureImg                *ebiten.Image            // 顯示俘獲文字的圖像
	CaptureOpt                *ebiten.DrawImageOptions // 俘獲文字的繪製參數
	ControlBool               bool                     // 是否啟用控制按鈕的切換，true為啟用，false為不啟用
	ControlOpt                *ebiten.DrawImageOptions // 控制文字的繪製參數
	DisableControlImg         *ebiten.Image            // 顯示為啟用控制文字的圖像
	EnableControlImg          *ebiten.Image            // 顯示啟用控制文字的圖像
	Show                      bool                     // 顯示的切換，true為顯示，false為否
	TitleCaptureAndControlImg *ebiten.Image            // 顯示標題文字俘獲或控制的透明圖像
	TitleCaptureAndControlOpt *ebiten.DrawImageOptions // 顯示標題文字俘獲或控制透明圖像的參數
	TitleOnlyCaptureImg       *ebiten.Image            // 顯示標題文字僅有俘獲的透明圖像
	TitleOnlyCaptureOpt       *ebiten.DrawImageOptions // 顯示標題文字僅有俘獲透明圖像的參數
}
