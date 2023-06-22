// 布陣控制物件
package arrangementholder

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/arrangementselect"
	roma "github.com/littletrainee/GunGiBoardGameGUI/gamestate/arrangementholder/recommendormanualarrangement"
)

// 布陣控制
type ArrangementHolder struct {
	TitleImg        *ebiten.Image                       // 玩家選擇推薦或是手動佈陣的標題圖像
	TitleOpt        *ebiten.DrawImageOptions            // 玩家選擇推薦或是手動佈陣標題圖像的參數
	ArrangementList []roma.RecommendOrManualArrangement // 布陣的選擇
	ArramgementBy   arrangementselect.ArrangementSelect // 在入門或出擊的推薦或是自訂布陣的選擇
}
