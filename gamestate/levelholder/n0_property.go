// 階級控制物件
package levelholder

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/levelholder/level"
	_level "github.com/littletrainee/gunginotationgenerator/enum/level"
)

// 選擇進行的階級
type LevelHolder struct {
	TitleImg     *ebiten.Image            // 階級標題的圖像
	TitleOpt     *ebiten.DrawImageOptions // 階級標題圖像的參數
	LevelList    []level.Level            // 階級選項切片
	CurrentLevel _level.Level             // 玩家所選擇的階級
	MaxLayer     int                      // 最大段位
}
