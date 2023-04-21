package move

import "github.com/littletrainee/GunGiBoardGameGUI/coordinate"

func PM槍() [][]coordinate.Coordinate {
	return [][]coordinate.Coordinate{
		// 一段
		{{X: -1, Y: -1}, {X: 0, Y: -1}, {X: 0, Y: -2}, {X: 1, Y: -1}, {X: 0, Y: 1}},

		// 二段
		{{X: -2, Y: -2}, {X: 0, Y: -3}, {X: 2, Y: -2}, {X: 0, Y: 2}},

		// 三段
		{{X: -3, Y: -3}, {X: 0, Y: -4}, {X: 3, Y: -3}, {X: 0, Y: 3}}}
}
