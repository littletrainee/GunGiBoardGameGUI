package phase

import (
	"github.com/littletrainee/gunginotationgenerator/enum/phase"
)

const (
	SELECT_COLOR phase.Phase = iota
	SET_COLOR
	SELECT_LEVEL
	SET_LEVEL
	RECOMMENDED_ARRANGEMENT
	SET_TO_PREPARE_FOR_GAMING
	INITILIZATION_BOARD_CLOCK_AND_CPU
	PREPARE_FOR_GAMING
	ARRANGEMENT_PHASE
	DUELING_PHASE_SELECT_KOMA
	DUELING_PHASE_MOVE_KOMA
	DUELING_PHASE_CLICK_CLOCK
	DUELING_PHASE_CAPTURE_OR_CONTROL_ASK
	CPU_SELECT_MOVE
	CPU_MOVE_KOMA
	CPU_CLICK_CLOCK
	DELAY
)
