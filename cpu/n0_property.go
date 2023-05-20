package cpu

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/player"
)

type CPU struct {
	*player.Player
	// DeclareSuMiPercentage float32
	// PercentagePhase       float32
	targetPosition image.Point
	targetKoma     []int
	TargetMove     []image.Point
}
