package moveablerange

import "image"

// MoveableRange筒 筒可以移動的範圍，回傳的是三維切片
func MoveableRange筒() [][][]image.Point {
	return [][][]image.Point{
		// 上
		{{{X: 0, Y: -2}}, {{X: 0, Y: -3}}, {{X: 0, Y: -4}}},
		// 右上
		{},
		// 右
		{},
		// 右下
		{{{X: 1, Y: 1}}, {{X: 2, Y: 2}}, {{X: 3, Y: 3}}},
		// 下
		{},
		// 左下
		{{{X: -1, Y: 1}}, {{X: -2, Y: 2}}, {{X: -3, Y: 3}}},
		// 左
		{},
		// 左上
		{},
	}
}
