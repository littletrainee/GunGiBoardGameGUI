package gamestate

import (
	"log"
	"testing"
)

func TestGameState_SetFirst(t *testing.T) {
	var G GameState = GameState{}
	G.SetFirst()
	log.Println(G.ColorHolder.Own)
}
