package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/playerstate"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
)

// 駒未被選取時
func (p *Player) SelectKomaFromKomaTai(GS gamestate.GameState, b board.Board) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		// 非第一巡
		if GS.ArrangementPhaseRoundCount != 1 {
			for i, v := range p.KomaTai {
				if v.Item1.OnKoma(float64(x), float64(y)) && v.Item2 > 0 {
					setOwnColor(b, false, GS)
					p.KomaTaiIndex = i
				}
			}
			if p.KomaTaiIndex != -1 {
				p.State = playerstate.IS_SELECT_KOMA
			}
			//第一巡
		} else if p.KomaTai[0].Item1.OnKoma(float64(x), float64(y)) {
			setOwnColor(b, false, GS)
			p.KomaTaiIndex = 0
			p.State = playerstate.IS_SELECT_KOMA
		}

	}
}
