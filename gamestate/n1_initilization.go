package gamestate

import (
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/colorholder"
	"golang.org/x/image/font"
)

// Initilization 初始化GameState物件，設定好預設的顯示文字與位置並回傳。
//
// 參數font用於文字顯示得字型
func Initilization(font font.Face) GameState {
	return GameState{
		Phase:       phase.SELECT_COLOR,
		ColorHolder: colorholder.Initilization(font),
	}
}
