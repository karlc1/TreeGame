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

	n := 800

	points := make([]models.Actor, 0, 0)
	lastPoint := box2d.B2Vec2{}
	for i := 0; i < n; i += 8 {
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

		if i > 1 {

			callBack := func(fixture *box2d.B2Fixture, point box2d.B2Vec2, normal box2d.B2Vec2, fraction float64) float64 {
				// tweak what is a hit here
				game.TrajectoryCollisionPoint = &point

				// terminate loop once something is hit
				i = n
				return 0
			}

			game.PhysWorld.RayCast(callBack, lastPoint, p)
		}

		points = append(points, a)
		lastPoint = p
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
