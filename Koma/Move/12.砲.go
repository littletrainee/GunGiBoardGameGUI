package Move

import "github.com/littletrainee/GunGiBoardGameGUI/Coordinate"

func PM砲() [][]Coordinate.Coordinate {
	return [][]Coordinate.Coordinate{
		// 一段
		{{X: 0, Y: -3}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}},

		// 二段
		{{X: 0, Y: -4}, {X: 2, Y: 0}, {X: 0, Y: 2}, {X: -2, Y: 0}},

		// 三段
		{{X: 0, Y: -5}, {X: 3, Y: 0}, {X: 0, Y: 3}, {X: -3, Y: 0}}}
}
