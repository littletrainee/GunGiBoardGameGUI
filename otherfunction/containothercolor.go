package otherfunction

import "github.com/littletrainee/GunGiBoardGameGUI/koma"

// containOtherColor 檢查是否有包含對方的駒，有包含回傳true，否則false
func containOtherColor(targetBlockKomaStack []koma.Koma, currentKoma koma.Koma) bool {
	for _, v := range targetBlockKomaStack {
		if v.Color != currentKoma.Color {
			return true
		}
	}
	return false
}
