package game

import (
	"karlc/treegame/internal/models"
	"karlc/treegame/internal/utils"

	"github.com/ByteArena/box2d"
)

var game *Game

type Game struct {
	PhysWorld *box2d.B2World
	AllActors []models.Actor
	AllJoints []*models.Joint
	Player    *models.Box
}

func NewGameObj() *Game {
	gravity := box2d.MakeB2Vec2(0, -10)
	world := box2d.MakeB2World(gravity)
	g := &Game{
		PhysWorld: &world,
	}
	game = g
	return g
}

func (g *Game) InitPlayer() {
	player := models.NewBox(g.PhysWorld, true, -29, 10, 0.5, 0.8)
	player.SetDensity(1)
	player.SetFriction(0.6)
	player.Fixture.SetRestitution(0.15)
	player.Body.SetFixedRotation(false)
	player.Fixture.SetUserData(player)
	g.AllActors = append(g.AllActors, player)
	g.Player = player
}

func (g *Game) InitGround() {
	ground := models.NewBox(g.PhysWorld, false, 0, -45, 100, 30)
	ground.SetFriction(0.6)
	g.AllActors = append(g.AllActors, ground)
}

func (g *Game) InitTestBox() {
	testBox := models.NewBox(g.PhysWorld, false, -20, 10, 1, 1)
	g.AllActors = append(g.AllActors, testBox)

	rope := models.NewJoint(g.PhysWorld, g.Player, testBox)
	g.AllJoints = append(g.AllJoints, rope)
}

func (g *Game) InitDecor(n int) {
	for i := 0; i < n; i++ {
		size := utils.RandFloat32(0, 0.3)
		posX := utils.RandFloat32(-100, 100)
		posY := utils.RandFloat32(-20, 50)

		// for parallax test
		z := int(size * 10)

		g.AllActors = append(g.AllActors, &models.DecorBox{
			Height: size,
			Width:  size,
			PosX:   posX,
			PosY:   posY,
			Zval:   z,
		})
	}
}

// UpdatePhysics steps the physics simulation forward
func (g *Game) UpdatePhysics() {
	timeStep := 1.0 / 30
	velocityIterations := 2
	positionIterations := 2
	g.PhysWorld.Step(
		timeStep,
		velocityIterations,
		positionIterations,
	)
}
