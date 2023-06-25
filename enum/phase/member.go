// 程式執行階段的列舉
package phase

import (
	"github.com/littletrainee/gunginotationgenerator/enum/phase"
)

const (
	SELECT_COLOR                           phase.Phase = iota // 選擇顏色
	SET_COLOR                                                 // 設定兩家的顏色
	SELECT_LEVEL                                              // 選擇階級
	SET_LEVEL                                                 // 設定階級
	SELECT_RECOMMEND_OR_MANUAL_ARRANGEMENT                    // 詢問布陣
	SET_COUNTDOWN_FOR_GAMING                                  // 設定對弈開始倒數計時器
	BEGIN_COUNTDOWN_FOR_GAMING                                // 對弈前的倒數計時
	INITILIZATION_EACH_OBJECT                                 // 初始化棋盤、棋鐘、電腦與詢問俘獲彈出視窗
	SELECT_KOMA                                               // 玩家與電腦的選駒
	MOVE_KOMA                                                 // 玩家的移駒
	CLICK_CLOCK                                               // 玩家的棋鐘
	PLAYER_CAPTURE_OR_CONTROL_ASK                             // 玩家的俘獲或控制視窗
	MUTINY                                                    // 玩家選擇是否要叛變對方的駒
	ANOTHER_ROUND_OR_END                                      // 詢問是否再來一局
	BEFORE_GAMING_BUFFER_ZONE                                 // 遊戲開始前的緩衝
	GAMING_BUFFER_ZONE                                        // 遊戲開始後緩衝區
)
