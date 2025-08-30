package sui

import (
	"fmt"

	"github.com/faiface/pixel"
)

func init() {
	fmt.Println("iniciando sizer...")
}

type BoxSize struct {
	*R
	marco *R
	Direction
	Objetos []Dimensionable
	dy      float64
	fondo   *Frame
	frente  *Frame
	Conf    *Css
}

func NewBoxSizer(direccion Direction, conf *Css) *BoxSize {
	bs := &BoxSize{}
	l := 10.
	bs.R = NewR(0, 0, l, l)
	bs.marco = NewR(0, 0, l, l)
	bs.Direction = direccion
	bs.Objetos = []Dimensionable{}
	bs.Conf = entregarCss(conf, CssDefaultBoxSize)
	// bs.fondo = NewFrame(bs.R, 0, 1, NotCircularEdges, Blanco, gris)
	// bs.frente = NewFrame(bs.marco, 0, 1, NotCircularEdges, Blanco, grafito)
	return bs
}

func (bs *BoxSize) Add(objetos ...Dimensionable) {
	bs.Objetos = append(bs.Objetos, objetos...)
	// bs.minimizar()
	bs.ubicarObjetos()

}

func (bs *BoxSize) Remove(i int) {
	bs.Objetos = append(bs.Objetos[:i], bs.Objetos[i+1:]...)
}

func (bs *BoxSize) Clear() {
	bs.Objetos = []Dimensionable{}
}

func (bs *BoxSize) SetPos(x float64, y float64) {
	bs.X = x
	bs.Y = y
	bs.ubicarObjetos()
}

func (bs *BoxSize) GetPos() (float64, float64) {
	return bs.W, bs.H
}

func (bs *BoxSize) SetSize(w float64, h float64) {
	bs.W = w
	bs.H = h
	bs.ajustarMarco()
	bs.ubicarObjetos()
}

func (bs *BoxSize) GetSize() (float64, float64) {
	return bs.W, bs.H
}

func (bs *BoxSize) GetProp() *Css {
	return bs.Conf
}

func (bs *BoxSize) ajustarMarco() {
	ml, mr := bs.margenesHorizontales()
	mt, mb := bs.margenesVerticales()
	bs.marco.W = bs.W - ml - mr
	bs.marco.H = bs.H - mt - mb
	bs.marco.X = bs.X + ml
	bs.marco.Y = bs.Y + mb
}

// entrega el margen izquierdo y derecho
func (bs *BoxSize) margenesHorizontales() (left float64, right float64) {
	conf := bs.Conf
	if conf.Margin > conf.Margin_left {
		left = conf.Margin
	} else {
		left = conf.Margin_left
	}

	if conf.Margin > conf.Margin_Right {
		right = conf.Margin
	} else {
		right = conf.Margin_Right
	}
	return left, right
}

// entrega el margen superior e inferior
func (bs *BoxSize) margenesVerticales() (top float64, buttom float64) {
	conf := bs.Conf
	if conf.Margin > conf.Margin_Top {
		top = conf.Margin
	} else {
		top = conf.Margin_Top
	}

	if conf.Margin > conf.Margin_Bottom {
		buttom = conf.Margin
	} else {
		buttom = conf.Margin_Bottom
	}
	return top, buttom
}

// entrega la dimension minima posible del sizer
func (bs *BoxSize) minimizar() {
	conf := bs.Conf    //configuracion de elementos
	sp := conf.Spacing // espaciado de elementos
	bs.marco.W = 0
	bs.marco.H = 0
	for _, obj := range bs.Objetos {
		w, h := obj.GetSize()
		switch bs.Direction {
		case Vertical:
			bs.marco.H += h + sp
			if bs.marco.W < w {
				bs.marco.W = w
			}
		case Horizontal:
			bs.marco.W += w + sp
			if bs.marco.H < h {
				bs.marco.H = h
			}
		}
	}

	switch bs.Direction {
	case Vertical:
		bs.marco.H -= sp
	case Horizontal:
		bs.marco.W -= sp
	}
	l, r := bs.margenesHorizontales()
	t, b := bs.margenesVerticales()
	bs.W = l + r + bs.marco.W
	bs.H = t + b + bs.marco.H
}

