package move

import (
	"image"
)

func PM兵() [][][]image.Point {
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
