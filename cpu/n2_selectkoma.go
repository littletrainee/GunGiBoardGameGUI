package cpu

import (
	"math/rand"
	"time"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
)

func (c *CPU) SelectKoma(b board.Board) {
	var probablyChoice [][]int
	rand.Seed(time.Now().UnixNano())
	for k, v := range b.Blocks {
		if len(v.KomaStack) > 0 && v.KomaStack[len(v.KomaStack)-1].Color == c.Player.SelfColor {
			probablyChoice = append(probablyChoice, []int{k.X, k.Y})
		}
	}

	for i, v := range c.Player.KomaTai {
		if v.Item2 > 0 {
			probablyChoice = append(probablyChoice, []int{i})
		}
	}
	c.targetKoma = probablyChoice[rand.Intn(len(probablyChoice))]
	// c.targetKoma = []int{3}

}
