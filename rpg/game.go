package rpg

import (
	"log"

	"engo.io/engo"
)

var gameScene *GameScene

type GameScene struct {
	Fonts *FontSet
	HUD   *HUD
	Log   *ActivityLog
}

func (scene *GameScene) Preload() {
	gameScene = scene
	scene.Fonts, err = PreloadFonts()
	if err != nil {
		log.Fatalln("failed to preload fonts:", err.Error())
	}
}

func (scene *GameScene) Setup(u engo.Updater) {
	gameScene.HUD, err = newHUD()
	if err != nil {
		log.Fatalln("error:", err.Error())
	}
	scene.Log = newActivityLog()
	scene.Log.Update("Welcome to the game.")
	scene.Log.Update("There are three skeletons near you.")
	scene.Log.Update("Try moving into them to attack.")
}
