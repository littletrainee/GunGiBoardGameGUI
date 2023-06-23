package moveablerange

import (
	"image"
)

// MoveableRange忍 忍可以移動的範圍，回傳的是三維切片
func MoveableRange忍() [][][]image.Point {
	return [][][]image.Point{
		// 上
		{},
		// 右上
		{{{X: 1, Y: -1}, {X: 2, Y: -2}}, {{X: 3, Y: -3}}, {{X: 4, Y: -4}}},
		// 右
		{},
		// 右下
		{{{X: 1, Y: 1}, {X: 2, Y: 2}}, {{X: 3, Y: 3}}, {{X: 4, Y: 4}}},
		// 下
		{},
		// 左下
		{{{X: -1, Y: 1}, {X: -2, Y: 2}}, {{X: -3, Y: 3}}, {{X: -4, Y: 4}}},
		// 左
		{},
		// 左上
		{{{X: -1, Y: -1}, {X: -2, Y: -2}}, {{X: -3, Y: -3}}, {{X: -4, Y: -4}}},
	}
}
