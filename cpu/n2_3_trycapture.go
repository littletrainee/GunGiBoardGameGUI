package cpu

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

func (c *CPU) TryCapture(b board.Board) {
	for k, v := range b.Blocks {
		var (
			komaStackLength      int = len(v.KomaStack)
			theTopOneOfKomaStack koma.Koma
			confirmPosition      image.Point
			wall                 bool
		)
		if len(v.KomaStack) > 0 && v.KomaStack[komaStackLength-1].Color == c.SelfColor {
			theTopOneOfKomaStack = v.KomaStack[komaStackLength-1]
			for _, direction := range theTopOneOfKomaStack.ProbablyMoveing {
				wall = false
				for i, layer := range direction {
					if i < komaStackLength {
						for _, pos := range layer {
							confirmPosition = image.Point{
								X: theTopOneOfKomaStack.CurrentPosition.X - pos.X,
								Y: theTopOneOfKomaStack.CurrentPosition.Y - pos.Y,
							}
							targetBlock, ok := b.Blocks[confirmPosition]
							if ok {
								if !wall {
									if len(targetBlock.KomaStack) >= komaStackLength {
										wall = true
									}
									for _, v := range targetBlock.KomaStack {
										if v.Color != c.SelfColor {
											c.TargetMove = []image.Point{k, confirmPosition}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
