package sui

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func init() {
	fmt.Println("iniciando superficie...")
}

type Surface struct {
	*R
	*pixelgl.Canvas
	fondo *Frame
	conf  *Css
}

func NewSurface(w float64, h float64, conf *Css) *Surface {
	sf := &Surface{}
	sf.R = NewR(0, 0, w, h)
	sf.Canvas = pixelgl.NewCanvas(pixel.R(0, 0, w, h))
	sf.fondo = NewFrame(sf.R, 5, 1, AllCircularEdges, negro, gris_claro)
	sf.conf = conf
	return sf
}

func (sf *Surface) SetSize(w float64, h float64) {
	sf.W = w
	sf.H = h
	sf.Canvas.SetBounds(pixel.R(0, 0, w, h))
}

func (sf *Surface) GetSize() (float64, float64) {
	return sf.W, sf.H
}

func (sf *Surface) SetPos(x float64, y float64) {
	sf.X = x
	sf.Y = y
}

func (sf *Surface) GetPos() (float64, float64) {
	return sf.X, sf.Y
}

func (sf *Surface) Accionar(pt *P) {

}

func (sf *Surface) CollideMouse() (choca bool) {
	if sf.CollideP(mouse.P) {
		choca = true
	}
	return choca
}

// entrega la posicion relativa del mouse con respecto al lienzo
func (sf *Surface) MousePos() *P {
	x := mouse.X - sf.X
	y := mouse.Y - sf.Y
	return &P{x, y}
}

func (sf *Surface) GetProp() *Css {
	return sf.conf
}

func (sf *Surface) Dib(target pixel.Target) {
	sf.fondo.Dib(target)
	// sf.Canvas.Clear(transparente)
	mat := pixel.IM
	x := sf.X + sf.Canvas.Bounds().W()/2
	y := sf.Y + sf.Canvas.Bounds().H()/2
	mat = mat.Moved(pixel.V(x, y))
	sf.Canvas.Draw(target, mat)
}
