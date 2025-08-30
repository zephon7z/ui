package sui

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func init() {
	fmt.Println("iniciando area...")
}

type Area struct {
	*R
	rSup       *R //rectangulo de la superficie
	superficie *pixelgl.Canvas
	sizer      *BoxSize
	barW       *spaceBar
	barH       *spaceBar
	conf       *Css
}

func NewArea(w float64, h float64, conf *Css) *Area {
	ar := &Area{}
	ar.R = NewR(0, 0, w, h)
	ar.rSup = NewR(0, 0, w, h)
	ar.superficie = pixelgl.NewCanvas(pixel.R(0, 0, w, h))
	ar.sizer = NewBoxSizer(Vertical, &Css{Margin: 5, Spacing: 2, Expand_width: true})
	ar.barH = NewSpaceBar(h, ar.sizer, ar.rSup, Vertical, nil)
	ar.conf = entregarCss(conf, CssdefaultArea)
	return ar
}

func (ar *Area) Add(lista ...Dimensionable) {
	ar.sizer.Add(lista...)
	ar.ajustarArea()
}

func (ar *Area) Remove(i int) {
	ar.sizer.Remove(i)
	ar.barH.calcLargoBarra()
}

func (ar *Area) SetPos(x float64, y float64) {
	ar.X = x
	ar.Y = y
	ar.barH.SetPos(x+ar.W-ar.barH.W, ar.Y)
	ar.rSup.X = x
	ar.rSup.Y = y
	ar.ajustarArea()
}

func (ar *Area) GetPos() (float64, float64) {
	return ar.X, ar.Y
}

func (ar *Area) SetSize(w float64, h float64) {
	ar.W = w
	ar.H = h
	ar.ajustarArea()
	ar.barH.calcLargoBarra()
}

func (ar *Area) GetSize() (float64, float64) {
	return ar.W, ar.H
}

func (ar *Area) GetProp() *Css {
	return ar.conf
}

func (ar *Area) dibSizer(target pixel.Target) {
	y := ar.superficie.Bounds().H() - ar.sizer.H + ar.sizer.dy
	if y != ar.sizer.Y {
		ar.sizer.SetPos(0, y)
	}
	ar.sizer.Dib(ar.superficie)
}

func (ar *Area) dibSuperficie(target pixel.Target) {
	mat := pixel.IM
	x := ar.X + ar.superficie.Bounds().W()/2
	y := ar.Y + ar.superficie.Bounds().H()/2
	mat = mat.Moved(pixel.V(x, y))
	ar.superficie.Draw(target, mat)
}

/*
ajusta la superficie de dibujo para dar espacio a la scrollbar dentro del ancho total,
redimensiona la scrollbar, ademas redimensiona el rect que se usa para saber si
colisiona el mouse hacer funcionar los widgets
*/
func (ar *Area) ajustarArea() {
	barH := (ar.sizer.H > ar.H)
	//se da espacio para la superficie y se deja espacio para la scrollbar
	if barH {
		ar.rSup.W = ar.W - ar.barH.W
		ar.rSup.H = ar.H
		ar.superficie.SetBounds(pixel.R(0, 0, ar.rSup.W, ar.rSup.H))
	} else {
		ar.rSup.W = ar.W
		ar.rSup.H = ar.H
		ar.superficie.SetBounds(pixel.R(0, 0, ar.W, ar.H))
	}
	w := ar.superficie.Bounds().W()
	_, h := ar.sizer.GetSize()
	ar.barH.SetSize(Bar_width, ar.superficie.Bounds().H())
	ar.sizer.SetSize(w, h)
}

func (ar *Area) dibBar(target pixel.Target) {
	barH := ar.sizer.H > ar.H
	if barH {
		ar.barH.Dib(target)
	}
}

func (ar *Area) Accionar(pt *P) {
	// barH := ar.sizer.H > ar.H
	// if barH {
	// }
	ar.barH.Accionar(pt)
	// if ar.rSup.CollideP(pt) {
	pti := &P{}
	pti.X = pt.X - ar.X
	pti.Y = pt.Y - ar.Y
	ar.sizer.Accionar(pti)
	// }
	/*Al agreegar el condicional de que choque el puntero con el area provoca bugs en
	el combobox y en otros porque no se actualizan los estados de los objetos
	contenidos */
}

func (ar *Area) Dib(target pixel.Target) {
	// fmt.Println(ar.sizer.W, ar.W, ar.superficie.Bounds().W(), ar.barH.W)
	ar.superficie.Clear(ar.conf.Background)
	ar.dibSizer(target)
	ar.dibSuperficie(target)
	ar.dibBar(target)
}
