package main

import (
	"fmt"
	"karlc/treegame/internal/config"
	"karlc/treegame/internal/game"
	"karlc/treegame/internal/physics"
	"karlc/treegame/internal/render"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	config := config.Default()
	win := setupWindow(config)
	fmt.Println(win.Bounds().Size())
	inputHandler := game.NewInputHandler(win)
	game := setupGame(config)
	defer game.ExitGame()
	camera := setupCamera(config, game, win)

	fpsTick := time.Tick(time.Second / time.Duration(config.TargetFPS))
	secondTick := time.Tick(time.Second)
	frames := 0

	for !win.Closed() {

		select {
		case <-secondTick:
			//fmt.Printf("FPS: %v \n", frames)
			frames = 0
			//fmt.Printf("Timestep: %v \n", elapsedTime.Seconds())
		default:
			frames++
		}

		inputHandler.HandleInput()
		game.UpdatePhysics()

		win.Clear(colornames.Darkslategray)
		camera.TestDraw()
		camera.DrawGame(game)
		win.Update()

		physics.AdjustAngularVelocity(
			game.Player.Box,
			game,
			config,
		)

		<-fpsTick

		if inputHandler.DestroyRope && !game.PhysWorld.IsLocked() && game.Rope != nil {
			game.PhysWorld.DestroyJoint(game.Rope.B2Joint)
			game.Rope = nil
		}

		if inputHandler.NewRope && game.Rope == nil {
			game.InitRope()
		}
	}
}

func setupWindow(cfg *config.Config) *pixelgl.Window {
	winCfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, cfg.ScreenWidth, cfg.ScreenHeight),
		VSync:  true,
		//Resizable: true,
	}

	win, err := pixelgl.NewWindow(winCfg)
	win.SetCursorVisible(false)
	if err != nil {
		panic(err)
	}
	return win
}

func setupGame(config *config.Config) *game.Game {
	gameObj := game.NewGameObj(config)
	gameObj.InitPlayer()
	gameObj.InitGround()
	gameObj.InitDecor(100)
	gameObj.InitTestBox()
	gameObj.InitRope()
	gameObj.InitContactListener()

	return gameObj
}

func setupCamera(cfg *config.Config, game *game.Game, win *pixelgl.Window) *render.Camera {
	camera := render.NewCamera(
		cfg,
		15,
		win,
	)
	camera.OffsetY = -13
	camera.AttachTo(game.Player.Box)
	return camera
}
