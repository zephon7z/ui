package sui

import (
	"fmt"

	"github.com/faiface/pixel"
)

func init() {
	fmt.Println("iniciando item...")
}

type MenuItem struct {
	*R
	fondo *Frame
	area  *Area
	conf  *Css
}

func NewMenuItem(w float64, h float64, conf *Css) *MenuItem {
	mi := &MenuItem{}
	mi.R = NewR(0, 0, w, h)
	mi.conf = entregarCss(conf, CssDefaultMenuItem)
	mi.area = NewArea(w-1, h-mi.conf.Border_Radius*2, nil)
	mi.fondo = NewFrame(mi.R, mi.conf.Border_Radius, mi.conf.Border_Width, AllCircularEdges, mi.conf.border_Color, mi.conf.Background)
	return mi
}

func (mi *MenuItem) SetPos(x float64, y float64) {
	mi.X = x
	mi.Y = y
	mi.area.SetPos(x, y+mi.conf.Border_Radius)
}

func (mi *MenuItem) SetSize(w float64, h float64) {
	mi.W = w
	mi.H = h
	mi.area.SetSize(w-1, h-mi.conf.Border_Radius*2)
}

func (mi *MenuItem) Add(objetos ...Dimensionable) {
	mi.area.Add(objetos...)
}

// sirve para accionar las acciones de los widgets contenidos
func (mi *MenuItem) Accionar(pt *P) {
	mi.area.Accionar(pt)
}

func (mi *MenuItem) Dib(target pixel.Target) {
	mi.fondo.Dib(target)
	mi.area.Dib(target)
}
