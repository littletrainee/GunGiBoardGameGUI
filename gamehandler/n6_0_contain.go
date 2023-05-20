package gamehandler

import "image"

func contain(target image.Point, array []image.Point) bool {
	for _, v := range array {
		if target == v {
			return true
		}
	}
	return false
}
