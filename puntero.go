package sui

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Puntero struct {
	*P
	Click    bool
	Soltar   bool
	Press    bool
	FijarPos bool
	FijarH   bool
	Objeto   Widget
	Dx       float64
	Dy       float64
	Scroll   float64
	foco     any
}

type Widget interface {
	SetPos(float64, float64)
	GetProp() *Css
	GetPos() (float64, float64)
	GetSize() (float64, float64)
	SetSize(float64, float64)
	Dib(pixel.Target)
}

func NewPuntero() *Puntero {
	fmt.Println("se creo un nuevo puntero")

	p := &Puntero{}
	p.P = &P{}
	return p
}

func (pt *Puntero) Detectar(win *pixelgl.Window) {
	//actualizacion de datos del puntero
	pt.X = win.MousePosition().X
	pt.Y = win.MousePosition().Y
	pt.Dx = pt.X - win.MousePreviousPosition().X
	pt.Dy = pt.Y - win.MousePreviousPosition().Y
	//se fija el movimiento del mouse
	if mouse.FijarPos {
		p := win.MousePreviousPosition()
		win.SetMousePosition(pixel.V(p.X, p.Y))
		pt.X = win.MousePosition().X
		pt.Y = win.MousePosition().Y
	}
	if mouse.FijarH {
		p := win.MousePreviousPosition()
		win.SetMousePosition(pixel.V(win.MousePosition().X, p.Y))
		pt.Y = win.MousePosition().Y
	}

	pt.Soltar = false
	if win.Pressed(pixelgl.MouseButtonLeft) {
		pt.Press = true
		pt.Soltar = false
		pt.Click = false
	}
	if win.JustPressed(pixelgl.MouseButtonLeft) {
		pt.Click = true
		pt.Soltar = false
	}
	if win.JustReleased(pixelgl.MouseButtonLeft) {
		pt.Press = false
		pt.Click = false
		pt.Soltar = true
		// pt.foco = nil
	}
	mouse.Scroll = win.MouseScroll().Y
}
