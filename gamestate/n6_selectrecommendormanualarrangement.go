package gamestate

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
)

// 設置預設布陣與否
func (g *GameState) SelectRecommendOrManualArrangement() {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		// 透過迴圈檢查兩個選項當中的其中一個是否有被點選
		for _, v := range g.ArrangementHolder.ArrangementList {
			if v.OnArrangementButton(x, y) {
				// 設定
				g.ArrangementHolder.ArramgementBy = v.ArrangementSelect
				g.Phase = phase.SET_COUNTDOWN_FOR_GAMING
				break
			}
		}
	}
}
