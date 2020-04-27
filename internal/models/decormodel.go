package models

// DecorBox is a drawable box that is not interactable
// Completely decoupled from physics
type DecorBox struct {
	PosX   float64
	PosY   float64
	Width  float64
	Height float64
	// should eventually determine which objects
	// are in front of which. High Zval is further
	// away from the camera
	// This could alway help with parallax?
	Zval int
}

func (d *DecorBox) GetPosition() (float64, float64) {
	return d.PosX, d.PosY
}

func (d *DecorBox) GetSize() (float64, float64) {
	return d.Width, d.Height
}

func (d *DecorBox) GetAngle() float64 {
	return 0 // can a decorbox have an angle
}

func (d *DecorBox) GetZVal() int {
	return d.Zval
}
