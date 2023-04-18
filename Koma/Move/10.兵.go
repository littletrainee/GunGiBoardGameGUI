package Move

import "github.com/littletrainee/GunGiBoardGameGUI/Coordinate"

func PM兵() [][]Coordinate.Coordinate {
	return [][]Coordinate.Coordinate{
		// 一段
		{{X: 0, Y: -1}, {X: 0, Y: 1}},

		// 二段
		{{X: 0, Y: -2}, {X: 0, Y: 2}},

		// 三段
		{{X: 0, Y: -3}, {X: 0, Y: 3}}}
}
