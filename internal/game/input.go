package game

import (
	"karlc/treegame/internal/physics"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func HandleInput() {

	if rl.IsKeyDown(rl.KeyUp) {
		physics.Jump(game.Player)
	}

	if rl.IsKeyDown(rl.KeyRight) {
		physics.WalkRight(game.Player)
	}

	if rl.IsKeyDown(rl.KeyLeft) {
		physics.WalkLeft(game.Player)
	}
}
