package gamehandler

import (
	"image/color"
	"time"

	"github.com/littletrainee/GunGiBoardGameGUI/anotherroundorend"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/capture"
	_color "github.com/littletrainee/GunGiBoardGameGUI/color"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/cpu"
	"github.com/littletrainee/GunGiBoardGameGUI/declaresumi"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/arrangementselect"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/cpuselect"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/phase"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/state"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate/recommendarrangement"
	"github.com/littletrainee/GunGiBoardGameGUI/player"
	"github.com/littletrainee/GunGiBoardGameGUI/timer"
	"github.com/littletrainee/gunginotationgenerator/enum/level"
)

func (g *Game) Update() error {
	switch g.GameState.Phase {

	case phase.SELECT_COLOR:
		g.GameState.SelectColor()

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

	case phase.SELECT_LEVEL:
		g.GameState.SelectLevel()

	case phase.SET_LEVEL:
		// 若選擇的階級是入門或是初級會有推薦或自訂的選擇
		if g.GameState.Level == level.BEGINNER || g.GameState.Level == level.ELEMENTARY {
			g.GameState.Phase = phase.ASK_ARRANGEMENT
			g.GameState.RecommendOrManual = recommendarrangement.Initilization()
			g.GameState.LevelList = nil
		} else {
			g.GameState.Phase = phase.BEGIN_COUNTDOWN_FOR_GAMING
			g.GameState.LevelList = nil
		}
		// 若選擇高級最高段數提高至3
		if g.GameState.Level == level.ADVANCED {
			g.GameState.MaxLayer = 3
		} else {
			g.GameState.MaxLayer = 2
		}
		// 先後手
		g.GameState.SetFirstAndTurn()
		// 設定駒台上駒的位置
		g.Player1.SetKomaTaiPosition(g.GameState.Level, g.Font)
		g.Player2.SetKomaTaiPosition(g.GameState.Level, g.Font)

	case phase.ASK_ARRANGEMENT:
		g.GameState.AskArrangement()

	case phase.SET_COUNTDOWN_FOR_GAMING:
		g.PrepareForGameing = timer.Initilization(constant.PREPARE_FOR_GAMING_TIME, 0, 0)
		g.PrepareForGameing.Op.GeoM.Scale(5, 5)
		g.GameState.Phase = phase.BEGIN_COUNTDOWN_FOR_GAMING
		g.PrepareForGameing.CountDown()

	case phase.BEGIN_COUNTDOWN_FOR_GAMING:
		if g.PrepareForGameing.RemainingTime == 0 {
			g.PrepareForGameing.StopCountDown <- true
			g.GameState.Phase = phase.INITILIZATION_EACH_OBJECT
		}

	case phase.INITILIZATION_EACH_OBJECT:
		g.Board = board.Initilization()
		g.Player1Timer = timer.Initilization(constant.REMAINING_TIME,
			constant.TIMER_X+1,
			constant.TIMER_Y+1+constant.TIMER_HEIGHT/2)
		g.Player2Timer = timer.Initilization(constant.REMAINING_TIME,
			constant.TIMER_X+1,
			constant.TIMER_Y+1)
		g.CPU = cpu.Initilization(&g.Player2)
		g.GameState.Phase = phase.SELECT_KOMA
		g.CurrentState = state.ARRANGEMENT
		if g.GameState.Turn == g.Player1.SelfColor {
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
		g.AnotherRoundOrEnd = anotherroundorend.Initilization(g.Font)
		g.DeclareSuMi = declaresumi.Initilization(g.Font)

	case phase.SELECT_KOMA:
		switch g.CurrentState {
		case state.ARRANGEMENT:
			// 若是使用建議布陣則直接進入對弈階段
			if g.GameState.ArramgementBy == arrangementselect.RECOMMEND {
				g.Player1.RecommendPosition(g.GameState, g.Board)
				g.Player2.RecommendPosition(g.GameState, g.Board)
				g.CurrentState = state.DUELING
				g.DeclareSuMi.Show = false

			} else if g.GameState.Turn == g.GameState.First { // 非建議布陣且動作權為玩家1
				if g.Player1.IsSuMi {
					// continue to opponent
					g.ContinueToOpponent()
				} else {
					g.ArrangementPhaseSelectKoma()
				}
			} else {
				g.CPU.ArrangementPhaseSelectKoma(g.GameState, &g.Board)
				if g.CPU.CurrentDeclareSuMiPercentage > g.CPU.DeclareSuMiTargetPercentage {
					g.delayedChangePhaseTo(phase.CLICK_CLOCK)
				} else {
					g.CPU.CurrentDeclareSuMiPercentage += g.CPU.DeclareSuMiPercentagePhase
					g.delayedChangePhaseTo(phase.MOVE_KOMA)
				}
			}
		case state.DUELING:
			if g.GameState.Turn == g.GameState.First {
				//選擇棋盤上的駒或是駒台上的駒
				g.DuelingPhaseSelectKoma()
			} else {
				g.CPU.DuelingPhaseSelectKoma(g.Board, g.GameState, &g.AnotherRoundOrEnd)
				switch g.CPU.Select {
				case cpuselect.None:
					g.CPU.DuelingPhaseSelectMove(g.GameState, g.Board)
					g.delayedChangePhaseTo(phase.MOVE_KOMA)
				case cpuselect.DEFENSE_CAPTURE,
					cpuselect.DEFENSE_AVOID,
					cpuselect.DEFENSE_ARATA,
					cpuselect.TRY_CAPTURE:
					g.delayedChangePhaseTo(phase.MOVE_KOMA)
				case cpuselect.BEEN_CHECKMATE:
					// 暫停計時器
					g.Player2Timer.StopCountDown <- true
					g.delayedChangePhaseTo(phase.ANOTHER_ROUND_OR_END)
					g.CPU.Select = cpuselect.WAIT_FOR_SELECT_ANOTHERROUND_OR_EXIT
				case cpuselect.WAIT_FOR_SELECT_ANOTHERROUND_OR_EXIT:
				}
			}
		}

	case phase.MOVE_KOMA:
		switch g.CurrentState {
		case state.ARRANGEMENT:
			if g.GameState.Turn == g.GameState.First {
				g.ArrangementPhaseMoveKoma()
			} else {
				g.CPU.ArrangementPhaseMoveKoma(g.Board)
				g.delayedChangePhaseTo(phase.CLICK_CLOCK)
			}
		case state.DUELING:
			if g.GameState.Turn == g.GameState.First {
				g.DuelingPhaseMoveKoma()
			} else {
				if len(g.CPU.MoveToTarget) != 0 {
					g.CPU.DuelingPhaseMoveKoma(&g.Board)
					g.SetMaxRange()
					g.delayedChangePhaseTo(phase.CLICK_CLOCK)
				} else {
					g.delayedChangePhaseTo(phase.ANOTHER_ROUND_OR_END)
				}
			}
		}

	case phase.CLICK_CLOCK:
		if g.GameState.Turn == g.GameState.First {
			g.ClickClock()
		} else {
			g.CPU.ClickClock(&g.GameState, &g.Player1Timer, &g.Player2Timer, g.CurrentState)
			if g.Player2.IsSuMi {
				g.CurrentState = state.DUELING
				g.Player1.KomaDaiBackGroundColor = _color.BoardColor
				g.Player2.KomaDaiBackGroundColor = _color.BoardColor
			}
			g.delayedChangePhaseTo(phase.SELECT_KOMA)
		}

	case phase.PLAYER_CAPTURE_OR_CONTROL_ASK:
		g.CaptureOrControl()

	case phase.ANOTHER_ROUND_OR_END:
		g.AnotherRoundOrEndGame()

	default:
	}
	return nil
}
