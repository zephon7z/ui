package sui

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
)

func init() {
	fmt.Println("iniciando slide...")
}

type Slider struct {
	*R
	*TypeWriter
	barra   *pixel.Sprite
	frente  *Frame
	estado  Estado
	cadena  string
	val     float64
	min     float64
	max     float64
	tiempo  float64
	corners [4]bool
	conf    *Css
}

func NewSlider(s string, w float64, atlas *text.Atlas, esq [4]bool, css *Css) *Slider {
	sl := &Slider{}
	sl.R = NewR(0, 0, w, Cell_height)
	conf := entregarCss(css, CssDefaultSlider)
	sl.TypeWriter = NewTypeWriter(s, w, Cell_height, atlas, conf)
	fmt.Println(conf.Aling, "slider", CssDefaultSlider.Aling, Center)
	img := CrearImgRect(w, Cell_height+1, conf.Border_Radius, conf.Bar_color, conf.border_Color, esq)
	sl.barra = pixel.NewSprite(img, img.Bounds())
	sl.frente = NewFrame(NewR(0, 0, w, Cell_height), conf.Border_Radius, conf.Border_Width, esq, conf.border_Color, conf.Normal_color)
	sl.cadena = s
	sl.corners = esq
	sl.conf = conf
	return sl
}

func (sl *Slider) SetValues(val float64, min float64, max float64) {
	sl.val = val
	sl.min = min
	sl.max = max
}

// func (sl *Slider)() {}

func (sl *Slider) SetPos(x float64, y float64) {
	sl.X = x
	sl.Y = y
	sl.frente.SetPos(x, y)
	sl.TypeWriter.SetPos(x, y)
}

func (sl *Slider) SetSize(w float64, h float64) {
	sl.W = w
	sl.H = h
	sl.texto.SetSize(w-sl.conf.Border_Radius*2, h)
	esq := sl.corners
	conf := sl.conf
	img := CrearImgRect(w, h, conf.Border_Radius, conf.Bar_color, conf.border_Color, esq)
	sl.barra = pixel.NewSprite(img, img.Bounds())
	sl.frente.SetSize(w, h)
}

func (sl *Slider) GetPos() (float64, float64) {
	return sl.X, sl.Y
}

func (sl *Slider) GetSize() (float64, float64) {
	return sl.W, sl.H
}

func (sl *Slider) GetProp() *Css {
	return sl.conf
}

func (sl *Slider) Desactivar() {
	sl.estado = Normal
	mouse.foco = nil
}

func (sl *Slider) SetVal(x float64) {
	sl.val = x
}

/*
calcula el porcentaje de avance del slider segun donde se encuentre el
puntero al momento de ser arrastrado segun el rectangulo de colision
*/
func (sl *Slider) calcValor(pt *P) float64 {
	if pt.X < sl.X {
		return 0.0
	} else if pt.X > sl.X+sl.W {
		return 1.0
	} else {
		return (pt.X - sl.X) / sl.W
	}
}

/*
cambia el estado entre normal, write que es para ingresar el valor manualemente
y press que es cuando se arrastra para cambiar el valor
*/
func (sl *Slider) cambiarEstado(pt *P) {
	if mouse.foco != sl {
		sl.estado = Normal
		if sl.CollideP(pt) {
			sl.estado = Over
			if mouse.Press {
				if mouse.Dx != 0 {
					sl.estado = Press
					mouse.foco = sl
				}
			}
			if mouse.Soltar {
				sl.estado = write
				mouse.foco = sl
			}
		}
	}
	if mouse.foco == sl {
		if mouse.Soltar {
			sl.Desactivar()
		}
	}
}

// sirve para las funciones
func (sl *Slider) Accionar(pt *P) {
	sl.cambiarEstado(pt)
	if sl.estado == Press {
		sl.val = sl.min + (sl.max-sl.min)*sl.calcValor(pt)
	}
}

func (sl *Slider) dibBarra(target pixel.Target) {
	r := pixel.R(0, 0, sl.W*((sl.val-sl.min)/(sl.max-sl.min)), sl.H)
	sl.barra.Set(sl.barra.Picture(), r)
	mat := pixel.IM
	x := sl.X + sl.barra.Frame().W()/2
	y := sl.Y + sl.barra.Frame().H()/2 - 1
	mat = mat.Moved(pixel.V(x, y))
	sl.barra.Draw(target, mat)
}

// dibuja el valor en el centro de la slice
func (sl *Slider) dibTexto(target pixel.Target) {
	sl.texto.S = fmt.Sprintf(sl.cadena, sl.val)
	sl.TypeWriter.Dib(target)
}

func (sl *Slider) Dib(target pixel.Target) {
	sl.frente.Dib(target)
	sl.dibBarra(target)
	sl.dibTexto(target)
}
