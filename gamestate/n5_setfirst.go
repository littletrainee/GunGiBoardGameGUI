package gamestate

import (
	"image/color"
)

// 設定先後手
func (g *GameState) SetFirstAndTurn() {
	// rand.Seed(time.Now().UnixNano())
	// target := rand.Intn(2)
	// if target == 0 {
	// 	g.First = color.Black
	// 	g.Turn = color.Black
	// } else {
	g.First = color.White
	g.Turn = color.White
	// }
}
