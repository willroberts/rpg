package main

import (
	"fmt"
	"os"

	"engo.io/engo"
	"github.com/willroberts/rpg/internal/scene"
	"go.uber.org/zap"
)

const (
	width  int = 720
	height int = 720
)

func main() {
	o := engo.RunOptions{
		Title:  "RPG",
		Width:  width,
		Height: height,
	}

	s := &scene.GameScene{}
	if err := s.InitLogger(); err != nil {
		fmt.Println("failed to initialize logger:", err)
		os.Exit(1)
	}
	s.Logger.Info("initialized logger")

	s.Logger.Info("running scene",
		zap.String("name", s.Type()),
		zap.Int("width", width),
		zap.Int("height", height),
	)
	engo.Run(o, s)
}
