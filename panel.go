package sui

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func init() {
	fmt.Println("iniciando panel...")
}

type Dimensionable interface {
	Dib(pixel.Target)
	SetPos(float64, float64)
	GetPos() (float64, float64)
	GetSize() (float64, float64)
	SetSize(float64, float64)
	GetProp() *Css
	Accionar(*P)
}

type Panel struct {
	*R
	superficie *pixelgl.Canvas
	direccion  Direction
	BtnTab     []*ToogleButton
	ListSizer  []*BoxSize
	Area       *Area
	fondo      *Frame
	posTab     Position
	Sizer      *BoxSize
	sizerTab   *BoxSize
	conf       *Css
}

func NewPanel(d Direction, posTab Position, conf *Css) *Panel {
	pl := &Panel{}
	pl.R = NewR(0, 0, 200, 200)
	pl.direccion = d
	pl.posTab = posTab
	pl.BtnTab = []*ToogleButton{} //lista que usa el toogleButton para cambiar
	pl.ListSizer = []*BoxSize{}   //lista de boxSiser por Tab que se muestran en el area
	pl.conf = entregarCss(conf, CssDefaultPanel)
	pl.Area = NewArea(100, 100, nil)
	pl.Sizer = NewBoxSizer(Horizontal, &Css{Margin: 10})
	pl.sizerTab = NewBoxSizer(Vertical, &Css{})
	pl.PosTab()
	// pl.Sizer.Add(pl.sizerTab, pl.Area)
	pl.fondo = NewFrame(pl.R, pl.conf.Border_Radius, pl.conf.Border_Width, AllCircularEdges, pl.conf.border_Color, pl.conf.Background)
	// pl.NewTab(nil, "default", 50)
	// pl.ListSizer[0] = pl.Area.sizer
	return pl
}

// agrega una nueba pestana al panel y el sizer que contiene sus elementos
func (pl *Panel) NewTab(ico *pixel.Sprite, s string, w float64) {
	var bordes [4]bool
	bordes = entregarBordes(pl.posTab)
	tb := NewToogelButton(ico, s, nil, bordes, nil)

	tb.Def = func() {
		for i, obj := range pl.BtnTab {
			if obj == tb {
				pl.Area.sizer = pl.ListSizer[i]
				pl.Area.barH.size = pl.ListSizer[i]
				pl.Area.barH.calcLargoBarra()
			}
		}
	}
	tb.SetSize(w, Line_height)

	pl.BtnTab = append(pl.BtnTab, tb)
	tb.List = &pl.BtnTab
	sizer := NewBoxSizer(Vertical, nil)

	pl.ListSizer = append(pl.ListSizer, sizer)
	pl.sizerTab.Add(tb)
	tb.List = &pl.BtnTab

	//en caso de ser la primera pestaña se borra la default y se deja como la principal
	pl.Area.sizer = pl.ListSizer[0]
	pl.BtnTab[0].Estado = Active

}

func (pl *Panel) PosTab() {
	s := NewBoxSizer(Horizontal, &Css{Margin: pl.conf.Border_Radius})
	switch pl.posTab {
	case Left:
		s.Add(pl.sizerTab, pl.Area)
	case Right:
		s.Add(pl.Area, pl.sizerTab)
	case Top:
		pl.sizerTab.Direction = Horizontal
		s.Direction = Vertical
		s.Add(pl.sizerTab, pl.Area)
	case Buttom:
		pl.sizerTab.Direction = Horizontal
		s.Direction = Vertical
		s.Add(pl.Area, pl.sizerTab)
	}
	pl.Sizer = s
}

// segun la posicion de las pestanas entrega un [4]bool para redondear el fondo
func entregarBordes(pos Position) (bordes [4]bool) {
	switch pos {
	case Left:
		bordes = LeftCircularEdges
	case Top:
		bordes = TopCircularEdges
	case Right:
		bordes = RightCircularEdges
	case Buttom:
		bordes = ButtomCircularEdges
	}
	return bordes
}

func (pl *Panel) Add(objetos ...Dimensionable) {
	pl.Area.Add(objetos...)
}

func (pl *Panel) AddIn(i int, objetos ...Dimensionable) {
	pl.ListSizer[i].Add(objetos...)
}

func (pl *Panel) SetPos(x float64, y float64) {
	pl.X = x
	pl.Y = y
	pl.Sizer.SetPos(x, y)
}

func (pl *Panel) GetPos() (float64, float64) {
	return pl.X, pl.Y
}

func (pl *Panel) SetSize(w float64, h float64) {
	pl.W = w
	pl.H = h
	pl.Sizer.SetSize(w, h)
}

func (pl *Panel) GetSize() (float64, float64) {
	if pl.W > 10 && pl.H > 10 {
		return pl.W, pl.H
	} else {
		return 0, 0
	}
}

func (pl *Panel) GetProp() *Css {
	return pl.conf
}

func (pl Panel) Redimensionar(pt *P) {
	// if pl.conf.Resize == true{
	// 	if mouse.Click && pl.CollideP(pt){
	// 		if Abs(pl.X -pt.X) < 10 || Abs(pl.X - (pl.X + pl.W))< 10 {

	// 		}
	// 	}
	// }
}

func (pl *Panel) Accionar(pt *P) {
	pl.Sizer.Accionar(pt)
}

func (pl *Panel) redimensionarArea() {
	var w, h float64
	if len(pl.BtnTab) <= 1 {
		w = pl.W - 20
		h = pl.H
	} else {
		w = pl.W - pl.sizerTab.W
		h = pl.H
	}
	pl.Area.SetSize(w, h)
}

func (pl *Panel) ubicarElementos() {
	pl.redimensionarArea()
}
func (pl *Panel) Dib(target pixel.Target) {
	if pl.W > 10 && pl.H > 10 {
		pl.fondo.Dib(target)
		pl.Sizer.Dib(target)
	}
}
