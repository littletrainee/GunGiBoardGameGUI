// 謀的叛變視窗物件
package mutiny

import "github.com/hajimehoshi/ebiten/v2"

// 謀的叛變
type Mutiny struct {
	Show      bool                     // 顯示切換，true為顯示，false為否
	TitleImg  *ebiten.Image            // 標題文字的透明圖像
	TitleOpt  *ebiten.DrawImageOptions // 標題文字透明圖像的參數
	MutinyImg *ebiten.Image            // 叛變文字的透明圖像
	MutinyOpt *ebiten.DrawImageOptions // 叛變文字透明視窗的參數
	CancelImg *ebiten.Image            // 取消文字的透明圖像
	CancelOpt *ebiten.DrawImageOptions // 取消文字透明圖像的參數
}
