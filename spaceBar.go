package sui

import (
	"fmt"

	"github.com/faiface/pixel"
)

func init() {
	fmt.Println("iniciando spaceBar...")
}

type spaceBar struct {
	*Frame
	bar          *Frame
	estado       Estado
	Area         *R
	size         *BoxSize
	val          float64
	max          float64
	posInitMouse float64
	posInitBar   float64
	orientacion  Direction
	conf         *Css
}

func NewSpaceBar(h float64, bs *BoxSize, area *R, orien Direction, conf *Css) *spaceBar {
	sb := &spaceBar{}
	sb.orientacion = orien
	conf = entregarCss(conf, CssDefaultSpaceBar)
	sb.conf = conf
	sb.size = bs
	sb.Area = area
	r1 := NewR(0, 0, Bar_width, h)
	sb.Frame = NewFrame(r1, conf.Border_Radius, conf.Border_Width, AllCircularEdges, conf.border_Color, conf.Background)
	r2 := NewR(0, 0, Bar_width, 10)
	sb.bar = NewFrame(r2, conf.Border_Radius, conf.Border_Width, AllCircularEdges, conf.border_Color, conf.Normal_color)
	sb.val = 0
	sb.max = 300
	sb.orientacion = Horizontal
	sb.calcLargoBarra()
	return sb
}

func (sb *spaceBar) SetPos(x float64, y float64) {
	sb.X = x
	sb.Y = y
	sb.posicionarBarra()
}

/*
cacula el largo de la barra desplazadora segun la proporcion de lo que se ve y el
tamaño del contenido, ademas calcula los valores max que se puede alcanzar.
*/
func (sb *spaceBar) calcLargoBarra() {
	/*porcentaje minimo que puede tener la barra en relacion al largo*/
	min := 0.07 //%
	largo := sb.H * sb.H / sb.size.H
	if largo < sb.H*min {
		largo = sb.H * min
	}
	sb.bar.H = largo
	sb.max = sb.size.H - sb.Area.H
}

func (sb *spaceBar) SetSize(w float64, h float64) {
	sb.Frame.SetSize(w, h)
	sb.calcLargoBarra()
}

// calcula la posicion de la barra segun el Dy del BoxSizer (sb.size)
func (sb *spaceBar) posicionarBarra() {
	sb.bar.X = sb.X
	sb.bar.Y = sb.Y + sb.H - sb.bar.H + (sb.H-sb.bar.H)*(sb.size.dy/sb.max)
}

func (sb *spaceBar) cambiarEstado(pt *P) {
	if mouse.foco == nil {
		sb.estado = Normal
		if sb.bar.CollideP(pt) {
			sb.estado = Over
			if mouse.Click {
				mouse.foco = sb
				sb.posInitMouse = pt.Y
				sb.posInitBar = sb.bar.Y
				sb.estado = Press
			}
		}
	}
	if mouse.foco == sb {
		sb.estado = Press
		if mouse.Soltar {
			mouse.foco = nil
			sb.estado = Normal
		}
	}
}

func (sb *spaceBar) cambiarColor() {
	switch sb.estado {
	case Normal:
		sb.bar.SetBgColor(sb.conf.Normal_color)
	case Press:
		sb.bar.SetBgColor(sb.conf.Press_color)
	case Over:
		sb.bar.SetBgColor(sb.conf.Over_color)
	}
}

func (sb *spaceBar) porcentaje() float64 {
	return ((sb.Y + sb.H) - (sb.bar.Y + sb.bar.H)) / (sb.H - sb.bar.H)
}

func (sb *spaceBar) Accionar(pt *P) {
	sb.cambiarEstado(pt)
	sb.moverBar(pt)
	sb.scroll(pt)
}

// mueve la barra de desplazamiento cuando se arrastra
func (sb *spaceBar) moverBar(pt *P) {
	//se mueve la barra desplazadora
	if mouse.foco == sb {
		//se ubica la barra a su nueva posicion
		sb.bar.Y = sb.posInitBar + (pt.Y - sb.posInitMouse)
		// -------- inicio restricciones --------
		//restricciones para que la barra no se salga de los limites
		if sb.bar.Y < sb.Y {
			sb.bar.Y = sb.Y
		}
		if sb.bar.Y+sb.bar.H > sb.Y+sb.H {
			sb.bar.Y = sb.Y + sb.H - sb.bar.H
		}
		//----------- fin restricciones ----------
		//se calcula el dy con el desenso de la barra
		sb.size.dy = sb.max * sb.porcentaje()
	}
}

/*Ajusta el parametro DY de boxSize para que se dibujen sus elementos*/
func (sb *spaceBar) scroll(pt *P) {
	if mouse.Scroll != 0 && sb.Area.CollideP(pt) && sb.size.H > sb.Area.H {
		sb.size.dy += Line_height * -mouse.Scroll
		//-----restricciones --------
		if sb.size.dy < 0 {
			sb.size.dy = 0
		}
		if sb.size.dy > sb.max {
			sb.size.dy = sb.max
		}
	}
}

func (sb *spaceBar) ubicarBarPorDy() {
	sb.bar.Y = sb.Y + sb.H - sb.bar.H - (sb.H-sb.bar.H)*sb.size.dy/sb.max
}

func (sb *spaceBar) Dib(target pixel.Target) {
	sb.cambiarColor()
	sb.Frame.Dib(target)
	sb.ubicarBarPorDy()
	sb.bar.Dib(target)
}

/*se podria cambiar el sistema de cambiar colores en los otros widgets, porque
se consume mas memoria aunque debe ser minima*/
