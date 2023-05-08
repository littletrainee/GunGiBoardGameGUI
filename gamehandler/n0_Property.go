package gamehandler

import (
	"image"

	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/cpu"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/player"
	"github.com/littletrainee/GunGiBoardGameGUI/timer"
	"golang.org/x/image/font"
)

type Game struct {
	GameState                  gamestate.GameState
	Player1                    player.Player
	Player2                    player.Player
	PrepareForGameing          timer.Timer
	Board                      board.Board
	Player1Timer               timer.Timer
	Player2Timer               timer.Timer
	Font                       font.Face
	CPU                        cpu.CPU
	WhichBlockBeSelect         image.Point
	WhichKomaOnKomaTaiBeSelect int
}
