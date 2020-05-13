package game

import (
	"github.com/faiface/pixel/pixelgl"
)

type InputHandler struct {
	win         *pixelgl.Window
	DestroyRope bool
	NewRope     bool
}

func NewInputHandler(win *pixelgl.Window) *InputHandler {
	return &InputHandler{
		win: win,
	}
}

func (i *InputHandler) HandleInput() {

	if i.win.Pressed(pixelgl.KeyUp) {
		Jump(game.Player.Box)
	}

	if i.win.Pressed(pixelgl.KeyRight) {
		WalkRight(game.Player.Box)
	}

	if i.win.Pressed(pixelgl.KeyLeft) {
		WalkLeft(game.Player.Box)
	}

	if i.win.Pressed(pixelgl.KeyEscape) {
		i.win.SetClosed(true)
	}

	if i.win.Pressed(pixelgl.KeySpace) {
		i.DestroyRope = true
		i.NewRope = false
	}

	if i.win.Pressed(pixelgl.KeyX) {
		i.NewRope = true
		i.DestroyRope = false
	}
}
