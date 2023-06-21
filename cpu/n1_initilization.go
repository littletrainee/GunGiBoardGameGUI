package cpu

import (
	"math/rand"
	"time"

	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/player"
)

// Initilization 初始化CPU物件，設定好本局宣告"布陣完成"的機率與繼承控制Player2的物件並回傳。
//
// 參數p用為CPU繼承控制的Player2
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
