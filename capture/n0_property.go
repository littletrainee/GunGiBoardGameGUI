// 詢問玩家是否要俘獲對方駒的物件
package capture

import "github.com/hajimehoshi/ebiten/v2"

// 詢問俘獲物件
type Capture struct {
	Show              bool                     // 顯示的切換，true為顯示，false為否
	ControlBool       bool                     // 是否啟用控制按鈕的切換，true為啟用，false為不啟用
	CaptureImg        *ebiten.Image            // 顯示俘獲文字的圖像
	CancelImg         *ebiten.Image            // 顯示取消文字的圖像
	EnableControlImg  *ebiten.Image            // 顯示啟用控制文字的圖像
	DisableControlImg *ebiten.Image            // 顯示為啟用控制文字的圖像
	CaptureOpt        *ebiten.DrawImageOptions // 俘獲文字的繪製參數
	CancelOpt         *ebiten.DrawImageOptions // 取消文字的繪製參數
	ControlOpt        *ebiten.DrawImageOptions // 控制文字的繪製參數
}
