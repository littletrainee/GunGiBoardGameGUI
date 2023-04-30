package playerstate

type PlayerState int

const (
	WAIT_FOR_SELECT_KOMA PlayerState = iota
	IS_SELECT_KOMA
	WAIT_TO_CLICK_CLOCK
	BLANK
)
