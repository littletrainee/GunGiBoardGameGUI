package gamehandler

func (g *Game) SetOwnMaxRange() {
	for k, v := range g.Board.Blocks {
		for _, e := range v.KomaStack {
			if e.Color == g.Player1.SelfColor && g.Player1.MaxRange < k.X {
				g.Player1.MaxRange = k.X
				break
			}
		}
	}

}
