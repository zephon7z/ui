package sui

import (
	"fmt"

	"github.com/faiface/pixel"
)

func init() {
	fmt.Println("iniciando listBox...")
}

type ListBox struct {
	*Frame
	Area *Area
	list []*etiqueta
	i    int
	Def  func()
	conf *Css
}

func NewListBox(w float64, h float64, esq [4]bool, conf *Css) *ListBox {
	lb := &ListBox{}
	lb.conf = entregarCss(conf, CssDefaultListBox)
	r := NewR(0, 0, w, h)
	lb.Area = NewArea(w-1, h-lb.conf.Border_Radius*2, nil)
	lb.Frame = NewFrame(r, lb.conf.Border_Radius, lb.conf.Border_Width, esq, lb.conf.border_Color, lb.conf.Background)
	lb.list = []*etiqueta{}
	return lb
}

func (lb *ListBox) SetItem(s string, i int) {
	lb.list[i].text.S = s
}

func (lb *ListBox) Add(lista ...string) {
	for _, nombre := range lista {
		et := NewEtiqueta(lb.W, Line_height, nil, nombre, "", nil)
		et.Def = func() {
			et.sel = true
			for i, obj := range lb.list {
				if obj != et {
					obj.sel = false
				} else {
					lb.i = i
				}
			}
			mouse.foco = nil
			FocoItems = nil
		}
		lb.Area.Add(et)
		lb.list = append(lb.list, et)
	}
}

func (lb *ListBox) Remove(i int) {
	lb.Area.Remove(i)
	lb.list = append(lb.list[:i], lb.list[i+1:]...)
	if lb.i > len(lb.list)-1 {
		lb.i = len(lb.list) - 1
	}
	lb.Area.sizer.minimizar()
	// lb.Area.sizer.minimizar()
}

func (lb *ListBox) Clear() {
	lb.Area.sizer.Objetos = []Dimensionable{}
	lb.list = []*etiqueta{}
}

func (lb *ListBox) SetPos(x float64, y float64) {
	lb.Frame.SetPos(x, y)
	lb.Area.SetPos(x, y+lb.conf.Border_Radius)
}

func (lb *ListBox) GetPos() (float64, float64) {
	return lb.X, lb.Y
}

func (lb *ListBox) SetSize(w float64, h float64) {
	lb.Area.SetSize(w, h-lb.conf.Border_Radius*2)
	lb.Frame.SetSize(w, h)
}

func (lb *ListBox) GetProp() *Css {
	return lb.conf
}

func (lb *ListBox) GetSize() (float64, float64) {
	return lb.W, lb.H
}

func (lb *ListBox) GetInt() int {
	return lb.i
}

func (lb *ListBox) GetString() string {
	if len(lb.list) > 0 {
		return lb.list[lb.i].text.S
	} else {
		return ""
	}
}

func (lb *ListBox) Accionar(pt *P) {
	lb.Area.Accionar(pt)
}

func (lb *ListBox) SetI(i int) {
	if i < 0 {
		i = 0
	}
	if i > len(lb.list)-1 {
		i = len(lb.list) - 1
	}
	lb.i = i
	for j, et := range lb.list {
		if j == i {
			et.sel = true
		} else {
			et.sel = false
		}
	}
}

func (lb *ListBox) Dib(targer pixel.Target) {
	lb.Frame.Dib(targer)
	lb.Area.Dib(targer)
}
