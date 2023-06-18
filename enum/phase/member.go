package phase

import (
	"github.com/littletrainee/gunginotationgenerator/enum/phase"
)

const (
	// 選擇顏色
	SELECT_COLOR phase.Phase = iota
	// 設定顏色
	SET_COLOR
	// 選擇階級
	SELECT_LEVEL
	// 設定階級
	SET_LEVEL
	// 詢問布陣
	ASK_ARRANGEMENT
	// 設定對弈開始倒數計時器
	SET_COUNTDOWN_FOR_GAMING
	// 對弈前的倒數計時
	BEGIN_COUNTDOWN_FOR_GAMING
	// 初始化棋盤、棋鐘、電腦與詢問俘獲彈出視窗
	INITILIZATION_EACH_OBJECT
	// 玩家與電腦的選駒
	SELECT_KOMA
	// 玩家的移駒
	MOVE_KOMA
	// 玩家的棋鐘
	CLICK_CLOCK
	// 玩家的俘獲或控制視窗
	PLAYER_CAPTURE_OR_CONTROL_ASK
	// 詢問是否再來一局
	ANOTHER_ROUND_OR_END
	// 緩衝區
	BUFFER_ZONE
)
