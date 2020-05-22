package game

import (
	"karlc/treegame/internal/models"
	"math"

	"github.com/ByteArena/box2d"
)

const (
	walking_speed = 20.0
	walking_acc   = 0.6
)

func WalkRight(b *models.Box) {
	if b.State != models.GROUNDED {
		//return
	}
	vel := b.Body.GetLinearVelocity()
	desiredVelocity := math.Max(vel.X+walking_acc, -walking_speed)
	velChange := desiredVelocity - vel.X
	impulse := b.Body.GetMass() * velChange
	b.Body.ApplyLinearImpulse(box2d.MakeB2Vec2(impulse, 0), b.Body.GetWorldCenter(), true)

}

func WalkLeft(b *models.Box) {

	if b.State != models.GROUNDED {
		//return
	}

	vel := b.Body.GetLinearVelocity()
	desiredVelocity := math.Min(vel.X-walking_acc, walking_speed)
	velChange := desiredVelocity - vel.X
	impulse := b.Body.GetMass() * velChange
	b.Body.ApplyLinearImpulse(box2d.MakeB2Vec2(impulse, 0), b.Body.GetWorldCenter(), true)

}

func Jump(b *models.Box) {
	if b.State != models.GROUNDED {
		return
	}

	impulse := b.Body.GetMass() * 10
	linVel := -(b.Body.GetLinearVelocity().X / 4)
	b.Body.ApplyLinearImpulse(box2d.MakeB2Vec2(linVel, impulse), b.Body.GetWorldCenter(), true)
}

// TODO: support grapple different objects
func Grapple(i *InputHandler, b *models.Box) {
	i.NewRope = true
	i.DestroyRope = false
	b.State = models.GRAPPLING
}

func LetGoGrapple(i *InputHandler, b *models.Box) {
	i.DestroyRope = true
	i.NewRope = false
	b.State = models.JUMPING
}
