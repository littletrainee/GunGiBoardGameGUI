package gamestate

import (
	"fmt"
	"testing"
)

func TestGameState_SetFirst(t *testing.T) {
	var G GameState = GameState{}
	G.SetFirst()
	fmt.Println(G.ColorHolder.Own)
}
