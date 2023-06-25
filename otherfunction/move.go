// 其他函式
package otherfunction

import (
	"github.com/littletrainee/GunGiBoardGameGUI/enum/action"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/levelholder"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
)

// Move 每個駒必定會碰到的狀況是可不可移動，可不可俘獲的狀況，因此將邏輯判斷放於此函式可以統一判斷是否核可，回傳的結果會因為判斷的不同而有不同的列舉值
//
// 參數targetKomaStack為目標位置的駒堆疊列，currentKomaStack為當前位置的駒堆疊列，levelholder為階級控制物件，用以判斷帥在初級以下不可向目標堆疊與中級以上可項目標堆疊的判斷
func Move(targetKomaStack, currentKomaStack []koma.Koma, levelholder levelholder.LevelHolder) action.Action {
	var (
		targetlen       int = len(targetKomaStack)
		targetLastKoma  koma.Koma
		currentlen      int = len(currentKomaStack)
		currentLastKoma koma.Koma
	)
	// 目標位置沒有駒 -> 直接移動
	if targetlen == 0 {
		return action.MOVE
	}

	// 目標位置有駒，段數 > 自己段數 -> X
	if targetlen > currentlen {
		return action.ILLEGAL_MOVE
	}
	targetLastKoma = targetKomaStack[targetlen-1]
	currentLastKoma = currentKomaStack[currentlen-1]

	// 自己駒是帥，階級是入門或是初階
	if currentLastKoma.Name == "帥" && levelholder.CurrentLevel == level.BEGINNER || levelholder.CurrentLevel == level.ELEMENTARY {
		// 非相同顏色 -> 俘獲
		if targetLastKoma.Color != currentLastKoma.Color {
			return action.CAPTURE_ONLY
		}
		// 相同顏色 -> X
		if targetLastKoma.Color == currentLastKoma.Color {
			return action.ILLEGAL_MOVE
		}
	} else {

		// 目標位置有駒，段數 = 自己段數
		if targetlen == currentlen {
			// 已達最大，非相同顏色
			if containOtherColor(targetKomaStack, currentLastKoma) && targetlen == levelholder.MaxLayer {
				// 最上方的顏色是自家的顏色-> X
				if targetLastKoma.Color == currentLastKoma.Color {
					return action.ILLEGAL_MOVE
					// 非自家俘獲
				} else {
					return action.CAPTURE_ONLY
				}
			}
			// 已達最大，相同顏色 -> X
			if !containOtherColor(targetKomaStack, currentLastKoma) && targetlen == levelholder.MaxLayer {
				return action.ILLEGAL_MOVE
			}
			// 未達最大，非相同顏色 -> 俘獲或控制
			if containOtherColor(targetKomaStack, currentLastKoma) && targetlen < levelholder.MaxLayer {
				return action.CAPTURE_OR_CONTROL
			}
			// 未達最大，相同顏色 -> 直接移動
			if !containOtherColor(targetKomaStack, currentLastKoma) && targetlen < levelholder.MaxLayer {
				return action.MOVE
			}
		}

		// 目標位置有駒，段數 < 自己段數
		if targetlen > 0 && targetlen < currentlen {
			// 相同顏色 -> 直接移動
			if !containOtherColor(targetKomaStack, currentLastKoma) {
				// move()
				return action.MOVE
			}
			// 非相同顏色 -> 俘獲或控制
			if containOtherColor(targetKomaStack, currentLastKoma) {
				return action.CAPTURE_OR_CONTROL
			}
		}
	}
	return action.NONE
}
