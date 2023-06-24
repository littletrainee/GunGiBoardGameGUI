package otherfunction

import (
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// notContain 不包含自己的顏色，若不包含回傳true，包含回傳false
//
// 參數komalist為駒列表，targetColor為要確認的顏色
func NoteContain(komaStack []koma.Koma, targetColor color.Gray16) bool {
	for _, v := range komaStack {
		if v.Color == targetColor {
			return false
		}
	}
	return true
}
