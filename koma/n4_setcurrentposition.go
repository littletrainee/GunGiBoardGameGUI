package koma

import "github.com/littletrainee/pair"

func (k *Koma) SetCurrentPosition(x, y int) {
	k.CurrentPosition = pair.Pair[int, int]{Item1: x, Item2: y}
}
