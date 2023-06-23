package moveablerange

import "image"

// MoveableRange馬 馬可以移動的範圍，回傳的是三維切片
func MoveableRange馬() [][][]image.Point {
	return [][][]image.Point{
		// 上
		{{{X: 0, Y: -1}, {X: 0, Y: -2}}, {{X: 0, Y: -3}}, {{X: 0, Y: -4}}},
		// 右上
		{},
		// 右
		{{{X: 1, Y: 0}}, {{X: 2, Y: 0}}, {{X: 3, Y: 0}}},
		// 右下
		{},
		// 下
		{{{X: 0, Y: 1}, {X: 0, Y: 2}}, {{X: 0, Y: 3}}, {{X: 0, Y: 4}}},
		// 左下
		{},
		// 左
		{{{X: -1, Y: 0}}, {{X: -2, Y: 0}}, {{X: -3, Y: 0}}},
		// 左上
		{},
	}
}
