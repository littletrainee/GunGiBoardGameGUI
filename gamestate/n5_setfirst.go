package gamestate

import (
	"image/color"
)

// SetFirst設定先後手
func (g *GameState) SetFirst() {
	// rand.Seed(time.Now().UnixNano())
	// target := rand.Intn(2)
	// if target == 0 {
	// 	g.First = color.Black
	// 	g.Turn = color.Black
	// } else {
	g.ColorHolder.Own = color.White
	g.ColorHolder.Turn = color.White
	// }
}
