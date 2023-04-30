package gamehandler

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

func (g *Game) Start() {
	// ebiten.SetVsyncEnabled(false)
	ebiten.SetWindowSize(constant.WINDOW_SIZE_WIDTH, constant.WINDOW_SIZE_HEIGHT)
	// ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle(constant.WINDOW_TITLE)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
