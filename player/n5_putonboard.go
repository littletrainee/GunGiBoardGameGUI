package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/board"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
	"github.com/littletrainee/GunGiBoardGameGUI/enum/playerstate"
	"github.com/littletrainee/GunGiBoardGameGUI/gamestate"
	"github.com/littletrainee/GunGiBoardGameGUI/koma"
)

// 駒被選取時
func (p *Player) PutOnBoard(b board.Board, GS gamestate.GameState) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for k, v := range b.Blocks {
			if v.OnBlock(x, y) && !v.HasSuI() && v.Name.Item2 > 6 && len(v.KomaStack) < GS.MaxLayer {
				var (
					tempKoma koma.Koma
					shift    float32
				)
				switch len(v.KomaStack) {
				case 0:
					shift = 0
				case 1:
					shift = constant.SECOND_DAN
				case 2:
					shift = constant.THIRD_DAN
				}
				//複製出來
				tempKoma = p.KomaTai[p.KomaTaiIndex].Item1.Clone()
				// 設定目標格的位置
				tempKoma.SetCurrentCoordinate(v.Coordinate, int(shift))
				//設定目標代號
				tempKoma.SetCurrentPosition(v.Name)
				// 設定文字顯示的位置
				tempKoma.SetGeoMetry(0)
				v.KomaStack = append(v.KomaStack, tempKoma)
				p.KomaTai[p.KomaTaiIndex].Item2--
				if p.KomaTai[p.KomaTaiIndex].Item2 == 0 {
					p.KomaTai[p.KomaTaiIndex].Item1 = koma.Koma{}
				}
				b.Blocks[k] = v
				setOwnColor(b, true, GS)
				p.KomaTaiIndex = -1
				p.State = playerstate.WAIT_TO_CLICK_CLOCK
				return
			}
		}
	}
}
