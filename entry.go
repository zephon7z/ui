package sui

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
)

func init() {
	fmt.Println("iniciando entry...")
}

type Entry struct {
	*R
	*TypeWriter
	corners [4]bool
	estado  Estado
	fondo   map[Estado]*Frame
	conf    *Css
}

func NewLEntry(s string, w float64, atlas *text.Atlas, esq [4]bool, conf *Css) *Entry {
	en := &Entry{}
	en.R = NewR(0, 0, w, Line_height)
	prop := entregarCss(conf, CssDefaultEntry)
	en.TypeWriter = NewTypeWriter(s, w, Line_height, atlas, prop)
	en.corners = esq
	en.conf = entregarCss(CssDefaultEntry, conf)
	en.fondo = en.crearFrames()
	en.corregirAnchoLabel()
	return en
}

func (en *Entry) SetPos(x float64, y float64) {
	en.X = x
	en.Y = y
	en.texto.SetPos(x+en.conf.Border_Radius, y)
}

func (en *Entry) GetPos() (float64, float64) {
	return en.X, en.Y
}

func (en *Entry) SetSize(w float64, h float64) {
	en.W = w
	en.H = h
	en.texto.SetSize(en.W-en.conf.Border_Radius*2, en.H)
	en.fondo[Normal].SetSize(w, h)
	en.fondo[Over].SetSize(w, h)
	en.fondo[Press].SetSize(w, h)
}

func (en *Entry) GetSize() (float64, float64) {
	return en.W, en.H
}

func (en *Entry) GetProp() *Css {
	return en.conf
}

func (en *Entry) Desactivar() {
	mouse.foco = nil
	focoTypeWriter = nil
	en.inicio = 0
	en.posCursor = 0
	en.TypeWriter.corregirInicio()
}
func (en *Entry) crearFrames() (dic map[Estado]*Frame) {
	prop := en.conf
	dic = map[Estado]*Frame{}
	dic[Normal] = NewFrame(en.R, prop.Border_Radius, prop.Border_Width, en.corners, prop.border_Color, prop.Normal_color)
	dic[Over] = NewFrame(en.R, prop.Border_Radius, prop.Border_Width, en.corners, prop.border_Color, prop.Over_color)
	dic[Press] = NewFrame(en.R, prop.Border_Radius, prop.Border_Width, en.corners, prop.border_Color, prop.Press_color)
	return dic
}

// cambia el ancho maximo del texto para que no choque con los bordes redondeados
func (en *Entry) corregirAnchoLabel() {
	w := en.W - en.conf.Border_Radius*2
	_, h := en.texto.GetSize()
	en.texto.SetSize(w, h)
}

func (en *Entry) Accionar(pt *P) {
	en.estado = Normal
	if mouse.foco == nil {
		if en.CollideP(pt) {
			en.estado = Over
			if mouse.Click {
				en.estado = Press
				mouse.foco = en
				focoTypeWriter = en.TypeWriter
				en.posCursor = len(en.Cadena)
				en.TypeWriter.corregirInicio()
			}
		}
	}
	if mouse.foco == en {
		en.estado = Press
		if !en.CollideP(pt) {
			if mouse.Click {
				en.Desactivar()
			}
		}
	}
}

func (en *Entry) Dib(target pixel.Target) {
	en.fondo[en.estado].Dib(target)
	en.TypeWriter.Dib(target)
}
