package moveablerange

import (
	"image"
)

func MoveableRange謀() [][][]image.Point {
	return [][][]image.Point{
		// 上
		{},
		// 右上
		{{{X: 1, Y: -1}}, {{X: 2, Y: -2}}, {{X: 3, Y: -3}}},
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
		{{{X: -1, Y: -1}}, {{X: -2, Y: -2}}, {{X: -3, Y: -3}}},
	}
}
