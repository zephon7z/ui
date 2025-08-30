package sui

import (
	"fmt"

	"github.com/faiface/pixel"
)

func init() {
	fmt.Println("iniciando combobox... ")
}

type ComboBox struct {
	*Button
	list []*etiqueta
	i    int
	mt   *MenuItem
	Def  func()
	conf *Css
}

func NewComboBox(w float64, fn func(), esq [4]bool, conf *Css) *ComboBox {
	cb := &ComboBox{}
	cb.conf = entregarCss(conf, CssDefaultComboBox)
	cb.Button = NewButton(nil, "Nil", nil, esq, cb.conf)
	cb.Button.SetSize(w, Line_height)
	cb.mt = NewMenuItem(w, Line_height*6, nil)
	cb.list = []*etiqueta{}
	cb.i = 0
	cb.Def = fn
	cb.Button.fn = cb.desplegarLista
	return cb
}

func (cb *ComboBox) SetInt(i int) {
	cb.i = i
	for j, et := range cb.list {
		if i == j {
			et.estado = Active
			et.sel = true
		} else {
			et.estado = Normal
			et.sel = false
		}
	}
	cb.Texto.S = cb.list[i].text.S
}

func (cb *ComboBox) SetItem(s string, i int) {
	cb.list[i].text.S = s
	if cb.i == i {
		cb.Texto.S = s
	}
}

func (cb *ComboBox) SetSurface(w float64, h float64) {
	cb.mt.SetSize(w, h)
}

func (cb *ComboBox) Add(list ...string) {
	for _, s := range list {
		et := NewEtiqueta(cb.W, cb.H, nil, s, "", nil)
		//funcion anonima para poder dejar seleccionado un elemento
		et.Def = func() {
			et.sel = true
			for i, obj := range cb.list {
				if obj != et {
					obj.sel = false
				} else {
					cb.i = i
					cb.Texto.S = obj.text.S
				}
			}
			if cb.Def != nil {
				cb.Def()
			}
			mouse.foco = nil
			FocoItems = nil
		}
		//-------------- * ---------------
		cb.list = append(cb.list, et)
		cb.mt.Add(et)
	}
}

func (cb *ComboBox) Selecionar(i int) {
	for _, et := range cb.list {
		et.sel = false
	}
	cb.list[i].sel = true
	cb.i = i
	cb.Texto.S = cb.list[i].text.S
}

func (cb *ComboBox) Remove(i int) {
	cb.list = append(cb.list[:i-1], cb.list[i:]...)
	// cb.mt.sizer.Clear()
	cb.AddEtiqueta(cb.list...)
}

// agrega una celda al combobox
func (cb *ComboBox) AddEtiqueta(objetos ...*etiqueta) {
	for _, et := range objetos {
		cb.mt.Add(et)
	}
}

func (cb *ComboBox) GetList() (lista []string) {
	for _, et := range cb.list {
		lista = append(lista, et.text.S)
	}
	return lista
}

// retorna el texto seleccionado
func (cb *ComboBox) GetString() string {
	return cb.list[cb.i].text.S
}

func (cb *ComboBox) GetInt() int {
	return cb.i
}

// elimina a todos lo objetos contenidos en el comboBox
func (cb *ComboBox) Clear() {
	cb.list = []*etiqueta{}
	cb.mt.area.sizer.Clear()
}

func (cb *ComboBox) desplegarLista() {
	FocoItems = []*MenuItem{cb.mt}
	cb.mt.SetPos(cb.Button.posAbs.X, cb.Button.posAbs.Y-cb.mt.H)
}

func (cb *ComboBox) Dib(target pixel.Target) {
	cb.Button.Dib(target)
	// fmt.Println(cb.mt.sizer.H, cb.mt.Superficie.Bounds().H(), "---------")
}
