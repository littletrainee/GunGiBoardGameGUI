package cpu

import (
	"github.com/littletrainee/GunGiBoardGameGUI/player"
)

func Initilization(p *player.Player) CPU {
	var c CPU = CPU{
		Player: p,
	}
	// c.Player = p
	// var count int
	// for _, v := range p.KomaTai {
	// 	count += v.Item2
	// }
	// c.PercentagePhase = float32(100 / count)
	// c.DeclareSuMiPercentage = c.PercentagePhase
	// c.target = rand.Float32() * 100.0
	return c
}
