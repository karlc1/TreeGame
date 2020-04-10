package physics

import (
	"github.com/ByteArena/box2d"
)

type World struct {
	PhysWorld *box2d.B2World
	Boxes     []*Box //TODO: use dict with pointer keys?
	Player    *Box
}

func NewWorld() *World {
	gravity := box2d.MakeB2Vec2(0, -10)
	world := box2d.MakeB2World(gravity)

	g := &World{
		PhysWorld: &world,
	}

	ground := g.NewBox(false, 320, 50, 100, 10)
	_ = ground
	//ground.SetDensity(3)
	//ground.SetFriction(3)

	player := g.NewBox(true, 200, 200, 20, 20)
	player.SetDensity(1)
	player.SetFriction(0.3)

	test := g.NewBox(true, 200, 250, 10, 20)
	test.SetDensity(1000)
	test.SetFriction(200)

	g.Player = player

	return g
}

type Box struct {
	Body    *box2d.B2Body
	Shape   *box2d.B2PolygonShape
	Fixture *box2d.B2Fixture
	Width   float64
	Height  float64
}

func (b *Box) SetDensity(d float64) {
	b.Fixture.SetDensity(d)
	b.Body.ResetMassData()
}

func (b *Box) SetFriction(f float64) {
	b.Fixture.SetFriction(f)
}

func (b *Box) GetPosition() box2d.B2Vec2 {
	return b.Body.M_xf.P
}

func (w *World) NewBox(dynamic bool, posX, posY, width, height float64) *Box {
	boxDef := box2d.MakeB2BodyDef()

	// static by default, only check if static
	// note that friction and density should be set
	// after initialization if the body is dynamic
	if dynamic {
		boxDef.Type = box2d.B2BodyType.B2_dynamicBody
	}

	boxDef.Position = box2d.MakeB2Vec2(posX, posY)
	boxShape := box2d.MakeB2PolygonShape()
	boxShape.SetAsBox(width, height)
	boxBody := w.PhysWorld.CreateBody(&boxDef)
	boxFixture := boxBody.CreateFixture(&boxShape, 0.0)

	box := &Box{
		Body:    boxBody,
		Shape:   &boxShape,
		Fixture: boxFixture,
		Width:   width * 2,
		Height:  height * 2,
	}

	w.Boxes = append(w.Boxes, box)
	return box
}
