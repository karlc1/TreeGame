package game

import (
	"karlc/treegame/internal/config"
	"karlc/treegame/internal/models"
	"karlc/treegame/internal/utils"

	"github.com/ByteArena/box2d"
)

var game *Game

type Game struct {
	Config           *config.Config
	PhysWorld        *box2d.B2World
	AllActors        []models.Actor
	AllJoints        []*models.Joint
	TrajectoryPoints []models.Actor
	Player           *models.Player
	TestBox          *models.Box
	Rope             *models.Joint
	Ground           *models.Box
	GravityX         float64
	GravityY         float64
}

func NewGameObj(config *config.Config) *Game {
	gravityX := 0.0
	gravityY := -20.0
	gravity := box2d.MakeB2Vec2(gravityX, gravityY)
	world := box2d.MakeB2World(gravity)

	g := &Game{
		Config:    config,
		PhysWorld: &world,
		GravityX:  gravityX,
		GravityY:  gravityY,
	}
	game = g

	return g
}

func (g *Game) ExitGame() {
	g.PhysWorld.Destroy()
}

func (g *Game) InitPlayer() {
	playerBox := models.NewBox(g.PhysWorld, true, -29, 10, 0.5, 0.8)
	playerBox.SetDensity(100)
	playerBox.SetFriction(0.6)
	playerBox.Fixture.SetRestitution(0.15)
	playerBox.Body.SetFixedRotation(false)
	playerBox.Fixture.SetUserData(playerBox)

	//playerBox.Body.SetAngularDamping(0.5)

	g.AllActors = append(g.AllActors, playerBox)
	g.Player = &models.Player{
		Box: playerBox,
	}
}

func (g *Game) InitGround() {
	ground := models.NewBox(g.PhysWorld, false, 0, -45, 100, 30)
	ground.SetFriction(0.6)
	g.AllActors = append(g.AllActors, ground)
	g.Ground = ground
}

func (g *Game) InitTestBox() {
	testBox := models.NewBox(g.PhysWorld, false, -20, 0, 1, 1)

	g.AllActors = append(g.AllActors, testBox)
	g.TestBox = testBox

	//rope := models.NewJoint(g.PhysWorld, g.Player, testBox)
	//g.AllJoints = append(g.AllJoints, rope)
}

func (g *Game) InitRope() {
	r := models.NewRope(g.PhysWorld, g.Player.Box, g.TestBox)
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

func (g *Game) InitContactListener() {
	contactListener := NewContactListener(g.Player.Box)
	g.PhysWorld.SetContactListener(contactListener)
}

// UpdatePhysics steps the physics simulation forward
func (g *Game) UpdatePhysics() {

	timeStep := 1.0 / float64(g.Config.TargetFPS)

	velocityIterations := 1
	positionIterations := 1
	g.PhysWorld.Step(
		timeStep,
		velocityIterations,
		positionIterations,
	)
}
