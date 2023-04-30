package main

import "github.com/littletrainee/GunGiBoardGameGUI/gamehandler"

func main() {
	//實例化Game物件並運行
	var Game gamehandler.Game = gamehandler.Game{}
	Game.Initilization()
	Game.Start()
}
