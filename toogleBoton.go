package sui

import (
	"fmt"

	"github.com/faiface/pixel"
)

func init() {
	fmt.Println("iciciando toogleBoton...")
}

type ToogleButton struct {
	*R
	posAbs  *P
	Texto   *Label
	Estado  Estado
	fondo   *Frame
	corners [4]bool
	Def     func()
	icon    *pixel.Sprite
	List    *[]*ToogleButton
	conf    *Css
}

func NewToogelButton(icon *pixel.Sprite, s string, fn func(), esq [4]bool, conf *Css) *ToogleButton {
	tbtn := &ToogleButton{}
	tbtn.R = NewR(0, 0, 100, Line_height)
	tbtn.posAbs = &P{0, 0}
	tbtn.Estado = Normal
	tbtn.corners = esq
	tbtn.Def = fn
	tbtn.icon = icon
	tbtn.conf = entregarCss(conf, CssDefaultToogleButton)
	tbtn.Texto = NewLabel(s, 100, Line_height, nil, tbtn.conf)
	tbtn.List = &[]*ToogleButton{}
	tbtn.fondo = NewFrame(
		tbtn.R,
		tbtn.conf.Border_Radius,
		tbtn.conf.Border_Width,
		esq,
		tbtn.conf.border_Color,
		tbtn.conf.Normal_color,
	)
	return tbtn
}

// func (tbtn *ToogleButton) Add(listTbtn ...*ToogleButton) {
// 	for _, obj := range listTbtn {
// 		tbtn.List = append(tbtn.List, obj)
// 	}
// }

func (tbtn *ToogleButton) SetPos(x float64, y float64) {
	tbtn.X = x
	tbtn.Y = y
	tbtn.ajustarTexto()
}

func (tbtn *ToogleButton) ajustarTexto() {
	var wIcon float64
	if tbtn.icon != nil {
		wIcon = tbtn.icon.Frame().W()
	}
	tbtn.Texto.SetSize(tbtn.W-wIcon, tbtn.H)
	tbtn.Texto.SetPos(tbtn.X+wIcon, tbtn.Y)

}

func (tbtn *ToogleButton) GetPos() (float64, float64) {
	return tbtn.X, tbtn.Y
}

func (tbtn *ToogleButton) SetSize(w float64, h float64) {
	tbtn.W = w
	tbtn.H = h
	tbtn.ajustarTexto()
}

func (tbtn *ToogleButton) GetSize() (float64, float64) {
	return tbtn.W, tbtn.H
}

func (tbtn *ToogleButton) GetProp() *Css {
	return tbtn.conf
}

func (tbtn *ToogleButton) dibIcon(target pixel.Target) {
	if tbtn.icon != nil {
		mat := pixel.IM
		mat = mat.Moved(pixel.V(tbtn.X, tbtn.Y))
		tbtn.icon.Draw(target, mat)
	}
}

func (tbtn *ToogleButton) Get() bool {
	var usado bool = false
	if tbtn.Estado == Active {
		usado = true
	}
	return usado
}

func (tbtn *ToogleButton) cambiarEstado(pt *P) {
	if tbtn.Estado != Active {
		tbtn.Estado = Normal
		if tbtn.CollideP(pt) {
			tbtn.posAbs.X = mouse.X - (pt.X - tbtn.X)
			tbtn.posAbs.Y = mouse.Y - (pt.Y - tbtn.Y)
			tbtn.Estado = Over
			if mouse.Press {
				tbtn.Estado = Press
			}
			if mouse.Soltar {
				//deja a toda la lista en estado normal
				for _, obj := range *tbtn.List {
					obj.Estado = Normal
				}
				//deja solamente al actual en activo
				tbtn.Estado = Active

				if tbtn.Def != nil {
					tbtn.Def()
				}
			}
		}
	} else {
		if len(*tbtn.List) == 0 {
			if tbtn.CollideP(pt) && mouse.Soltar {
				tbtn.Estado = Normal
			}

		}
	}
}

func (tbtn *ToogleButton) Accionar(pt *P) {
	tbtn.cambiarEstado(pt)
	switch tbtn.Estado {
	case Normal:
		tbtn.fondo.colorBg = tbtn.conf.Normal_color
	case Over:
		tbtn.fondo.colorBg = tbtn.conf.Over_color
	case Press:
		tbtn.fondo.colorBg = tbtn.conf.Press_color
	case Active:
		tbtn.fondo.colorBg = tbtn.conf.Active_color
	}
}

func (tbtn *ToogleButton) Dib(target pixel.Target) {
	tbtn.fondo.Dib(target)
	tbtn.dibIcon(target)
	tbtn.Texto.Dib(target)
}
