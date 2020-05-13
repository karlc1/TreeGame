package physics

import (
	"karlc/treegame/internal/config"
	"karlc/treegame/internal/game"
	"karlc/treegame/internal/models"

	"github.com/ByteArena/box2d"
)

var lastVel float64

// AdjustPlayerAngularVelocity is used to always have the
// player land upright
// https://www.youtube.com/watch?v=BZwizmCI_g0
func AdjustAngularVelocity(obj *models.Box, game *game.Game, cfg *config.Config) {

	n := 100

	points := make([]models.Actor, n, n)
	for i := 0; i < n; i++ {
		startPos := obj.Body.GetPosition()
		startVel := obj.Body.GetLinearVelocity()
		fps := cfg.TargetFPS
		gravity := game.PhysWorld.M_gravity
		p := getTrajectoryPoint(startPos, startVel, gravity, fps, float64(i))

		a := &models.DecorBox{
			PosX:   p.X,
			PosY:   p.Y,
			Width:  0.1,
			Height: 0.1,
		}
		points[i] = a
	}

	game.TrajectoryPoints = points

}

func getTrajectoryPoint(startPos, stepVelocity, gravity box2d.B2Vec2, fps, n float64) box2d.B2Vec2 {
	t := float64(1) / fps
	stepVelocity.OperatorScalarMulInplace(t)
	gravity.OperatorScalarMulInplace(t * t)

	stepVelocity.OperatorScalarMulInplace(n)
	gravity.OperatorScalarMulInplace(0.5 * (n*n + n))

	startPos.X += stepVelocity.X + gravity.X
	startPos.Y += stepVelocity.Y + gravity.Y

	return startPos
}
