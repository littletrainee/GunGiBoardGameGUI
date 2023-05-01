package gamestate

import (
	"fmt"
	"testing"
)

func TestGameState_SetFirst(t *testing.T) {
	var G GameState = GameState{}
	G.SetFirstAndTurn()
	fmt.Println(G.First)
}
