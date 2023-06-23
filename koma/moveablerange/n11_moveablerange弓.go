package moveablerange

import "image"

// MoveableRange弓 弓可以移動的範圍，回傳的是三維切片
func MoveableRange弓() [][][]image.Point {
	return [][][]image.Point{
		// 上
		{{{X: 0, Y: -2}}, {{X: 0, Y: -3}}, {{X: 0, Y: -4}}},
		// 右上
		{{{X: 1, Y: -2}}, {{X: 2, Y: -3}}, {{X: 3, Y: -4}}},
		// 右
		{},
		// 右下
		{},
		// 下
		{{{X: 0, Y: 1}}, {{X: 0, Y: 2}}, {{X: 0, Y: 3}}},
		// 左下
		{},
		// 左
		{},
		// 左上
		{{{X: -1, Y: -2}}, {{X: -2, Y: -3}}, {{X: -3, Y: -4}}},
	}
}
