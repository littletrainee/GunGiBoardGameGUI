package defaultposition

import (
	"image"

	"github.com/littletrainee/pair"
)

/*
0,5.9
1,6.9
2.4.9
4,6.7
4.4.7
5.8.8
5.2.8
6.5.8
7.7.7
7.3.7
9.9.7
9.5.7
9.1.7
*/

func OwnBeginner() []pair.Pair[int, image.Point] {
	return []pair.Pair[int, image.Point]{
		{Item1: 0, Item2: image.Point{X: 5, Y: 9}},

		{Item1: 1, Item2: image.Point{X: 6, Y: 9}},

		{Item1: 2, Item2: image.Point{X: 4, Y: 9}},

		{Item1: 4, Item2: image.Point{X: 6, Y: 7}},
		{Item1: 4, Item2: image.Point{X: 4, Y: 7}},

		{Item1: 5, Item2: image.Point{X: 8, Y: 8}},
		{Item1: 5, Item2: image.Point{X: 2, Y: 8}},

		{Item1: 6, Item2: image.Point{X: 5, Y: 8}},

		{Item1: 7, Item2: image.Point{X: 7, Y: 7}},
		{Item1: 7, Item2: image.Point{X: 3, Y: 7}},

		{Item1: 9, Item2: image.Point{X: 9, Y: 7}},
		{Item1: 9, Item2: image.Point{X: 5, Y: 7}},
		{Item1: 9, Item2: image.Point{X: 1, Y: 7}},
	}
}
