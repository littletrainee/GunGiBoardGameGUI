package Move

import "github.com/littletrainee/GunGiBoardGameGUI/Coordinate"

func PM大() [][]Coordinate.Coordinate {
	return [][]Coordinate.Coordinate{
		// 一段
		{{X: -1, Y: -1}, {X: 1, Y: -1}, {X: 1, Y: 1}, {X: -1, Y: 1}},

		// 二段
		{{X: -2, Y: -2}, {X: 2, Y: -2}, {X: 2, Y: 2}, {X: -2, Y: 2}},

		// 三段
		{{X: -3, Y: -3}, {X: 3, Y: -3}, {X: 3, Y: 3}, {X: -3, Y: 3}},

		// 無視段
		{{X: 0, Y: -1}, {X: 0, Y: -2}, {X: 0, Y: -3}, {X: 0, Y: -4},
			{X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}, {X: 4, Y: 0},
			{X: 0, Y: 1}, {X: 0, Y: 2}, {X: 0, Y: 3}, {X: 0, Y: 4},
			{X: -1, Y: 0}, {X: -2, Y: 0}, {X: -3, Y: 0}, {X: -4, Y: 0}}}
}
