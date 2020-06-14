package render

import (
	"fmt"
	"image"
	"karlc/treegame/internal/models"
	"math"
	"os"

	_ "image/jpeg"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	//rl "github.com/gen2brain/raylib-go/raylib"
)

// Renderer should by used to draw on screen
// Any position and scale transformations should
// be done by the camera and sent to the renderer
// for drawing only
type Renderer struct {
	window *pixelgl.Window
}

func loadSprite(path string) *pixel.Sprite {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("Error opening file '%s': %s", path, err.Error()))
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(fmt.Sprintf("Error decoding file '%s': %s", path, err.Error()))
	}
	pic, err := pixel.PictureDataFromImage(img), nil
	if err != nil {
		panic(fmt.Sprintf("Error loading picture from file '%s': %s", path, err.Error()))
	}
	return pixel.NewSprite(pic, pic.Bounds())
}

func NewRenderer(window *pixelgl.Window) *Renderer {

	//treePic, treeTile := loadPicture("./assets/sprites/treetile-small.jpg")
	//treeTile.Set(treePic, pixel.R(
	//treeTile.Frame().Min.X,
	//treeTile.Frame().Min.Y,
	//treeTile.Frame().Max.X,
	//treeTile.Frame().Max.Y),
	//)

	return &Renderer{
		window: window,
		//TreeTile: treeTile,
		//TreePic:  treePic,
	}
}

func (r *Renderer) DrawTreeTile(t *models.Tree) {

	for _, p := range t.SpriteVecs {
		m := pixel.IM.Moved(p)
		t.Sprite.Draw(r.window, m)
	}

	//r.TreeTile.Draw(r.window, pixel.IM)
}

// DrawLine between point a and b
func (r *Renderer) DrawLine(aX, aY, bX, bY float64) {
	imd := imdraw.New(nil)
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(aX, aY))
	imd.Push(pixel.V(bX, bY))
	imd.Line(1)
	imd.Draw(r.window)
}

func (r *Renderer) DrawText(x, y float64, message string) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(x, y), basicAtlas)
	fmt.Fprintln(basicTxt, message)
	basicTxt.Draw(r.window, pixel.IM.Scaled(basicTxt.Orig, 1))
}

func (r *Renderer) drawRectWithColor(x, y, w, h, a float64, c pixel.RGBA) {
	tlx, tly := x-w/2, y+h/2 // top left
	trx, try := x+w/2, y+h/2 // top right
	blx, bly := x-w/2, y-h/2 // bottom left
	brx, bry := x+w/2, y-h/2 // bottom right

	// if rect has angle, translate corner points
	if a != 0 {
		tlx, tly = rotate(x, y, tlx, tly, a)
		trx, try = rotate(x, y, trx, try, a)
		blx, bly = rotate(x, y, blx, bly, a)
		brx, bry = rotate(x, y, brx, bry, a)
	}

	imd := imdraw.New(nil)
	imd.Color = c
	imd.EndShape = imdraw.RoundEndShape
	imd.Push(pixel.V(tlx, tly))
	imd.Push(pixel.V(trx, try))
	imd.Push(pixel.V(brx, bry))
	imd.Push(pixel.V(blx, bly))
	imd.Push(pixel.V(tlx, tly))
	imd.Polygon(0)
	imd.Draw(r.window)
}

func (r *Renderer) DrawRect(x, y, w, h, a float64) {
	c := pixel.RGB(1, 1, 1)
	r.drawRectWithColor(x, y, w, h, a, c)
}

func (r *Renderer) DrawRectRed(x, y, w, h, a float64) {
	c := pixel.RGB(1, 0, 0)
	r.drawRectWithColor(x, y, w, h, a, c)
}

func (r *Renderer) DrawRectGreen(x, y, w, h, a float64) {
	c := pixel.RGB(0, 1, 0)
	r.drawRectWithColor(x, y, w, h, a, c)
}

func (r *Renderer) DrawPoint(x, y float64) {
	r.DrawRectGreen(x, y, 15, 15, 0)
}

func (r *Renderer) DrawCanvas(t *models.Tree) {

}

// rotate rotates a corner around a center point theta radians
func rotate(centerX, centerY, cornerX, cornerY, theta float64) (float64, float64) {
	tempX, tempY := cornerX-centerX, cornerY-centerY
	rotatedX := tempX*math.Cos(theta) - tempY*math.Sin(theta)
	rotatedY := tempX*math.Sin(theta) + tempY*math.Cos(theta)
	return rotatedX + centerX, rotatedY + centerY
}

// remove later, used to experiment with pixelgl easily
func (r *Renderer) Test() {
	//
}
