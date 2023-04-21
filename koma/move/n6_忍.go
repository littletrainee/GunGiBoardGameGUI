package move

import "github.com/littletrainee/GunGiBoardGameGUI/coordinate"

func PM忍() [][]coordinate.Coordinate {
	return [][]coordinate.Coordinate{
		// 一段
		{{X: -1, Y: -1}, {X: -2, Y: -2}, {X: 1, Y: -1}, {X: 2, Y: -2},
			{X: 1, Y: 1}, {X: 2, Y: 2}, {X: -1, Y: 1}, {X: -2, Y: 2}},

		// 二段
		{{X: -3, Y: -3}, {X: 3, Y: -3}, {X: 3, Y: 3}, {X: -3, Y: 3}},

		// 三段
		{{X: -4, Y: -4}, {X: 4, Y: -4}, {X: 4, Y: 4}, {X: -4, Y: 4}}}
}
