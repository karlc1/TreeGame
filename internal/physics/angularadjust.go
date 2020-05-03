package physics

import (
	"fmt"
	"karlc/treegame/internal/models"
	"math"
	"time"
)

var lastVel float64

// AdjustPlayerAngularVelocity is used to always have the
// player land upright
// https://www.youtube.com/watch?v=BZwizmCI_g0
func AdjustAngularVelocity(player *models.Player, ground *models.Box, gravY float64) {

	pVelY := player.Box.Body.GetLinearVelocity().Y
	defer func() {
		lastVel = pVelY
	}()

	if !(lastVel > 0 && pVelY < 0) {
		return
	}

	//only adjust if player is falling downwards
	//if pVelY > 0 {
	//fmt.Println("--------")
	//return
	//}

	t := getTimeUntilGrounded(player, ground, gravY)

	go func() {
		timer := time.NewTimer(time.Millisecond * time.Duration(t*1000))
		<-timer.C
		fmt.Println("TIMER DONE")
	}()

	fullCircle := math.Pi * 2

	// do the rotation modulo a full rotation in radians
	currentAngle := (math.Mod(player.Box.GetAngle(), fullCircle))

	angularVel := (player.Box.Body.GetAngularVelocity())

	// predicted landing angle
	pred := math.Mod(currentAngle+angularVel*t, fullCircle)

	fmt.Println(t)

	_ = pred
	_ = t
	_ = angularVel

}

func getTimeUntilGrounded(player *models.Player, ground *models.Box, gravY float64) float64 {
	// player coordinates
	_, ph := player.Box.GetSize()
	_, py := player.Box.GetPosition()
	// bottom of player box
	feet := py - ph/2

	// ground coordinates
	_, gh := ground.GetSize()
	_, gy := ground.GetPosition()
	// top of ground box
	groundTop := gy + gh/2

	jumpHeight := feet - groundTop

	// time until feet hits ground
	t := math.Sqrt((2 * jumpHeight) / math.Abs(gravY))

	//fmt.Printf("playerY: %v \n", py)
	//fmt.Printf("groundY: %v \n", gy)
	//fmt.Printf("feet: %v \n", feet)
	//fmt.Printf("ground: %v \n", groundTop)
	//fmt.Printf("jumpHeight: %v \n", jumpHeight)
	//fmt.Printf("Time until ground: %v \n \n", t)

	return t
}

func getTimeUntilGrounded2(player *models.Player, ground *models.Box, gravY, velY float64) float64 {
	// player coordinates
	_, ph := player.Box.GetSize()
	_, py := player.Box.GetPosition()
	// bottom of player box
	feet := py - ph/2

	// ground coordinates
	_, gh := ground.GetSize()
	_, gy := ground.GetPosition()
	// top of ground box
	groundTop := gy + gh/2

	jumpHeight := feet - groundTop

	// time until feet hits ground
	t := math.Sqrt((2 * (jumpHeight + velY)) / math.Abs(gravY))

	//fmt.Printf("playerY: %v \n", py)
	//fmt.Printf("groundY: %v \n", gy)
	//fmt.Printf("feet: %v \n", feet)
	//fmt.Printf("ground: %v \n", groundTop)
	//fmt.Printf("jumpHeight: %v \n", jumpHeight)
	//fmt.Printf("Time until ground: %v \n \n", t)

	return t
}
