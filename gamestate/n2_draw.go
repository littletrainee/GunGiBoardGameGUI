package gamestate

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
)

// Draw 繪製GameState相關的畫面，例如最初的玩家選擇顏色、遊玩的階級、並且基於選擇入門或是初級階級，而有選擇推薦布陣或是手動布陣的選項
//
// 參數screen為要繪製的目標視窗，font為要繪製的文字
func (g *GameState) Draw(screen *ebiten.Image) {
	switch g.Phase {
	case phase.SELECT_COLOR:
		g.ColorHolder.Draw(screen)
	case phase.SELECT_LEVEL:
		for _, v := range g.LevelHolder.LevelList {
			v.Draw(screen)
		}
	case phase.RECOMMEND_OR_MANUAL_ARRANGEMENT:
		for _, v := range g.ArrangementHolder.RecommendOrManual {
			v.Draw(screen)
		}
	}
}
