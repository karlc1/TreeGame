package game

import (
	"karlc/treegame/internal/physics"

	"github.com/faiface/pixel/pixelgl"
)

type InputHandler struct {
	win *pixelgl.Window
}

func NewInputHandler(win *pixelgl.Window) *InputHandler {
	return &InputHandler{
		win: win,
	}
}

func (i *InputHandler) HandleInput() {

	if i.win.Pressed(pixelgl.KeyUp) {
		physics.Jump(game.Player)
	}

	if i.win.Pressed(pixelgl.KeyRight) {
		physics.WalkRight(game.Player)
	}

	if i.win.Pressed(pixelgl.KeyLeft) {
		physics.WalkLeft(game.Player)
	}

	if i.win.Pressed(pixelgl.KeyEscape) {
		i.win.SetClosed(true)
	}
}
