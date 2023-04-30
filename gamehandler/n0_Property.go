package gamehandler

import (
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/cpu"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/player"
	"github.com/littletrainee/GunGiBoardGameGUI/timer"
	"github.com/littletrainee/GunGiBoardGameGUI/timerhandler"
	"golang.org/x/image/font"
)

type Game struct {
	GameState         gamestate.GameState
	Player1           player.Player
	Player2           player.Player
	PrepareForGameing timer.Timer
	Board             board.Board
	// TimerPair         pair.Pair[timer.Timer, timer.Timer]
	TimerHandler timerhandler.TimerHandler
	Font         font.Face
	CPU          cpu.CPU
}
