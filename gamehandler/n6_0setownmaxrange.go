package gamehandler

func (g *Game) SetOwnMaxRange() {
	for k, v := range g.Board.Blocks {
		if len(v.KomaStack) != 0 && v.KomaStack[len(v.KomaStack)-1].Color == g.Player1.SelfColor && k.Y < g.Player1.MaxRange {
			g.Player1.MaxRange = k.Y
		}
	}
}
