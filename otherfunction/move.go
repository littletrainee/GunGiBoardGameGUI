package otherfunction

import (
	"github.com/littletrainee/GunGiBoardGameGUI/enum/action"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/levelholder"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
)

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
			// 已達最大，非相同顏色 -> 俘獲
			if targetLastKoma.Color != currentLastKoma.Color && targetlen == levelholder.MaxLayer {
				return action.CAPTURE_ONLY
			}
			// 已達最大，相同顏色 -> X
			if targetLastKoma.Color == currentLastKoma.Color && targetlen == levelholder.MaxLayer {
				return action.ILLEGAL_MOVE
			}
			// 未達最大，非相同顏色 -> 俘獲或控制
			if targetLastKoma.Color != currentLastKoma.Color && targetlen < levelholder.MaxLayer {
				return action.CAPTURE_OR_CONTROL
			}
			// 未達最大，相同顏色 -> 直接移動
			if targetLastKoma.Color == currentLastKoma.Color && targetlen < levelholder.MaxLayer {
				return action.MOVE
			}
		}

		// 目標位置有駒，段數 < 自己段數
		if targetlen > 0 && targetlen < currentlen {
			// 相同顏色 -> 直接移動
			if targetLastKoma.Color == currentLastKoma.Color {
				// move()
				return action.MOVE
			}
			// 非相同顏色 -> 俘獲或控制
			if targetLastKoma.Color != currentLastKoma.Color {
				return action.CAPTURE_OR_CONTROL
			}
		}
	}
	return action.NONE
}
