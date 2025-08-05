package sui

import (
	"fmt"

	// "image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"

	// "golang.org/x/image/colornames"
	"golang.org/x/image/font/gofont/goregular"
)

func init() {
	fmt.Println("iniciando label...")
}

type Label struct {
	*R
	S          string
	i          int
	txt        *text.Text
	atlas      *text.Atlas
	conf       *Css
	tresPuntos bool
}

func NewLabel(s string, w float64, h float64, atlas *text.Atlas, conf *Css) *Label {
	la := &Label{}
	la.S = s
	la.i = len(s) - 1
	la.conf = entregarCss(conf, CssDefaultLabel)
	la.atlas = entregarAtlas(atlas)
	la.txt = text.New(pixel.ZV, la.atlas)
	la.R = NewR(0, 0, w, h)
	la.moverI(w)
	la.escribir()
	la.tresPuntos = true
	return la
}

func entregarAtlas(atlas *text.Atlas) (at *text.Atlas) {
	if atlas != nil {
		at = atlas
	} else {
		f, _ := truetype.Parse(goregular.TTF)
		face := truetype.NewFace(f, &truetype.Options{Size: TextSize})
		at = text.NewAtlas(face, text.ASCII)
	}
	return at
}

func (la *Label) SetPos(x float64, y float64) {
	la.X = x
	la.Y = y
}

func (la *Label) GetPos() (float64, float64) {
	return la.X, la.Y
}
func (la *Label) GetSize() (float64, float64) {
	return la.W, la.H
}
func (la *Label) SetSize(w float64, h float64) {
	la.moverI(w)
	la.W = w
	la.escribir()
	la.H = h
}
func (la *Label) GetProp() *Css {
	return la.conf
}

// solo esta por requisito
func (la *Label) Accionar(p *P) {}

// ajusta el ancho del label para que calce el string
func (la *Label) AjustarAlAncho() {
	var w float64
	w = la.txt.BoundsOf(la.S).W()
	la.SetSize(w, la.H)
}

/*
mueve el valor i para acortar el texto de manera de no sobrepasar el ancho maximo
*/
func (la *Label) moverI(w float64) {
	la.i = len(la.S) - 1
	if la.i < 0 {
		la.i = 0
	}
	if la.tresPuntos {
		w = w - la.txt.BoundsOf("...").W()
	}
	if la.txt.BoundsOf(la.S[:la.i]).W() > w {
		for la.txt.BoundsOf(la.S[:la.i]).W() > w {
			if la.S[:la.i] == "" {
				break
			}
			la.i--
		}
	}
}

// escribe la cadena de texto hasta su i caracter de manera que sobrepase su ancho
func (la *Label) escribir() {
	la.moverI(la.W)
	la.txt.Clear()
	la.txt.Color = la.conf.Text_color
	if la.i == len(la.S)-1 {
		la.txt.WriteString(la.S)
	} else {
		if la.i > 0 {
			if la.tresPuntos {
				la.txt.WriteString(la.S[:la.i] + "...")
			} else {
				la.txt.WriteString(la.S[:la.i])
			}
		} else { //en case de que el texto sea muy corto no se escribe
			la.txt.WriteString("")
		}
	}
	//acatualizacion del ancho de texto
	// la.W = la.txt.Bounds().W()
}

// dibuja el label en la superficie especificada
func (la *Label) Dib(target pixel.Target) {
	la.escribir()
	mat := pixel.IM
	correccion := 0.8
	y := la.Y + (la.H-la.txt.LineHeight*correccion)/2
	switch la.conf.Text_Aling {
	case Right:
		mat = mat.Moved(pixel.V(la.X+la.W-la.txt.Bounds().W(), y))
	case Center:
		mat = mat.Moved(pixel.V(la.X+(la.W-la.txt.Bounds().W())/2, y))
	case correrCentro:
		mat = mat.Moved(pixel.V(la.X-la.txt.Bounds().W()/2, y))
	case correrDerecha:
		mat = mat.Moved(pixel.V(la.X-la.txt.Bounds().W(), y))
	default:
		mat = mat.Moved(pixel.V(la.X, y))
	}
	la.txt.Draw(target, mat)
}
