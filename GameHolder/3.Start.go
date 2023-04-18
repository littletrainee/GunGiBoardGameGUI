package GameHolder

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/Const"
)

func (g *Game) Start() {

	ebiten.SetWindowSize(Const.WINDOW_SIZE_WIDTH, Const.WINDOW_SIZE_HEIGHT)
	// ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle(Const.WINDOW_TITLE)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
