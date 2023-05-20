package koma

import (
	"image"

	"github.com/littletrainee/pair"
)

func (k *Koma) SetCurrentPosition(p pair.Pair[int, int]) {
	k.CurrentPosition = image.Point{X: p.Item1, Y: p.Item2}
}
