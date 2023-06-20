package cpu

import (
	"math/rand"
	"time"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/player"
)

func Initilization(p *player.Player) CPU {
	var c CPU = CPU{
		Player: p,
	}
	c.Player = p
	var count int
	for _, v := range p.KomaDai {
		count += v.Item2
	}
	c.DeclareSuMiPercentagePhase = float32(100 / count)
	rand.Seed(time.Now().UnixNano())
	c.DeclareSuMiTargetPercentage = float32(rand.Intn(constant.CPU_DECLARE_SUMI_TARGET_PERCENTAGE))
	return c
}
