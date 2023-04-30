package cpu

import "github.com/littletrainee/GunGiBoardGameGUI/player"

type CPU struct {
	Player                *player.Player
	DeclareSuMiPercentage float32
	PercentagePhase       float32
	target                float32
}
