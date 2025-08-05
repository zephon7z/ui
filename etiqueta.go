package sui

import (
	"fmt"

	"github.com/faiface/pixel"
)

func init() {
	fmt.Println("inciando etiqueta...")
}

type etiqueta struct {
	*R
	text   *Label
	label  *Label
	icon   *pixel.Sprite
	fondo  *Frame
	sel    bool
	estado Estado
	Def    func()
	conf   *Css
}

func NewEtiqueta(w float64, h float64, icon *pixel.Sprite, s string, la string, conf *Css) *etiqueta {
	et := &etiqueta{}
	et.R = NewR(0, 0, w, h)
	et.conf = entregarCss(conf, CssDefaultEtiqueta)
	et.text = NewLabel(s, 100, h, nil, et.conf)
	et.text.tresPuntos = false
	et.text.AjustarAlAncho()
	et.label = NewLabel(la, 100, h, nil, nil)
	et.label.tresPuntos = false
	et.label.AjustarAlAncho()
	et.estado = Normal
	et.fondo = NewFrame(NewR(0, 0, w, h), et.conf.Border_Radius, et.conf.Border_Width, AllCircularEdges, et.conf.border_Color, et.conf.Normal_color)
	et.icon = icon
	return et
}

func (et etiqueta) GetProp() *Css {
	return et.conf
}

func (et *etiqueta) GetPos() (float64, float64) {
	return et.X, et.Y
}

func (et *etiqueta) SetPos(x float64, y float64) {
	et.X = x
	et.Y = y
	et.ubicarElementos()
}

func (et *etiqueta) SetSize(w float64, h float64) {
	et.W = w
	et.H = h
	et.fondo.SetSize(w, h)
	et.ubicarElementos()
}

func (et etiqueta) GetSize() (float64, float64) {
	return et.W, et.H
}

func (et *etiqueta) cambiarEstado(pt *P) {
	if et.sel {
		et.estado = Active
	} else {
		et.estado = Normal
	}
	if mouse.foco == nil {
		if et.CollideP(pt) && et.estado != Active {

			et.estado = Over
			if mouse.Press {
				et.estado = Press
			}
			if mouse.Soltar {
				if et.Def != nil {
					et.Def()
				} else {
					fmt.Println("no hay funcion")
				}
			}
		}
	}
}

func (et etiqueta) Accionar(pt *P) {
	et.cambiarEstado(pt)
	switch et.estado {
	case Normal:
		et.fondo.colorBg = et.conf.Normal_color
	case Over:
		et.fondo.colorBg = et.conf.Over_color
	case Press:
		et.fondo.colorBg = et.conf.Press_color
	case Active:
		et.fondo.colorBg = et.conf.Active_color
	}
	// et.fondo.colorBg = et.conf.Normal_color
	// if et.CollideP(pt) {
	// 	et.fondo.colorBg = et.conf.Over_color
	// 	if mouse.Press {
	// 		et.fondo.colorBg = et.conf.Press_color
	// 	}
	// 	if mouse.Soltar && et.Def != nil {
	// 		et.Def()
	// 	}
	// }
	// if et.sel {
	// 	et.fondo.colorBg = et.conf.Active_color
	// }
}

func (et *etiqueta) ubicarElementos() {
	et.fondo.SetPos(et.X, et.Y)
	var x float64
	x = et.conf.Border_Radius + et.X + et.conf.Margin
	if et.icon != nil {
		x += et.icon.Frame().W()
	}
	et.text.SetPos(x, et.Y)

	et.label.X = et.X + et.W - et.label.W - et.conf.Border_Radius - et.conf.Margin
	et.label.Y = et.Y
}

func (et *etiqueta) Dib(target pixel.Target) {
	et.fondo.Dib(target)
	if et.icon != nil {
		mat := pixel.IM
		mat = mat.Moved(pixel.V(et.X, et.Y))
		et.icon.Draw(target, mat)
	}
	et.text.Dib(target)
	et.label.Dib(target)
}
