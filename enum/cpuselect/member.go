package cpuselect

type CPUSelect int

const (
	None                                 CPUSelect = iota // CPU未選取
	DEFENSE_CAPTURE                                       // CPU選擇透俘獲防守
	DEFENSE_AVOID                                         // CPU選擇移動帥防守
	DEFENSE_ARATA                                         // CPU選擇"新"防守
	TRY_CAPTURE                                           // CPU嘗試俘獲對方的駒
	BEEN_CHECKMATE                                        // CPU被"將軍"
	WAIT_FOR_SELECT_ANOTHERROUND_OR_EXIT                  // 等待玩家選擇再來一局還是離開遊戲
)
