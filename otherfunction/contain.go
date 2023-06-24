package otherfunction

import "image"

// contain 確認是否有包含目標物件，若咬包含則回傳true，否則false
//
// 參數slice為要確認的切片，target為目標元素
func Contain(slice []image.Point, target image.Point) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}
