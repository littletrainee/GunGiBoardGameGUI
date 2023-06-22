package gamehandler

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/littletrainee/GunGiBoardGameGUI/constant"
)

// Start 在初始化完成後的執行GUI的顯示
func (g *Game) Start() {
	// ebiten.SetVsyncEnabled(false)
	ebiten.SetWindowSize(constant.WINDOW_SIZE_WIDTH, constant.WINDOW_SIZE_HEIGHT)
	// ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle(constant.WINDOW_TITLE)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
