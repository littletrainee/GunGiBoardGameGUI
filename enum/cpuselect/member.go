package cpuselect

type CPUSelect int

const (
	// CPU未選取
	None CPUSelect = iota
	// CPU選擇透俘獲防守
	DEFENSE_CAPTURE
	// CPU選擇移動帥防守
	DEFENSE_AVOID
	// CPU選擇"新"防守
	DEFENSE_ARATA
	// CPU嘗試俘獲對方的駒
	TRY_CAPTURE
	// CPU被"將軍"
	BEEN_CHECKMATE
)
