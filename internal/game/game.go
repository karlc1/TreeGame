package game

import (
	"karlc/treegame/internal/models"
	"karlc/treegame/internal/utils"
	"time"

	"github.com/ByteArena/box2d"
)

var game *Game

type Game struct {
	PhysWorld *box2d.B2World
	AllActors []models.Actor
	AllJoints []*models.Joint
	Player    *models.Box
	TestBox   *models.Box
	Rope      *models.Joint
	Tree      *models.Tree
}

func NewGameObj() *Game {
	gravity := box2d.MakeB2Vec2(0, -40)
	world := box2d.MakeB2World(gravity)
	g := &Game{
		PhysWorld: &world,
	}
	game = g

	return g
}

func (g *Game) InitPlayer() {
	player := models.NewBox(g.PhysWorld, true, -29, 10, 0.5, 0.8)
	player.SetDensity(100)
	player.SetFriction(0.6)
	player.Fixture.SetRestitution(0.15)
	player.Body.SetFixedRotation(false)
	player.Fixture.SetUserData(player)
	g.AllActors = append(g.AllActors, player)
	g.Player = player
}

func (g *Game) InitTree() {
	t := models.NewTree(
		46,
		100,
		45,
		-15,
	)
	g.Tree = t
}

func (g *Game) InitGround() {
	ground := models.NewBox(g.PhysWorld, false, 0, -45, 100, 30)
	ground.SetFriction(0.6)
	g.AllActors = append(g.AllActors, ground)
}

func (g *Game) InitTestBox() {
	testBox := models.NewBox(g.PhysWorld, false, -20, 0, 1, 1)
	g.AllActors = append(g.AllActors, testBox)
	g.TestBox = testBox

	//rope := models.NewJoint(g.PhysWorld, g.Player, testBox)
	//g.AllJoints = append(g.AllJoints, rope)
}

func (g *Game) InitRope() {
	r := models.NewRope(g.PhysWorld, g.Player, g.TestBox)
	g.Rope = r
}

func (g *Game) InitDecor(n int) {
	for i := 0; i < n; i++ {
		size := utils.RandFloat64(0, 0.3)
		posX := utils.RandFloat64(-100, 100)
		posY := utils.RandFloat64(-20, 50)

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
func (g *Game) UpdatePhysics(elapsed time.Duration) {

	//timeStep := 1.0 / 60
	timeStep := elapsed.Seconds()

	velocityIterations := 1
	positionIterations := 1
	g.PhysWorld.Step(
		timeStep,
		velocityIterations,
		positionIterations,
	)
}
