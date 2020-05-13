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
}

func (c ContactListener) EndContact(contact box2d.B2ContactInterface) {
	a := contact.GetFixtureA().GetUserData()
	b := contact.GetFixtureB().GetUserData()

	if a == c.player || b == c.player {
		c.player.State = models.JUMPING
	}
}

func (c ContactListener) PreSolve(contact box2d.B2ContactInterface, oldManifold box2d.B2Manifold) {
	//fmt.Println("pre solve")
}

func (c ContactListener) PostSolve(contact box2d.B2ContactInterface, impulse *box2d.B2ContactImpulse) {
	//fmt.Println("post solve")
}
