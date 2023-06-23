package moveablerange

import "image"

// MoveableRange砦 砦可以移動的範圍，回傳的是三維切片
func MoveableRange砦() [][][]image.Point {
	return [][][]image.Point{
		// 上
		{{{X: 0, Y: -1}}, {{X: 0, Y: -2}}, {{X: 0, Y: -3}}},
		// 右上
		{},
		// 右
		{{{X: 1, Y: 0}}, {{X: 2, Y: 0}}, {{X: 3, Y: 0}}},
		// 右下
		{{{X: 1, Y: 1}}, {{X: 2, Y: 2}}, {{X: 3, Y: 3}}},
		// 下
		{},
		// 左下
		{{{X: -1, Y: 1}}, {{X: -2, Y: 2}}, {{X: -3, Y: 3}}},
		// 左
		{{{X: -1, Y: 0}}, {{X: -2, Y: 0}}, {{X: -3, Y: 0}}},
		// 左上
		{},
	}
}
