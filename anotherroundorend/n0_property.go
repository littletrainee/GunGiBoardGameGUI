// 詢問玩家是否要再來一局或是離開遊戲的物件
package anotherroundorend

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/checkmate"
)

// 再來一局或是離開遊戲
type AnotherRoundOrEnd struct {
	Show                 bool                     // 顯示的切換，true為顯示，false為否
	CheckMateEnumeration checkmate.CheckMate      // 將軍還是被將軍，CHECKMATE為將軍，BE_CHECKMATE為被將軍
	TitleCheckMateImg    *ebiten.Image            // 顯示標題將軍文字的透明圖像
	TitleCheckMateOpt    *ebiten.DrawImageOptions // 顯示標題將軍文字透明圖像的參數
	TitleBeCheckMateImg  *ebiten.Image            // 顯示標題被將軍文字的透明圖像
	TitleBeCheckMateOpt  *ebiten.DrawImageOptions // 顯示標題被將軍文字透明圖像的參數
	AnotherRoundImg      *ebiten.Image            // 顯示再來一局文字的透明圖像
	AnotherRoundOpt      *ebiten.DrawImageOptions // 再來一局文字的繪製參數
	EndImg               *ebiten.Image            // 顯示離開遊戲文字的透明圖像
	EndOpt               *ebiten.DrawImageOptions // 離開遊戲文字的繪製參數
}
