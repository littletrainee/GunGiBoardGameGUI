package gamehandler

import (
	"fmt"
	"image/color"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/cpu"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/recommendarrangement"
	"github.com/littletrainee/GunGiBoardGameGUI/player"
	"github.com/littletrainee/GunGiBoardGameGUI/timer"
	"github.com/littletrainee/GunGiBoardGameGUI/timerhandler"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
)

func (g *Game) Update() error {
	switch g.GameState.Phase {
	// 選擇顏色
	case phase.SELECT_COLOR:
		g.GameState.SelectColor()

	// 設定顏色
	case phase.SET_COLOR:
		// 設定顏色與是否為自家，並初始化駒台切片
		if g.GameState.First == color.Black {
			g.Player1 = player.Initilization(color.Black, true)
			g.Player2 = player.Initilization(color.White, false)
		} else {
			g.Player1 = player.Initilization(color.White, true)
			g.Player2 = player.Initilization(color.Black, false)
		}
		g.GameState.Phase = phase.SELECT_LEVEL

	// 選擇階級
	case phase.SELECT_LEVEL:
		g.GameState.SelectLevel()

	// 設定階級與初始化雙方的駒台切片元素
	case phase.SET_LEVEL:
		if g.GameState.Level == level.BEGINNER ||
			g.GameState.Level == level.ELEMENTARY {
			g.GameState.Phase = phase.RECOMMENDED_ARRANGEMENT
			g.GameState.ArrangementList = recommendarrangement.Initilization()
			g.GameState.LevelList = nil
		} else {
			g.GameState.Phase = phase.PREPARE_FOR_GAMING
			g.GameState.LevelList = nil
		}
		g.GameState.SetFirst()
		g.Player1.SetPositionAndState(g.GameState.Level)
		g.Player2.SetPositionAndState(g.GameState.Level)

	// 若階級是入門或初級詢問是否用推薦的配置
	case phase.RECOMMENDED_ARRANGEMENT:
		g.GameState.SetRecommendedArrangement()

	// 設定對弈前的倒數計時
	case phase.SET_TO_PREPARE_FOR_GAMING:
		g.PrepareForGameing = timer.Initilization(
			constant.PREPARE_FOR_GAMING_TIME,
			0, 0)
		g.PrepareForGameing.Op.GeoM.Scale(5, 5)
		g.GameState.Phase = phase.PREPARE_FOR_GAMING
		g.PrepareForGameing.CountDown()

	// 對弈前的倒數計時
	case phase.PREPARE_FOR_GAMING:
		if g.PrepareForGameing.RemainingTime == 0 {
			g.PrepareForGameing.StopCountDown <- true
			g.GameState.Phase = phase.INITILIZATION_BOARD_CLOCK_AND_CPU
		}

	// 初始化棋盤、騎鐘與電腦
	case phase.INITILIZATION_BOARD_CLOCK_AND_CPU:
		g.Board = board.Initilization()
		g.TimerHandler = timerhandler.Initilization()
		g.CPU = cpu.Initilization(&g.Player2)
		g.GameState.Phase = phase.ARRANGEMENT_PHASE

	// 布陣階段
	case phase.ARRANGEMENT_PHASE:
		if g.GameState.RecommendedArramgement {
			g.Player1.DefaultPosiont(g.GameState, g.Board)
			g.Player2.DefaultPosiont(g.GameState, g.Board)
			g.GameState.Phase = phase.DUELING_PHASE
		}
		// if g.GameState.Turn == g.GameState.First {
		// 	// 確認本家是否已經有動作
		// 	switch g.Player1.PlayerState {
		// 	case playerstate.WAIT_FOR_SELECT_KOMA:
		// 		g.Player1.SelectKomaFromKomaTai(g.GameState, g.Board)
		// 	case playerstate.IS_SELECT_KOMA:
		// 		g.Player1.PutOnBoard(g.Board, g.GameState)
		// 	case playerstate.WAIT_TO_CLICK_CLOCK:
		// 		g.TimerHandler.ChangeAnotherOne(&g.GameState)
		// 		g.Player2.PlayerState = playerstate.IS_SELECT_KOMA
		// 	}
		// } else {
		// 	switch g.Player2.PlayerState {
		// 	case playerstate.IS_SELECT_KOMA:
		// 		g.cpu.AutoSetKoma(g.GameState, &g.Board)
		// 		time.Sleep(time.Second)
		// 	case playerstate.WAIT_TO_CLICK_CLOCK:
		// 		g.cpu.ClickClockOrSuMi(&g.TimerHandler, &g.GameState)
		// 		time.Sleep(time.Second)
		// 		g.Player1.PlayerState = playerstate.WAIT_FOR_SELECT_KOMA
		// 		g.GameState.ArrangementPhaseRoundCount++
		// 	}
		// }
	case phase.DUELING_PHASE:
		fmt.Println("Hello")
	default:
	}
	// }
	return nil
}