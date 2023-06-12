package defaultposition

import (
	"image"

	"github.com/littletrainee/pair"
)

func OpponentElementary() []pair.Pair[int, image.Point] {
	return []pair.Pair[int, image.Point]{
		{Item1: 0, Item2: image.Point{X: 5, Y: 1}},

		{Item1: 1, Item2: image.Point{X: 4, Y: 1}},

		{Item1: 2, Item2: image.Point{X: 6, Y: 1}},

		{Item1: 4, Item2: image.Point{X: 6, Y: 3}},
		{Item1: 4, Item2: image.Point{X: 4, Y: 3}},

		{Item1: 5, Item2: image.Point{X: 2, Y: 2}},

		{Item1: 6, Item2: image.Point{X: 5, Y: 2}},

		{Item1: 7, Item2: image.Point{X: 7, Y: 3}},
		{Item1: 7, Item2: image.Point{X: 3, Y: 3}},

		{Item1: 8, Item2: image.Point{X: 8, Y: 2}},

		{Item1: 9, Item2: image.Point{X: 9, Y: 3}},
		{Item1: 9, Item2: image.Point{X: 5, Y: 3}},
		{Item1: 9, Item2: image.Point{X: 1, Y: 3}},

		{Item1: 10, Item2: image.Point{X: 7, Y: 2}},
		{Item1: 10, Item2: image.Point{X: 3, Y: 2}},
	}
}
