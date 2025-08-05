package sui

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Alineacion int

type Div struct {
	*R
	*pixelgl.Canvas
	Ancho   float64
	Alto    float64
	Padre   *Div
	Lista   []Widget
	cambiar Position
	Prop    *Css
}

// crea in Div. Es el encargado de organizar los widgets que se agregen en
// pantalla a menos que se quiera hacer de forma manual
func NewDiv(cont *Div, w float64, h float64, conf *Css) *Div {
	fmt.Println("new div")
	obj := &Div{NewR(0, 0, w, h), pixelgl.NewCanvas(pixel.R(0, 0, w, h)), w, h, cont, []Widget{}, Center, conf}
	if obj.Padre == nil {
		div_main.Add(obj)
	} else {
		obj.Padre.Add(obj)
	}
	return obj
}

// redimensiona el div y entrega los valores con los que se redimensiono ya que
// puede ocurrir que no se pueda achicar mas y retornara su minimo
func (dv *Div) SetSize(w float64, h float64) {
	dv.W = w
	dv.H = h
	dv.R.W = w
	dv.R.H = h
	dv.Canvas.SetBounds(pixel.R(0, 0, w, h))
}

// retorna el ancho y alto del div
func (dv *Div) GetSize() (float64, float64) {
	return dv.W, dv.H
}

func (dv *Div) GetPos() (float64, float64) {
	return dv.X, dv.Y
}

// agrega un widget al div para ser ordenado
func (dv *Div) Add(wg Widget) {
	dv.Lista = append(dv.Lista, wg)
}

func (dv *Div) SetPos(x float64, y float64) {
	dv.X = x
	dv.Y = y
}

// redimensiona los anchos de los elementos contenidos al maximo posible en el div
func (dv *Div) RedimAncho() {
	var (
		total            float64  = 2 * dv.Prop.Padding //suma de los anchos,espacios,padding y sin contar los redimensionables
		redimensionables []Widget = []Widget{}          //se usa para redimensionar su ancho
		avance           float64  = 0.0                 //se usa para ubicar los widgets (posicion X y es acumulativo)
		wi               float64  = 0.0                 //ancho del widget al que debe quedar
	)

	/*se calcula el ancho utilizado por los widgets y se agregan a una lista los
	redimensionables*/
	for i, wg := range dv.Lista {
		css := wg.GetProp()
		if i != len(dv.Lista)-1 {
			total += css.Spacing
		}
		if css.Resize == false {
			w, _ := wg.GetSize()
			total += w
		} else {
			redimensionables = append(redimensionables, wg)
		}
	}
	if len(redimensionables) > 0 {
		wi = (dv.W - total) / float64(len(redimensionables))
	}
	/*se ordenan los widgets (se cambia su posicion)*/
	css := dv.Prop
	avance += css.Padding

	for _, wg := range dv.Lista {
		css = wg.GetProp()
		_, h := wg.GetSize()
		if css.Resize {
			wg.SetSize(wi, h)
		}
		wg.SetPos(avance, 0)
		w, _ := wg.GetSize()
		avance += w + css.Spacing
	}
}

// redimewnsiona los altos de los elementos contenidos al posimo posible en el div
func (dv *Div) redimAlto() {

}

func (dv *Div) GetProp() *Css {
	return dv.Prop
}

// sirve para cambiar el tamaño del div cuando se hace click y se arrastra
func (dv *Div) Redimensionar() {
	rd := NewR(dv.X+dv.W-5, dv.Y, 5, dv.H)
	ri := NewR(dv.X, dv.Y, 5, dv.H)
	if dv.CollideP(mouse.P) && mouse.Click {
		if rd.CollideP(mouse.P) {
			dv.cambiar = Left

		}
		if ri.CollideP(mouse.P) {
			dv.cambiar = Right
		}
	}
	if mouse.Soltar {
		dv.cambiar = Center
	}

	if dv.cambiar != Center {
		p := dv.Prop
		switch dv.cambiar {
		case Left:
			w := mouse.X - dv.X
			if w < p.Min {
				w = p.Min
			}
			dv.SetSize(w, dv.H)
		case Right:
			w := dv.X + dv.W - mouse.X
			if w < p.Min {
				w = p.Min
			}
			dv.SetSize(w, dv.H)

		}
	}
}

// ajusta el tamaño del widget cuando cambia su tamaño
func (dv *Div) Ajustar() {
	//dimensiones que se usaran para dibujar el div
	var resto, expandidos, minimo float64
	switch dv.Prop.Arrangement {
	case Horizontal:
		resto = dv.W - dv.Prop.Margin*2
		for i, wg := range dv.Lista {
			wi, _ := wg.GetSize()
			p := wg.GetProp()
			minimo += p.Min
			if i < len(dv.Lista)-1 {
				resto -= p.Spacing
			}
			if p.Expand {
				expandidos++
			} else {
				resto -= wi
			}
		}
	case Vertical:
		resto = dv.H - dv.Prop.Margin*2
		for _, wg := range dv.Lista {
			_, hi := wg.GetSize()
			p := wg.GetProp()
			resto -= hi - p.Spacing
			if p.Expand {
				expandidos++
			}
		}
	}

	//************** se dibujan los elementos del div *************
	//se cambia el origen
	dv.SetMatrix(pixel.IM.Moved(pixel.V(0, dv.H)))
	//se dibujan los elementos contenidos
	var x, y float64
	x = dv.Prop.Margin
	y = dv.Prop.Margin
	for _, wg := range dv.Lista {
		if dv.Prop.Arrangement == Horizontal {
			if wg.GetProp().Expand {
				// ancho, _ := wg.GetSize()
				w := resto / expandidos
				p := wg.GetProp()
				if w < p.Min {
					w = p.Min
				}
				wg.SetSize(w, dv.H)
			}
			wg.SetPos(x, 0)
		} else {
			wg.SetPos(0, y)
		}
		wg.Dib(dv)
		//calculo de la posicion x e y de los contenidos
		ancho, alto := wg.GetSize()
		pr := wg.GetProp()
		x += ancho + pr.Spacing
		y += alto + pr.Spacing
	}
}

// muestra el Div y los elementos en su lista de widgets
func (dv *Div) Dib(lienzo pixel.Target) {
	// //se limpia el lienzo
	// dv.Clear(dv.Prop.Background)
	// // se redimensiona el div
	// if dv.Prop.Resize {
	// 	dv.Redimensionar()
	// }
	// dv.Ajustar()
	// //*************** se dibuja el div ******************************
	// mat := pixel.IM
	// if dv == div_main {
	// 	mat = mat.Moved(pixel.V(dv.W/2, dv.H/2))
	// } else {
	// 	mat = mat.Moved(pixel.V(dv.X+dv.W/2, -dv.Y-dv.H/2))
	// }

	// dv.Draw(lienzo, mat)
}
