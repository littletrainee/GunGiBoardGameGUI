package cpu

import (
	"github.com/littletrainee/GunGiBoardGameGUI/enum/cpuselect"
	"github.com/littletrainee/GunGiBoardGameGUI/player"
)

type CPU struct {
	*player.Player
	DeclareSuMiTargetPercentage  float32 // 宣告布陣完畢的機率
	DeclareSuMiPercentagePhase   float32 // 宣告布陣完畢的區間值
	CurrentDeclareSuMiPercentage float32 // 宣告布陣完畢的當前機率
	// checkmateBy                  image.Point         // 哪個位置可以將軍
	MoveToTarget []int               // 從何處移動到目標位置
	Select       cpuselect.CPUSelect // CPU的當前選擇
}
