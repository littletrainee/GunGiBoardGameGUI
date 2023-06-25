// 整體程式的主要控制項
package gamehandler

import (
	"image"

	aroe "github.com/littletrainee/GunGiBoardGameGUI/anotherroundorend"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/capture"
	"github.com/littletrainee/GunGiBoardGameGUI/cpu"
	"github.com/littletrainee/GunGiBoardGameGUI/declaresumi"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/state"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/mutiny"
	"github.com/littletrainee/GunGiBoardGameGUI/player"
	"github.com/littletrainee/GunGiBoardGameGUI/timer"
	"golang.org/x/image/font"
)

// 程式相關的數據管理
type Game struct {
	GameState             gamestate.GameState     // 遊戲狀態
	Font                  font.Face               // 顯示字體
	Player1               player.Player           // 玩家1
	Player2               player.Player           // 玩家2
	PrepareForGameing     timer.Timer             // 準備進入布陣階段的倒數計時器
	Board                 board.Board             // 棋盤
	Player1Timer          timer.Timer             // 玩家1的棋鐘
	Player2Timer          timer.Timer             // 玩家2的棋鐘
	CPU                   cpu.CPU                 // 電腦
	CurrentState          state.State             // 當前的階段
	WhichKomaBeenSelected []int                   // 哪個駒被選取
	TargetPosition        image.Point             // 目標位置
	DeclareSuMi           declaresumi.DeclareSuMi // 宣告布陣完畢
	Capture               capture.Capture         // 俘獲視窗
	ConfirmPositionSlice  []image.Point           // 核可的位置
	AnotherRoundOrEnd     aroe.AnotherRoundOrEnd  // 遊戲結束或是再來一局視窗
	Mutiny                mutiny.Mutiny           // 叛變視窗
}
