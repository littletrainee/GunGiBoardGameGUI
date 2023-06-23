package moveablerange

import (
	"image"
)

// MoveableRange兵 兵可以移動的範圍，回傳的是三維切片
func MoveableRange兵() [][][]image.Point {
	return [][][]image.Point{
		// 上
		{{{X: 0, Y: -1}}, {{X: 0, Y: -2}}, {{X: 0, Y: -3}}},
		// 右上
		{},
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
		{},
	}
}
