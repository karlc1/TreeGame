package render

import (
	"karlc/treegame/internal/models"

	"github.com/faiface/pixel/pixelgl"
)

type TreeCamera struct {
	PosX           float64
	PosY           float64
	viewportWidth  int
	viewportHeight int
	unitScale      float64
	zoom           float64
	zoomList       []float64 // used to average out zoom values for smoother zoom
	zoomIndex      int       // used for zoomList
	OffsetY        float64
	OffsetX        float64

	renderer *Renderer
	Tree     *models.Tree
	Player   models.PhysicsActor
}

func NewTreeCamera(w, h int, scale float64, win *pixelgl.Window, player models.PhysicsActor, tree *models.Tree) *TreeCamera {

	// TODO: adjust number of avg vals
	avgVals := 400
	zoomList := make([]float64, avgVals, avgVals)
	for i := 0; i < avgVals; i++ {
		zoomList[i] = scale
	}

	_, py := player.GetPosition()

	return &TreeCamera{
		PosX:           tree.PosX,
		PosY:           py,
		viewportWidth:  w,
		viewportHeight: h,
		unitScale:      scale,
		renderer:       NewRenderer(win),
		zoomList:       zoomList,
		zoom:           scale,
		Player:         player,
		Tree:           tree,
	}
}

func (t *TreeCamera) DrawPlayer() {
	adjustedX, adjustedY := t.TranslatePosition(t.Player.GetPosition())
	adjustedW, adjustedH := t.TranslateSize(t.Player.GetSize())

	angle := t.Player.GetAngle()
	t.renderer.DrawRect(adjustedX, adjustedY, adjustedW, adjustedH, float64(angle))
}

func (t *TreeCamera) updateCameraPosition() {
	//_, aY := t.Player.GetPosition()
	//t.PosY = aY * t.zoom

	//playerPosX, _ := t.Player.GetPosition()
	//t.PosX = math.Mod(t.Tree.PosX-playerPosX, t.Tree.Circumference)
}

//func (t *TreeCamera) DrawDecor() {

//for _, d := range t.Tree.Decor {
//}
//}

func (t *TreeCamera) TranslatePosition(x, y float64) (adjustedX, adjustedY float64) {

	adjustedOffsetY := float64(t.viewportHeight) * t.OffsetY / 100
	adjustedOffsetX := float64(t.viewportHeight) * t.OffsetX / 100
	adjustedX = x*t.zoom - float64(t.PosX) + adjustedOffsetX + float64(t.viewportWidth/2)
	adjustedY = y*t.zoom - float64(t.PosY) + adjustedOffsetY + float64(t.viewportHeight/2)
	return
}

func (t *TreeCamera) TranslateSize(w, h float64) (adjustedW, adjustedH float64) {
	adjustedW = (w * t.zoom)
	adjustedH = (h * t.zoom)
	return
}

func (t *TreeCamera) DrawGame() {
	t.updateCameraPosition()
	t.DrawPlayer()
}
