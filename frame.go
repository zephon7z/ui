package sui

import (
	"fmt"
	"image/color"
	. "math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

func init() {
	fmt.Println("iniciando frame...")
}

type Frame struct {
	*imdraw.IMDraw
	*R
	Radio     float64
	Espesor   float64
	Esquinas  [4]bool
	colorLine color.Color
	colorBg   color.Color
}

func NewFrame(r *R, radio float64, espesor float64, esq [4]bool, ln color.Color, bg color.Color) *Frame {
	fr := &Frame{}
	fr.IMDraw = imdraw.New(nil)
	fr.R = r
	fr.Radio = radio
	fr.Espesor = espesor
	fr.Esquinas = esq
	fr.colorLine = ln
	fr.colorBg = bg
	return fr
}

func (fr *Frame) SetPos(x float64, y float64) {
	fr.X = x
	fr.Y = y
}

func (fr *Frame) SetSize(w float64, h float64) {
	fr.W = w
	fr.H = h
}

func (fr *Frame) SetBgColor(bg color.Color) {
	fr.colorBg = bg
}

func (fr *Frame) Dib(targer pixel.Target) {
	fr.Clear()

	//forma del termino de la linea o formas
	fr.EndShape = imdraw.RoundEndShape

	//BORDES REDONDEADOS

	//radio de las esquinas
	var r1, r2, r3, r4 float64
	if fr.Esquinas[0] {
		r1 = fr.Radio
	} else {
		r1 = 0
	}
	if fr.Esquinas[1] {
		r2 = fr.Radio
	} else {
		r2 = 0
	}
	if fr.Esquinas[2] {
		r3 = fr.Radio
	} else {
		r3 = 0
	}
	if fr.Esquinas[3] {
		r4 = fr.Radio
	} else {
		r4 = 0
	}
	// ESQUINAS

	//esquina inferior izquierda
	fr.Color = fr.colorBg //fondo
	fr.Push(pixel.V(fr.X+r3, fr.Y+r3))
	fr.CircleArc(r3, Pi, 1.5*Pi, 0)
	fr.Color = fr.colorLine //linea
	fr.Push(pixel.V(fr.X+r3, fr.Y+r3))
	fr.CircleArc(r3, Pi, 1.5*Pi, fr.Espesor)
	//esquina superior izquierda
	fr.Color = fr.colorBg
	fr.Push(pixel.V(fr.X+r1, fr.Y+fr.H-r1))
	fr.CircleArc(r1, 0.5*Pi, Pi, 0)
	fr.Color = fr.colorLine
	fr.Push(pixel.V(fr.X+r1, fr.Y+fr.H-r1))
	fr.CircleArc(r1, 0.5*Pi, Pi, fr.Espesor)
	//esquina superior derecha
	fr.Color = fr.colorBg
	fr.Push(pixel.V(fr.X+fr.W-r2, fr.Y+fr.H-r2))
	fr.CircleArc(r2, 0, 0.5*Pi, 0)
	fr.Color = fr.colorLine
	fr.Push(pixel.V(fr.X+fr.W-r2, fr.Y+fr.H-r2))
	fr.CircleArc(r2, 0, 0.5*Pi, fr.Espesor)
	//esquina inferior derecha
	fr.Color = fr.colorBg
	fr.Push(pixel.V(fr.X+fr.W-r4, fr.Y+r4))
	fr.CircleArc(r4, 1.5*Pi, 2*Pi, 0)
	fr.Color = fr.colorLine
	fr.Push(pixel.V(fr.X+fr.W-r4, fr.Y+r4))
	fr.CircleArc(r4, 1.5*Pi, 2*Pi, fr.Espesor)

	//PARTE RECTANGULAR

	//fondo
	fr.Color = fr.colorBg
	fr.Push(pixel.V(fr.X, fr.Y+r3))
	fr.Push(pixel.V(fr.X, fr.Y+fr.H-r1))
	fr.Push(pixel.V(fr.X+r1, fr.Y+fr.H))
	fr.Push(pixel.V(fr.X+fr.W-r2, fr.Y+fr.H))
	fr.Push(pixel.V(fr.X+fr.W, fr.Y+fr.H-r2))
	fr.Push(pixel.V(fr.X+fr.W, fr.Y+r4))
	fr.Push(pixel.V(fr.X+fr.W-r4, fr.Y))
	fr.Push(pixel.V(fr.X+r3, fr.Y))
	fr.Polygon(0)
	//bordes
	fr.Color = fr.colorLine //linea vertical izquierda
	fr.Push(pixel.V(fr.X, fr.Y+r3))
	fr.Push(pixel.V(fr.X, fr.Y+fr.H-r1))
	fr.Line(fr.Espesor)                    // linea horizontal superior
	fr.Push(pixel.V(fr.X+r1-1, fr.Y+fr.H)) //EL -1 ES POR QUE NO CONSIDERA EL FON DEL CIRCULO
	fr.Push(pixel.V(fr.X+fr.W-r2, fr.Y+fr.H))
	fr.Line(fr.Espesor) //linea vertical derecha
	fr.Push(pixel.V(fr.X+fr.W, fr.Y+fr.H-r2))
	fr.Push(pixel.V(fr.X+fr.W, fr.Y+r4))
	fr.Line(fr.Espesor)               // linea horizontal inferior
	fr.Push(pixel.V(fr.X+r3-1, fr.Y)) //EL -1 ES POR QUE NO CONSIDERA EL FON DEL CIRCULO
	fr.Push(pixel.V(fr.X+fr.W-r4, fr.Y))
	fr.Line(fr.Espesor)

	fr.Draw(targer)
}
