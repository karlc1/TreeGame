package game

import (
	"karlc/treegame/internal/models"

	"github.com/ByteArena/box2d"
)

type ContactListener struct {
	player *models.Box
}

func NewContactListener(a *models.Box) ContactListener {
	return ContactListener{
		player: a,
	}
}

func (c ContactListener) BeginContact(contact box2d.B2ContactInterface) {
	a := contact.GetFixtureA().GetUserData()
	b := contact.GetFixtureB().GetUserData()

	if a == c.player || b == c.player {
		c.player.State = models.GROUNDED
	}

	// needs to be called outsite of world update somehow
	//game.Player.Box.Body.SetTransform(box2d.MakeB2Vec2(0, 0), 0)
	c.player.Body.SetFixedRotation(true)
}

func (c ContactListener) EndContact(contact box2d.B2ContactInterface) {
	a := contact.GetFixtureA().GetUserData()
	b := contact.GetFixtureB().GetUserData()

	if a == c.player || b == c.player {
		c.player.State = models.JUMPING
	}

	c.player.Body.SetFixedRotation(false)
}

func (c ContactListener) PreSolve(contact box2d.B2ContactInterface, oldManifold box2d.B2Manifold) {
	//fmt.Println("pre solve")
}

func (c ContactListener) PostSolve(contact box2d.B2ContactInterface, impulse *box2d.B2ContactImpulse) {
	//fmt.Println("post solve")
}
