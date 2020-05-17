package models

import (
	"fmt"

	"github.com/ByteArena/box2d"
)

type State int

const (
	GROUNDED State = iota + 1
	JUMPING
	GRAPPLING
)

type Box struct {
	Body    *box2d.B2Body
	Shape   *box2d.B2PolygonShape
	Fixture *box2d.B2Fixture
	Width   float64
	Height  float64
	State   State
}

func (b *Box) SetDensity(d float64) {
	b.Fixture.SetDensity(d)
	b.Body.ResetMassData()
}

func (b *Box) SetFriction(f float64) {
	b.Fixture.SetFriction(f)
}

func (b *Box) GetPosition() (float64, float64) {
	return b.Body.M_xf.P.X, b.Body.M_xf.P.Y
}

func (b *Box) GetLinearVelocity() (float64, float64) {
	return b.Body.GetLinearVelocity().X, b.Body.GetLinearVelocity().Y
}

func (b *Box) GetSize() (float64, float64) {
	return b.Width, b.Height
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

func NewBox(world *box2d.B2World, dynamic bool, posX, posY, width, height float64) *Box {
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
	boxBody := world.CreateBody(&boxDef)
	boxFixture := boxBody.CreateFixture(&boxShape, 0.0)

	box := &Box{
		Body:    boxBody,
		Shape:   &boxShape,
		Fixture: boxFixture,
		Width:   width * 2,
		Height:  height * 2,
		State:   JUMPING,
	}

	return box
}

type Joint struct {
	B2Joint  box2d.B2JointInterface
	AnchorAX float64
	AnchorAY float64
	AnchorBX float64
	AnchorBY float64
}

func NewJoint(world *box2d.B2World, A, B *Box) *Joint {
	jointDef := box2d.MakeB2RopeJointDef()
	jointDef.BodyA = A.Body
	jointDef.BodyB = B.Body
	jointDef.CollideConnected = true

	jointDef.MaxLength = 5

	j := world.CreateJoint(&jointDef)

	joint := &Joint{
		B2Joint: j,
	}

	return joint
}

func NewRope(world *box2d.B2World, player, box *Box) *Joint {
	jointDef := box2d.MakeB2RopeJointDef()
	jointDef.SetBodyA(player.Body)
	jointDef.SetBodyB(box.Body)

	jointDef.MaxLength = 8
	jointDef.CollideConnected = true

	jointDef.LocalAnchorA = box2d.MakeB2Vec2(0, 1)
	jointDef.LocalAnchorB = box2d.MakeB2Vec2(0, -1)

	fmt.Printf("ANCHOR PLAYER X: %v \n", jointDef.LocalAnchorA.X)
	fmt.Printf("ANCHOR PLAYER Y: %v \n", jointDef.LocalAnchorA.Y)

	fmt.Printf("ANCHOR BOX X: %v \n", jointDef.LocalAnchorB.X)
	fmt.Printf("ANCHOR BOX Y: %v \n", jointDef.LocalAnchorB.Y)

	j := world.CreateJoint(&jointDef)

	return &Joint{
		B2Joint:  j,
		AnchorAX: jointDef.LocalAnchorA.X,
		AnchorAY: jointDef.LocalAnchorA.Y,
		AnchorBX: jointDef.LocalAnchorB.X,
		AnchorBY: jointDef.LocalAnchorB.Y,
	}
}
