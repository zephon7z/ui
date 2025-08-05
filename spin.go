package sui

import (
	"fmt"

	"strconv"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
)

func init() {
	fmt.Println("iniciando spinCtrl...")
}

type Spin struct {
	*R
	*TypeWriter
	centroR    *R
	cadenaS    string
	Val        float64
	incremento float64
	tiempo     float64
	btnIzq     *Button
	btnDer     *Button
	estado     Estado
	fondos     map[Estado]*Frame
	corners    [4]bool
	conf       *Css
}

func NewSpin(w float64, s string, atlas *text.Atlas, esq [4]bool, conf *Css) *Spin {
	Cell_height := Cell_height
	sp := &Spin{}
	sp.conf = entregarCss(conf, CssDefaultSpin)
	sp.R = NewR(0, 0, w, Cell_height)
	sp.TypeWriter = NewTypeWriter(s, w-40, Cell_height, atlas, sp.conf)
	sp.TypeWriter.enter = sp.quitarFoco
	sp.cadenaS = s
	sp.corners = esq
	sp.incremento = 1.
	esqI := [4]bool{esq[0], false, esq[1], false}
	sp.btnIzq = NewButton(nil, "<", sp.restaVal, esqI, sp.conf)
	sp.btnIzq.SetSize(20, Cell_height)
	esqD := [4]bool{false, esq[2], false, esq[3]}
	sp.btnDer = NewButton(nil, ">", sp.sumarVal, esqD, sp.conf)
	sp.btnDer.SetSize(20, Cell_height)
	sp.fondos = map[Estado]*Frame{
		Normal: NewFrame(sp.R, sp.conf.Border_Radius, sp.conf.Border_Width, esq, sp.conf.border_Color, sp.conf.Normal_color),
		Over:   NewFrame(sp.R, sp.conf.Border_Radius, sp.conf.Border_Width, esq, sp.conf.border_Color, sp.conf.Over_color),
		Press:  NewFrame(sp.R, sp.conf.Border_Radius, sp.conf.Border_Width, esq, sp.conf.border_Color, sp.conf.Active_color),
		Active: NewFrame(sp.R, sp.conf.Border_Radius, sp.conf.Border_Width, esq, sp.conf.border_Color, sp.conf.Active_color),
	}
	sp.centroR = NewR(0, 0, w, Cell_height)
	sp.setCentroR()
	return sp
}

func (sp *Spin) setCentroR() {
	sp.centroR.X = sp.X + sp.btnIzq.W
	sp.centroR.Y = sp.Y
	sp.centroR.W = sp.W - sp.btnIzq.W - sp.btnDer.W
	sp.centroR.H = sp.H
}

func (sp *Spin) SetPos(x float64, y float64) {
	sp.X = x
	sp.Y = y
	sp.btnIzq.SetPos(x, y)
	sp.btnDer.SetPos(x+sp.W-sp.btnIzq.W, y)
	sp.TypeWriter.SetPos(x+sp.btnDer.W, y)
	sp.centroR.X = sp.X + sp.btnIzq.W
	sp.centroR.Y = sp.Y
}

func (sp *Spin) GetPos() (float64, float64) {
	return sp.X, sp.Y
}

func (sp *Spin) SetSize(w float64, h float64) {
	sp.W = w
	sp.H = h
	sp.texto.SetSize(w-sp.conf.Border_Radius*2, h)
	sp.setCentroR()
}

func (sp *Spin) GetSize() (float64, float64) {
	return sp.W, sp.H
}

func (sp *Spin) GetProp() *Css {
	return sp.conf
}

func (sp *Spin) sumarVal() {
	sp.Val += sp.incremento
}

func (sp *Spin) restaVal() {
	sp.Val -= sp.incremento
}

func (sp *Spin) cambiarEstado(pt *P) {
	if mouse.foco == nil {
		sp.estado = Normal
		if sp.centroR.CollideP(pt) {
			sp.estado = Over
			if mouse.Soltar {
				sp.estado = Active
				sp.darFoco()
			}
			if mouse.Press && mouse.Dx != 0 {
				sp.estado = Press
				mouse.foco = sp
			}
		}
	}
	if mouse.foco == sp {
		if sp.estado == Press {
			sp.estado = Press
			sp.btnDer.Estado = Press
			sp.btnIzq.Estado = Press
			if mouse.Soltar {
				sp.quitarFoco()
			}
		} else {
			sp.estado = Active

		}
		if !sp.centroR.CollideP(pt) && mouse.Click {
			sp.quitarFoco()
		}

	}
}

func (sp *Spin) Accionar(pt *P) {
	sp.btnIzq.Accionar(pt)
	sp.btnDer.Accionar(pt)
	sp.cambiarEstado(pt)
	if sp.estado == Press {
		if mouse.Dx != 0 {
			sp.Val += (mouse.Dx * sp.incremento) / 3
		}

	}
}

func (sp *Spin) darFoco() {
	focoTypeWriter = sp.TypeWriter
	mouse.foco = sp
	sp.posCursor = 0
	sp.TypeWriter.Cadena = ""
}

func (sp *Spin) quitarFoco() {
	sp.estado = Normal
	focoTypeWriter = nil
	mouse.foco = nil
	val, err := strconv.ParseFloat(sp.Cadena, 64)
	if err == nil {
		sp.Val = val
		sp.Cadena = ""
	}
	sp.texto.S = fmt.Sprintln(sp.cadenaS, sp.Val)
}

func (sp *Spin) Dib(target pixel.Target) {
	sp.fondos[sp.estado].Dib(target)
	sp.btnIzq.Dib(target)
	sp.btnDer.Dib(target)
	if focoTypeWriter != sp.TypeWriter {
		sp.texto.S = fmt.Sprintf(sp.cadenaS, sp.Val)
	}
	sp.TypeWriter.Dib(target)
}
