package main

import (
	"engo.io/engo"
)

type GameWorld struct{}

func (game *GameWorld) Type() string {
	return "GameWorld"
}

func main() {
	opts := engo.RunOptions{
		Title:  "RPG",
		Width:  960,
		Height: 720,
	}
	engo.Run(opts, &GameWorld{})
}
