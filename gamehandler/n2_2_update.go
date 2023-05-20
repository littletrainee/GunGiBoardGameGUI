package gamehandler

import (
	"image/color"
	"time"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/capture"
	_color "github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/cpu"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/recommendarrangement"
	"github.com/littletrainee/GunGiBoardGameGUI/player"
	"github.com/littletrainee/GunGiBoardGameGUI/timer"
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
		if g.GameState.Level == level.ADVANCED {
			g.GameState.MaxLayer = 3
		} else {
			g.GameState.MaxLayer = 2
		}
		g.GameState.SetFirstAndTurn()
		g.Player1.SetKomaTaiPosition(g.GameState.Level, g.Font)
		g.Player2.SetKomaTaiPosition(g.GameState.Level, g.Font)

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

	// 初始化棋盤、棋鐘、電腦與詢問俘獲彈出視窗
	case phase.INITILIZATION_BOARD_CLOCK_AND_CPU:
		g.Board = board.Initilization()
		g.Player1Timer = timer.Initilization(constant.REMAINING_TIME,
			constant.TIMER_POSITION_X+1,
			constant.TIMER_POSITION_Y+1+constant.TIMER_BASE_HEIGHT/2)
		g.Player2Timer = timer.Initilization(constant.REMAINING_TIME,
			constant.TIMER_POSITION_X+1,
			constant.TIMER_POSITION_Y+1)
		// g.TimerHandler = timerhandler.Initilization(g.GameState)
		g.CPU = cpu.Initilization(&g.Player2)
		g.GameState.Phase = phase.ARRANGEMENT_PHASE
		if g.GameState.Turn == g.Player1.SelfColor {
			// g.TimerHandler.Player1Timer.BackgroundColor = _color.ConfirmColor
			g.Player1Timer.BackgroundColor = _color.ConfirmColor
			g.Player2Timer.BackgroundColor = _color.TimerPauseColor
			g.Player1Timer.SetGeoM(false)
			g.Player2Timer.SetGeoM(true)

			g.Player1Timer.CountDown()
			g.Player2Timer.CountDown()
			go func() {
				time.Sleep(time.Second / 2)
				g.Player2Timer.StopCountDown <- true
			}()
		} else {
			g.Player1Timer.BackgroundColor = _color.TimerPauseColor
			g.Player2Timer.BackgroundColor = _color.ConfirmColor
			g.Player1Timer.SetGeoM(false)
			g.Player2Timer.SetGeoM(true)

			g.Player2Timer.CountDown()
			g.Player1Timer.CountDown()
			go func() {
				time.Sleep(time.Second / 2)
				g.Player1Timer.StopCountDown <- true
			}()
		}
		g.Capture = capture.Initilization(g.Font)
	// 布陣階段
	case phase.ARRANGEMENT_PHASE:
		if g.GameState.RecommendedArramgement {
			g.Player1.DefaultPosition(g.GameState, g.Board)
			g.Player2.DefaultPosition(g.GameState, g.Board)
			g.delayedChangePhaseTo(phase.DUELING_PHASE_SELECT_KOMA)
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
	case phase.DUELING_PHASE_SELECT_KOMA:
		if g.GameState.Turn == g.Player1.SelfColor {
			//選擇棋盤上的駒或是駒台上的駒
			g.SelectKoma()
		} else {
			g.CPU.SelectKoma(g.Board)
			if len(g.CPU.TargetMove) != 0 {
				g.delayedChangePhaseTo(phase.CPU_MOVE_KOMA)
			} else {
				g.delayedChangePhaseTo(phase.CPU_SELECT_MOVE)
			}
		}

	case phase.CPU_SELECT_MOVE:
		g.CPU.SelectMove(g.GameState, g.Board)
		g.delayedChangePhaseTo(phase.CPU_MOVE_KOMA)
	case phase.CPU_MOVE_KOMA:
		g.CPU.MoveKoma(g.Board)
		g.delayedChangePhaseTo(phase.CPU_CLICK_CLOCK)
	case phase.CPU_CLICK_CLOCK:
		g.CPU.ClickClock(&g.GameState, &g.Player1Timer, &g.Player2Timer)
		g.delayedChangePhaseTo(phase.DUELING_PHASE_SELECT_KOMA)
	case phase.DUELING_PHASE_MOVE_KOMA:
		g.MoveKoma()
	case phase.DUELING_PHASE_CAPTURE_OR_CONTROL_ASK:
		g.CaptureOrControl()

	case phase.DUELING_PHASE_CLICK_CLOCK:
		g.ClickClock()
	default:
	}
	return nil
}
