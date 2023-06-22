package gamehandler

import "image"

// contain 包含物件
func contain(target image.Point, array []image.Point) bool {
	for _, v := range array {
		if target == v {
			return true
		}
	}
	return false
}
