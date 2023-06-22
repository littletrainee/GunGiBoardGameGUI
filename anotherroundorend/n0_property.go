// 詢問玩家是否要再來一局或是離開遊戲的物件
package anotherroundorend

import "github.com/hajimehoshi/ebiten/v2"

// 再來一局或是離開遊戲
type AnotherRoundOrEnd struct {
	Show            bool                     // 顯示的切換，true為顯示，false為否
	TitleImg        *ebiten.Image            // 顯示標題文字的透明圖像
	AnotherRoundImg *ebiten.Image            // 顯示再來一局文字的透明圖像
	EndImg          *ebiten.Image            // 顯示離開遊戲文字的透明圖像
	TitleOpt        *ebiten.DrawImageOptions // 標題文字的繪製參數
	AnotherRoundOpt *ebiten.DrawImageOptions // 再來一局文字的繪製參數
	EndOpt          *ebiten.DrawImageOptions // 離開遊戲文字的繪製參數
}
