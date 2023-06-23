// 電腦物件，使城市可以自動控制駒的物件
package cpu

import (
	"github.com/littletrainee/GunGiBoardGameGUI/enum/cpuselect"
	"github.com/littletrainee/GunGiBoardGameGUI/player"
)

// 電腦物件
type CPU struct {
	*player.Player
	CurrentDeclareSuMiPercentage float32             // 宣告布陣完畢的當前機率
	DeclareSuMiPercentagePhase   float32             // 宣告布陣完畢的區間值
	DeclareSuMiTargetPercentage  float32             // 宣告布陣完畢的機率
	MoveToTarget                 []int               // 從何處移動到目標位置
	Select                       cpuselect.CPUSelect // CPU的當前選擇
}
