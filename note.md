func (g *Game) Draw(screen *ebiten.Image) {
vector.DrawFilledCircle(screen, 40, 40, 40, color.RGBA{255, 0, 0, 0}, true)
}

screen size 1024,768
boardsize 550,550 border 1
block size 60 60 not border
own komaDai size 182 position 788 , 477 border 1
opponent komadat size 182 position 54,109 border 1
