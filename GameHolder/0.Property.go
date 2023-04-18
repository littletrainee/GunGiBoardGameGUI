package GameHolder

import (
	"github.com/littletrainee/GunGiBoardGameGUI/Board"
	"github.com/littletrainee/GunGiBoardGameGUI/GameState"
	"github.com/littletrainee/GunGiBoardGameGUI/Player"
	"golang.org/x/image/font"
)

type Game struct {
	Own       Player.Player
	Opponent  Player.Player
	Board     Board.Board
	GameState GameState.GameState
	Font      font.Face
}
