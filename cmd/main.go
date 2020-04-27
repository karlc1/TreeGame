package main

import (
	"karlc/treegame/internal/game"
	"karlc/treegame/internal/physics"
	"karlc/treegame/internal/render"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	SCREEN_WITH   = 1200
	SCREEN_HEIGHT = 600
	TARGET_FPS    = 60
)

func main() {
	pixelgl.Run(run)

	//render.InitWindow(SCREEN_WITH, SCREEN_HEIGHT, TARGET_FPS)

	//gameObj := game.NewGameObj()
	//gameObj.InitPlayer()
	//gameObj.InitGround()
	//gameObj.InitDecor(500)
	//gameObj.InitTestBox()

	//camera := render.NewCamera(
	//SCREEN_WITH,
	//SCREEN_HEIGHT,
	//20,
	//)
	//camera.OffsetY = -20

	//contactListener := physics.NewContactListener(gameObj.Player)
	//gameObj.PhysWorld.SetContactListener(contactListener)

	//camera.AttachTo(gameObj.Player)

	//for !rl.WindowShouldClose() {
	//gameObj.UpdatePhysics()
	//game.HandleInput()
	//camera.DrawGame(gameObj)
	//}

	//rl.CloseWindow()
}

func run() {
	win := setupWindow()
	inputHandler := game.NewInputHandler(win)
	game := setupGame()
	camera := setupCamera(game, win)

	fps := time.Tick(time.Second / 60)

	for !win.Closed() {
		inputHandler.HandleInput()
		game.UpdatePhysics()

		win.Clear(colornames.Black)
		//camera.TestDraw()
		camera.DrawGame(game)
		win.Update()

		<-fps
	}

}

func setupWindow() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, SCREEN_WITH, SCREEN_HEIGHT),
		//VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	return win
}

func setupGame() *game.Game {
	gameObj := game.NewGameObj()
	gameObj.InitPlayer()
	gameObj.InitGround()
	//gameObj.InitDecor(500)
	//gameObj.InitTestBox()
	contactListener := physics.NewContactListener(gameObj.Player)
	gameObj.PhysWorld.SetContactListener(contactListener)
	return gameObj
}

func setupCamera(game *game.Game, win *pixelgl.Window) *render.Camera {
	camera := render.NewCamera(
		SCREEN_WITH,
		SCREEN_HEIGHT,
		20,
		win,
	)
	camera.OffsetY = -20
	camera.AttachTo(game.Player)
	return camera
}
