package scenes

import (
	"image/color"
	"log"

	"github.com/willroberts/rpg/rpg"

	"engo.io/ecs"
	"engo.io/engo/common"
)

type BasicScene struct{}

// Preload() validates and loads assets.
func (scene *BasicScene) Preload() {
	log.Println("preloading")
	rpg.PreloadMapAssets("maps/basic.tmx")
	_, err := rpg.PreloadSprites()
	if err != nil {
		log.Fatalln("error:", err.Error())
	}
}

// Setup() initializes all systems necessary for the game to function.
func (scene *BasicScene) Setup(w *ecs.World) {
	log.Println("creating scene")
	common.SetBackground(color.Black)
	w.AddSystem(&common.RenderSystem{})

	log.Println("loading map")
	_, tiles, err := rpg.LoadMap("maps/basic.tmx")
	if err != nil {
		log.Fatalln("error:", err.Error())
	}

	log.Println("configuring systems")
	for _, sys := range w.Systems() {
		switch s := sys.(type) {
		case *common.RenderSystem:
			log.Println("configuring render system")
			for _, t := range tiles {
				s.Add(&t.BasicEntity, &t.RenderComponent, &t.SpaceComponent)
			}
		}
	}
}

// Type() returns the name of the scene to satisfy engo's Scene interface.
func (scene *BasicScene) Type() string {
	return "BasicScene"
}