// entrega el espacio disponible para que se expandan los objetos segun la direccion del sizer
func (bs *BoxSize) Disponible() (total float64, cant float64) {
	sp := bs.Conf.Spacing
	switch bs.Direction {
	case Vertical:
		total = bs.marco.H
	case Horizontal:
		total = bs.marco.W
	}
	for _, obj := range bs.Objetos {
		conf := obj.GetProp()
		w, h := obj.GetSize()
		switch bs.Direction {
		case Vertical:
			if !conf.Expand && !conf.Expand_height {
				total -= h
			} else {
				cant++
			}
		case Horizontal:
			if !conf.Expand && !conf.Expand_width {
				total -= w
			} else {
				cant++
			}
		}
	}
	total -= sp * float64(len(bs.Objetos)-1)
	if total < 0 {
		total = 0
	}
	return total, cant
}

// cambia el ancho del sizer y los elementos contenidos
func (bs *BoxSize) redimensionar() {

}

// ubica y redimensiona los objetos dentro del marco segun sus propiedades
func (bs *BoxSize) ubicarObjetos() {
	var maxW, maxH float64
	/*maxW y maxH se usan para comparar si Height o el Width del marco son menores y
	se cambia de tamaño el sizer, pero siempre que no se expanda en la direccion de la
	variable*/

	bs.ajustarMarco()
	x := bs.marco.X
	y := bs.marco.Y + bs.marco.H
	total, n := bs.Disponible()
	sp := bs.Conf.Spacing
	for _, obj := range bs.Objetos {
		conf := obj.GetProp()
		w, h := obj.GetSize()
		switch bs.Direction {
		case Vertical:
			if conf.Expand {
				obj.SetSize(bs.marco.W, total/n)
			} else if conf.Expand_width {
				obj.SetSize(bs.marco.W, h)
			} else if conf.Expand_height {
				obj.SetSize(w, total/n)
			}

			y -= h
			switch conf.Aling {
			case Center:
				obj.SetPos(x+(bs.marco.W-w)/2, y)
			case Right:
				obj.SetPos(x+bs.marco.W-w, y)
			default:
				obj.SetPos(x, y)
			}
			y -= sp //tiene que ir sp despues porque es el espacio para el siguiente objeto
		case Horizontal:
			if conf.Expand {
				obj.SetSize(total/n, bs.marco.H)
			} else if conf.Expand_width {
				obj.SetSize(total/n, h)
			} else if conf.Expand_height {
				obj.SetSize(w, bs.marco.H)
			}
			w, h = obj.GetSize()
			switch conf.Aling {
			case Center:
				obj.SetPos(x, y-h-(bs.marco.H-h)/2)
			case Buttom:
				obj.SetPos(x, y-bs.marco.H)
			default:
				obj.SetPos(x, y-h)
			}
			x += w + sp
		}

		w, h = obj.GetSize()
		switch bs.Direction {
		case Vertical:
			maxH += h + sp
			if maxW < w {
				maxW = w
			}
		case Horizontal:
			maxW += w + sp
			if maxH < h {
				maxH = h
			}
		}
		//se podria refactorizar un poco para eliminar los 2 ultimos switch
		//-----------------------------------
		// maxH += hf
		// if maxW < wf {
		// 	fmt.Println(wf, maxW, "---")
		// 	maxW = wf
		// }
		//-----------------------------------
	}
	switch bs.Direction {
	case Vertical:
		maxH -= sp
	case Horizontal:
		maxW -= sp
	}
	if maxH > bs.marco.H && (!bs.Conf.Expand && !bs.Conf.Expand_height) {
		mt, mb := bs.margenesVerticales()
		bs.H = maxH + mt + mb
	}
	if maxW > bs.marco.W && (!bs.Conf.Expand && !bs.Conf.Expand_width) {
		ml, mr := bs.margenesHorizontales()
		bs.W = maxW + ml + mr
	}
	bs.ajustarMarco()
}

func (bs *BoxSize) Accionar(pt *P) {
	// if bs == SizerPrincipal {
	// 	w, h := SizerPrincipal.Objetos[0].GetSize()
	// 	fmt.Println(len(bs.Objetos), w, h)
	// }
	for _, obj := range bs.Objetos {
		obj.Accionar(pt)
	}
}

func (bs *BoxSize) Dib(target pixel.Target) {
	// bs.fondo.Dib(target)
	// bs.frente.Dib(target)
	for _, obj := range bs.Objetos {
		obj.Dib(target)
	}
}
