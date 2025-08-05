package sui

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

func init() {
	fmt.Println("cargando cursor...")
}

// este esta encargado de copiar al porta papeles
type cursor struct {
	*R
	*imdraw.IMDraw
}

func newCursor(h float64) *cursor {
	cr := cursor{}
	cr.IMDraw = imdraw.New(nil)
	cr.R = NewR(0, 0, 1, h)
	return &cr
}

func (cr *cursor) Dib(target pixel.Target) {
	cr.Clear()
	cr.Color = color.RGBA{60, 170, 255, 100}
	cr.Push(pixel.V(cr.X, cr.Y))
	cr.Push(pixel.V(cr.X+cr.W, cr.Y+cr.H))
	cr.Rectangle(0)
	cr.Draw(target)
}
