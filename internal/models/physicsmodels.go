package models

import (
	"karlc/treegame/internal/contact"
	"math"

	"github.com/ByteArena/box2d"
)

type PhysicalWorld struct {
	PhysWorld *box2d.B2World
	Boxes     []*Box //TODO: use dict with pointer keys?
	Player    *Box
}

func NewPhysicalWorld() *PhysicalWorld {
	gravity := box2d.MakeB2Vec2(0, -10)

	world := box2d.MakeB2World(gravity)

	g := &PhysicalWorld{
		PhysWorld: &world,
	}

	cl := contact.ContactListener{}
	world.SetContactListener(cl)

	//ground := g.NewBox(false, 0, 50, 100, 10)
	//_ = ground
	//ground.SetDensity(3)
	//ground.SetFriction(3)

	//test := g.NewBox(true, 200, 250, 10, 20)
	//test.SetDensity(10)
	//test.SetFriction(4)

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

func (b *Box) GetPosition() (float32, float32) {
	return float32(b.Body.M_xf.P.X), float32(b.Body.M_xf.P.Y)
}

func (b *Box) GetSize() (float32, float32) {
	return float32(b.Width), float32(b.Height)
}

func (b *Box) GetAngle() float64 {
	return b.Body.GetAngle()
}

func (b *Box) GetZVal() int {
	// a physics object should always
	// be the same distance from the
	// camera
	return 0
}

func (b *Box) WalkRight() {
	speed := -15.0
	acceleration := 0.5
	vel := b.Body.GetLinearVelocity()
	desiredVelocity := math.Max(vel.X-acceleration, speed)
	velChange := desiredVelocity - vel.X
	impulse := b.Body.GetMass() * velChange
	b.Body.ApplyLinearImpulse(box2d.MakeB2Vec2(impulse, 0), b.Body.GetWorldCenter(), true)
}

func (b *Box) WalkLeft() {
	speed := 15.0
	acceleration := 0.5
	vel := b.Body.GetLinearVelocity()
	desiredVelocity := math.Min(vel.X+acceleration, speed)
	velChange := desiredVelocity - vel.X
	impulse := b.Body.GetMass() * velChange
	b.Body.ApplyLinearImpulse(box2d.MakeB2Vec2(impulse, 0), b.Body.GetWorldCenter(), true)
}

func (b *Box) Jump() {

	impulse := b.Body.GetMass() * 2
	b.Body.ApplyLinearImpulse(box2d.MakeB2Vec2(0, impulse), b.Body.GetWorldCenter(), true)
}

func (w *PhysicalWorld) NewBox(dynamic bool, posX, posY, width, height float64) *Box {
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
