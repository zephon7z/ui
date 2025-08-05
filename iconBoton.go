package sui

import (
	"fmt"

	"github.com/faiface/pixel"
)

func init() {
	fmt.Println("iniciando iconBoton...")
}

type IconButton struct {
	*ToogleButton
	mt   *MenuItem
	conf *Css
}

func NewIconButton(w float64, icon *pixel.Sprite, s string, conf *Css) *IconButton {
	prop := entregarCss(conf, CssDefaultIconButton)
	ib := &IconButton{}
	ib.ToogleButton = NewToogelButton(icon, s, nil, AllCircularEdges, prop)
	ib.conf = entregarCss(conf, CssDefaultComboBox)
	ib.mt = NewMenuItem(200, 200, nil)
	ib.fn = ib.desplegarLista
	return ib
}

func (ib *IconButton) Add(list ...Dimensionable) {
	ib.mt.Add(list...)
}

func (ib *IconButton) SetSizeSurface(w float64, h float64) {
	ib.mt.SetSize(w, h)
}

// agrega una celda al combobox
func (ib *IconButton) AddEtiqueta(objetos ...*etiqueta) {
	for _, et := range objetos {
		ib.mt.Add(et)
	}
}

func (ib *IconButton) desplegarLista() {
	FocoItems = []*MenuItem{ib.mt}
	ib.mt.SetPos(100, 100)
	ib.mt.SetPos(ib.ToogleButton.posAbs.X, ib.ToogleButton.posAbs.Y-ib.mt.H)
	ib.ToogleButton.Estado = Active
}

func (ib *IconButton) Accionar(pt *P) {
	ib.ToogleButton.Accionar(pt)
	if FocoItems == nil && ib.Estado == Active {
		ib.Estado = Normal
	}

}

func (ib *IconButton) Dib(target pixel.Target) {
	ib.ToogleButton.Dib(target)
	// fmt.Println(ib.mt.sizer.H, ib.mt.Superficie.Bounds().H(), "---------")
}
