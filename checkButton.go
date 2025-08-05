package sui

import (
	"fmt"

	"github.com/faiface/pixel"
)

func init() {
	fmt.Println("iniciando checkButton...")
}

type CheckButtom struct {
	*R
	text   *Label
	cuadro *Frame
	estado Estado
	conf   *Css
}

func NewCheckButtom(w float64, s string, conf *Css) *CheckButtom {
	cb := &CheckButtom{}
	cb.conf = entregarCss(conf, CssDefaultCheckButton)
	cb.R = NewR(0, 0, w, Line_height)
	r := NewR(0, 0, Check_width, Check_width)
	cb.cuadro = NewFrame(r, cb.conf.Border_Radius, cb.conf.Border_Width, AllCircularEdges, cb.conf.border_Color, cb.conf.Normal_color)
	cb.estado = Normal
	cb.text = NewLabel(s, w-cb.cuadro.W-cb.conf.Spacing, Line_height, nil, nil)
	return cb
}

func (cb *CheckButtom) SetPos(x float64, y float64) {
	cb.X = x
	cb.Y = y
	cb.cuadro.SetPos(x, y)
	cb.ajustarTexto()
}

func (cb *CheckButtom) GetPos() (float64, float64) {
	return cb.X, cb.Y
}

func (cb *CheckButtom) SetSize(w float64, h float64) {
	cb.W = w
	cb.H = h
	cb.ajustarTexto()
}

func (cb *CheckButtom) GetSize() (float64, float64) {
	return cb.W, cb.H
}

func (cb *CheckButtom) GetProp() *Css {
	return cb.conf
}

func (cb *CheckButtom) ajustarTexto() {
	cb.text.SetSize(cb.W-cb.cuadro.W-cb.conf.Spacing, cb.H)
	cb.text.SetPos(cb.X+cb.cuadro.W+cb.conf.Spacing, cb.Y)
}

func (cb *CheckButtom) Get() bool {
	var activo bool = false
	if cb.estado == Active {
		activo = true
	}
	return activo
}

func (cb *CheckButtom) cambiarEstado(pt *P) {
	if cb.estado != Active {
		cb.estado = Normal
		if cb.CollideP(pt) {
			cb.estado = Over
			if mouse.Press {
				cb.estado = Press
			}
			if mouse.Soltar {
				cb.estado = Active
			}
		}
	} else {
		if cb.CollideP(pt) {
			if mouse.Soltar {
				cb.estado = Over
			}
		}
	}
}

func (cb *CheckButtom) Accionar(pt *P) {
	cb.cambiarEstado(pt)
	switch cb.estado {
	case Normal:
		cb.cuadro.colorBg = cb.conf.Normal_color
	case Over:
		cb.cuadro.colorBg = cb.conf.Over_color
	case Press:
		cb.cuadro.colorBg = cb.conf.Press_color
	case Active:
		cb.cuadro.colorBg = cb.conf.Active_color
	}
}

func (cb *CheckButtom) Dib(target pixel.Target) {
	cb.cuadro.Dib(target)
	cb.text.Dib(target)
}
