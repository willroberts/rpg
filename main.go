package main

import (
	"log"

	"engo.io/engo"
	"github.com/willroberts/rpg/rpg"
)

func main() {
	log.Println("[engine] configuring engine")
	opts := engo.RunOptions{
		Title:  "RPG",
		Width:  960,
		Height: 720,
	}
	log.Println("[engine] starting game")
	engo.Run(opts, &rpg.DefaultScene{})
}
