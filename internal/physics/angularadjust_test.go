package physics

import (
	"karlc/treegame/internal/game"
	"testing"
	"time"

	"github.com/ByteArena/box2d"
)

// timerContactListener is a contact listener
// meant for testing timing predictions
type timerContactListener struct {
	contact chan bool
	t       *testing.T
}

func (c timerContactListener) BeginContact(contact box2d.B2ContactInterface) {
	c.contact <- true
}
func (c timerContactListener) EndContact(contact box2d.B2ContactInterface) {}
func (c timerContactListener) PreSolve(contact box2d.B2ContactInterface, oldManifold box2d.B2Manifold) {
}
func (c timerContactListener) PostSolve(contact box2d.B2ContactInterface, impulse *box2d.B2ContactImpulse) {
}

func setupGame() {

}

func TestFallingFromCloseDistanceCanBePredicted(t *testing.T) {
	gameObj := game.NewGameObj()
	gameObj.InitGround()
	gameObj.InitPlayer()
	defer gameObj.ExitGame()

	// move player to 20 units above ground
	//gameObj.Player.Box.Body.SetTransform(
	//box2d.MakeB2Vec2(0, 20),
	//gameObj.Player.Box.Body.GetAngle(),
	//)

	//fpsTick := time.Tick(time.Second / 60)
	lastFrame := time.Now()

	contact := make(chan bool)
	contactListener := timerContactListener{contact: contact, t: t}
	gameObj.PhysWorld.SetContactListener(contactListener)

	startTime := time.Now()
	var contactTime time.Time

	timer := time.NewTimer(time.Second * 5)

	for alive := true; alive; {
		//gameObj.UpdatePhysics(time.Since(lastFrame))
		lastFrame = time.Now()
		_ = lastFrame

		select {
		//case <-contact:
		//contactTime = time.Now()
		//alive = false
		case <-timer.C:
			t.Fatal("time end")
			alive = false
		default:
			t.Error("d")
		}
	}

	actualFallTIme := startTime.Sub(contactTime)
	t.Fatal(actualFallTIme)
}

func TestTmp(t *testing.T) {
	gameObj := game.NewGameObj()
	gameObj.InitGround()
	gameObj.InitPlayer()
	defer gameObj.ExitGame()

	timer := time.NewTimer(time.Second * 2)

	for alive := true; alive; {

		gameObj.UpdatePhysics()
		_, Y := gameObj.Player.Box.GetPosition()

		t.Errorf("y: %v\n", Y)

		select {
		case <-timer.C:
			t.Fatal("timeout")
			alive = false
		default:
			continue
		}
	}
}
