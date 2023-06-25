// CPU演算後的決策列舉
package cpuselect

type CPUSelect int

const (
	None            CPUSelect = iota // CPU未選取
	DEFENSE_CAPTURE                  // CPU選擇透俘獲防守
	DEFENSE_AVOID                    // CPU選擇移動帥防守
	DEFENSE_ARATA                    // CPU選擇"新"防守
	TRY_CAPTURE                      // CPU嘗試俘獲對方的駒
	BEEN_CHECKMATE                   // CPU被"將軍"
	RANDOM_SELECT                    // CPU一般的選擇
	CHECKMATE                        // CPU"將軍"玩家
)
