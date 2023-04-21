package gamehandler

import (
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/player"
	"golang.org/x/image/font"
)

type Game struct {
	Own       player.Player
	Opponent  player.Player
	Board     board.Board
	GameState gamestate.GameState
	Font      font.Face
}
