package main

import "github.com/littletrainee/GunGiBoardGameGUI/GameHolder"

func main() {
	var Game GameHolder.Game = GameHolder.Game{}
	Game.Initilization()
	Game.Start()
}
