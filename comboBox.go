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
	conf *Css
}

func NewComboBox(w float64, list []string, conf *Css) *ComboBox {
	cb := &ComboBox{}
	cb.conf = entregarCss(conf, CssDefaultComboBox)
	cb.Button = NewButton(nil, "Nil", nil, AllCircularEdges, cb.conf)
	cb.mt = NewMenuItem(w, Line_height*6, nil)
	cb.Add(list...)
	cb.i = 0
	cb.Button.fn = cb.desplegarLista
	return cb
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
			mouse.foco = nil
			FocoItems = nil
		}
		//-------------- * ---------------
		cb.list = append(cb.list, et)
		cb.mt.Add(et)
		//en caso de ser el primer lemento se selecciona
		if len(cb.list) == 1 {
			cb.Selecionar(0)
		}
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

// retorna el texto seleccionado
func (cb *ComboBox) Get() string {
	return cb.list[cb.i].text.S
}

func (cb *ComboBox) desplegarLista() {
	FocoItems = []*MenuItem{cb.mt}
	cb.mt.SetPos(cb.Button.posAbs.X, cb.Button.posAbs.Y-cb.mt.H)
}

func (cb *ComboBox) Dib(target pixel.Target) {
	cb.Button.Dib(target)
	// fmt.Println(cb.mt.sizer.H, cb.mt.Superficie.Bounds().H(), "---------")
}
