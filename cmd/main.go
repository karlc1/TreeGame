package main

import (
	"karlc/treegame/internal/game"
	"karlc/treegame/internal/physics"
	"karlc/treegame/internal/render"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_WITH   = 1200
	SCREEN_HEIGHT = 600
	TARGET_FPS    = 60
)

func main() {
	render.InitWindow(SCREEN_WITH, SCREEN_HEIGHT, TARGET_FPS)

	gameObj := game.NewGameObj()
	gameObj.InitPlayer()
	gameObj.InitGround()
	gameObj.InitDecor(500)
	gameObj.InitTestBox()

	camera := render.NewCamera(
		SCREEN_WITH,
		SCREEN_HEIGHT,
		20,
	)
	camera.OffsetY = -20

	contactListener := physics.NewContactListener(gameObj.Player)
	gameObj.PhysWorld.SetContactListener(contactListener)

	camera.AttachTo(gameObj.Player)

	for !rl.WindowShouldClose() {
		gameObj.UpdatePhysics()
		game.HandleInput()
		camera.DrawGame(gameObj)
	}

	rl.CloseWindow()
}
