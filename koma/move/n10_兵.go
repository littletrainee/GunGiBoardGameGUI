package move

import "github.com/littletrainee/GunGiBoardGameGUI/coordinate"

func PM兵() [][]coordinate.Coordinate {
	return [][]coordinate.Coordinate{
		// 一段
		{{X: 0, Y: -1}, {X: 0, Y: 1}},

		// 二段
		{{X: 0, Y: -2}, {X: 0, Y: 2}},

		// 三段
		{{X: 0, Y: -3}, {X: 0, Y: 3}}}
}
