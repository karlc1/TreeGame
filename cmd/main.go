package main

import (
	"fmt"
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
	TARGET_FPS    = 144
)

func main() {
	pixelgl.Run(run)
}

func run() {
	win := setupWindow()
	inputHandler := game.NewInputHandler(win)
	game := setupGame()
	camera := setupCamera(game, win)
	//treeCamera := setupTreeCamera(game, win)

	fpsTick := time.Tick(time.Second / TARGET_FPS)
	secondTick := time.Tick(time.Second)
	frames := 0
	lastFrame := time.Now()
	var elapsedTime time.Duration

	for !win.Closed() {

		elapsedTime = time.Since(lastFrame)
		lastFrame = time.Now()

		select {
		case <-secondTick:
			fmt.Printf("FPS: %v \n", frames)
			frames = 0
			fmt.Printf("Timestep: %v \n", elapsedTime.Seconds())
		default:
			frames++
		}

		inputHandler.HandleInput()
		game.UpdatePhysics(elapsedTime)

		win.Clear(colornames.Black)
		camera.TestDraw()
		camera.DrawGame(game)
		//treeCamera.DrawGame()
		win.Update()

		playerPosX, _ := game.Player.GetPosition()
		if playerPosX >= game.Tree.PosX-3 && playerPosX <= game.Tree.PosX+3 {
			camera.TreeView = true
		}

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

func setupWindow() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, SCREEN_WITH, SCREEN_HEIGHT),
		VSync:  false,
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
	gameObj.InitDecor(100)
	gameObj.InitTestBox()
	//gameObj.InitRope()
	gameObj.InitTree()
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

func setupTreeCamera(game *game.Game, win *pixelgl.Window) *render.TreeCamera {
	camera := render.NewTreeCamera(
		SCREEN_WITH,
		SCREEN_HEIGHT,
		20,
		win,
		game.Player,
		game.Tree,
	)
	camera.OffsetY = -20
	return camera
}
