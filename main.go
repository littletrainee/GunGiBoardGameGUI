package main

import "github.com/littletrainee/GunGiBoardGameGUI/gamehandler"

func main() {
	var Game gamehandler.Game = gamehandler.Game{}
	Game.Initilization()
	Game.Start()
}
