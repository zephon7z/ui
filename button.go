package sui

import (
	"fmt"

	"github.com/faiface/pixel"
)

func init() {
	fmt.Println("iniciando boton...")
}

type Estado int

const (
	Normal Estado = iota
	Over
	Press
	Disable
	Active
	write
)

type Button struct {
	*R
	posAbs  *P
	Texto   *Label
	Estado  Estado
	fondo   *Frame
	corners [4]bool
	fn      func()
	icon    *pixel.Sprite
	conf    *Css
}

func NewButton(icon *pixel.Sprite, s string, fn func(), esq [4]bool, clase *Css) *Button {
	btn := &Button{}
	btn.conf = entregarCss(clase, CssDefaultBoton)
	btn.R = NewR(0, 0, 100, Line_height)
	btn.posAbs = &P{0, 0}
	btn.Texto = NewLabel(s, 100, Line_height, nil, btn.conf)
	btn.Estado = Normal
	btn.corners = esq
	btn.fn = fn
	btn.icon = icon
	btn.fondo = NewFrame(btn.R, btn.conf.Border_Radius, btn.conf.Border_Width, esq, btn.conf.border_Color, btn.conf.Normal_color)
	return btn
}

// func (btn *Button) crearFrames() (lista map[Estado]*Frame) {
// 	prop := btn.conf
// 	lista = map[Estado]*Frame{}
// 	lista[Normal] = NewFrame(btn.R, prop.Border_Radius, prop.Border_Width, btn.corners, prop.border_Color, prop.Normal_color)
// 	lista[Over] = NewFrame(btn.R, prop.Border_Radius, prop.Border_Width, btn.corners, prop.border_Color, prop.Over_color)
// 	lista[Press] = NewFrame(btn.R, prop.Border_Radius, prop.Border_Width, btn.corners, prop.border_Color, prop.Press_color)
// 	return lista
// }

func (btn *Button) SetFn(def func()) {
	btn.fn = def
}

// cambia la posicion del boton
func (btn *Button) SetPos(x float64, y float64) {
	btn.X = x
	btn.Y = y
	btn.Texto.SetPos(x+btn.conf.Border_Radius, y)
}

// cambia las dimensiones del boton
func (btn *Button) SetSize(w float64, h float64) {
	btn.W = w
	btn.H = h
	if btn.W < 10 {
		btn.W = 10
	}
	var ancho, wIcon float64
	if btn.icon != nil {
		wIcon = btn.icon.Frame().W()
	}
	ancho = btn.W - btn.conf.Border_Radius*2 - wIcon
	if ancho < 0 {
		ancho = 0
	}
	btn.Texto.SetSize(ancho, btn.H)

}

// entrega la posicion del boton
func (btn *Button) GetPos() (float64, float64) {
	return btn.X, btn.Y
}

// entrega el ancho y alto del boton
func (btn *Button) GetSize() (float64, float64) {
	return btn.W, btn.H
}

// entrega los estilos del boton
func (btn *Button) GetProp() *Css {
	return btn.conf
}

func (btn *Button) desactivar() {
}

func (btn *Button) ubicarTexto() {
	var w float64 = 0
	if btn.icon != nil {
		w = btn.icon.Frame().W()
	}
	correccion := 0.8
	btn.Texto.X = btn.X + btn.conf.Border_Radius + w
	btn.Texto.Y = btn.Y + (btn.H-btn.Texto.atlas.LineHeight()*correccion)/2
}

// ejecurta la funcion que contiene el boton
func (btn *Button) ejecutar() {
	if btn.fn != nil {
		btn.fn()
	} else {
		fmt.Println("no hay funcion que aplicar")
	}
}

// cambia el estado del boton y ejecuta la funcion al hacer click
func (btn *Button) Accionar(pt *P) {
	btn.Estado = Normal
	if mouse.foco == nil {
		if btn.CollideP(pt) {
			btn.posAbs.X = mouse.X - (pt.X - btn.X)
			btn.posAbs.Y = mouse.Y - (pt.Y - btn.Y)
			btn.Estado = Over
			if mouse.Press {
				btn.Estado = Press
			}
			if mouse.Soltar {
				btn.ejecutar()
				btn.Estado = Press
			}
		}
	}

}
func (btn *Button) dibfondo(target pixel.Target) {
	switch btn.Estado {
	case Normal:
		btn.fondo.colorBg = btn.conf.Normal_color
	case Over:
		btn.fondo.colorBg = btn.conf.Over_color
	case Press:
		btn.fondo.colorBg = btn.conf.Press_color
	}
	btn.fondo.Dib(target)
	// btn.fondos[btn.Estado].Dib(target)
}

// dibuja el boton en pantalla
func (btn *Button) Dib(target pixel.Target) {
	btn.dibfondo(target)
	btn.Texto.Dib(target)
}
