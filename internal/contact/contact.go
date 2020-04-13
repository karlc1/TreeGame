package contact

import (
	"fmt"

	"github.com/ByteArena/box2d"
)

type ContactListener struct{}

func (c ContactListener) BeginContact(contact box2d.B2ContactInterface) {
	fmt.Println("begin contact")
}

func (c ContactListener) EndContact(contact box2d.B2ContactInterface) {
	fmt.Println("end contact")
}

func (c ContactListener) PreSolve(contact box2d.B2ContactInterface, oldManifold box2d.B2Manifold) {
	//fmt.Println("pre solve")
}

func (c ContactListener) PostSolve(contact box2d.B2ContactInterface, impulse *box2d.B2ContactImpulse) {
	//fmt.Println("post solve")
}
