package cpu

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/player"
)

type CPU struct {
	*player.Player
	// DeclareSuMiPercentage float32
	// PercentagePhase       float32
	checkmateBy          image.Point
	targetPosition       image.Point
	targetKoma           []int
	CaptureForDefense    []image.Point
	AvoidForDefense      []image.Point
	ARaTaForDefense      []int
	CaptureForMotivation []image.Point
}
