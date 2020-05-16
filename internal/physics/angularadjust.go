package physics

import (
	"fmt"
	"karlc/treegame/internal/config"
	"karlc/treegame/internal/game"
	"karlc/treegame/internal/models"
	"math"

	"github.com/ByteArena/box2d"
)

var lastVel float64

// AdjustPlayerAngularVelocity is used to always have the
// player land upright
// https://www.youtube.com/watch?v=BZwizmCI_g0
func AdjustAngularVelocity(obj *models.Box, game *game.Game, cfg *config.Config) {

	n := 400

	points := make([]models.Actor, 0, 0)
	prevPoint := box2d.B2Vec2{}
	collisionFound := false

	t := 0
	for i := 0; i < n; i += 2 {
		startPos := obj.Body.GetPosition()
		startVel := obj.Body.GetLinearVelocity()
		fps := cfg.TargetFPS
		gravity := game.PhysWorld.M_gravity
		currPoint := getTrajectoryPoint(startPos, startVel, gravity, fps, float64(i))

		// skip first iteration to avoid obj colliding with itself
		if i > 1 {

			callBack := func(fixture *box2d.B2Fixture, point box2d.B2Vec2, normal box2d.B2Vec2, fraction float64) float64 {

				// store collision point for drawing
				game.TrajectoryCollisionPoint = &point

				// store amount of timesteps until collision
				t = i

				// terminate loop at first collision
				i = n

				collisionFound = true

				// terminate ray cast
				return 0
			}

			game.PhysWorld.RayCast(callBack, prevPoint, currPoint)
		}

		a := &models.DecorBox{
			PosX:   currPoint.X,
			PosY:   currPoint.Y,
			Width:  0.25,
			Height: 0.25,
		}
		points = append(points, a)
		prevPoint = currPoint
	}

	game.TrajectoryPoints = points

	// if no collision was found in the search space n,
	// don't adjust yet
	if !collisionFound {
		return
	}

	// if timeSteps until collision is either min or max
	// nothing useful was found
	if t == 0 || t == n {
		return
	}
	angularVel := obj.Body.GetAngularVelocity()

	// if angular velocit is already 0, there is no need to adjust
	if angularVel == 0 {
		return
	}

	currentAngle := obj.Body.GetAngle()

	var finalAngle float64

	time := float64(t) * 1 / cfg.TargetFPS

	fullRotation := math.Pi * 2
	if angularVel < 0 {
		fullRotation = -fullRotation
	}

	finalAngle = math.Mod(currentAngle+angularVel*time, -fullRotation)

	if angularVel > 0 {
		if finalAngle > 0 {
			adjust(obj, time, false)
		} else {
			adjust(obj, time, true)
		}
	} else {
		if finalAngle > 0 {
			adjust(obj, time, false)
		} else {
			adjust(obj, time, true)
		}

	}

	if tmp%4 == 0 {
		fmt.Printf("Final angle: %v\n Current vel: %v\n\n", finalAngle, angularVel)
	}

	tmp++
}

func adjust(b *models.Box, t float64, increase bool) {

	// adjust more if there is less time until collision
	//timeComponent := 1 / t

	currV := b.Body.GetAngularVelocity()
	var newV float64
	if increase {
		//newV = currV + currV/500
		newV = currV + 0.1
	} else {
		//newV = currV - currV/500
		newV = currV - 0.1
	}

	b.Body.SetAngularVelocity(newV)
}

var tmp int64 = 0

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
