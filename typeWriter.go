package sui

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.design/x/clipboard"
)

func init() {
	fmt.Println("iniciando entry...")
	clipboard.Init()
}

type TypeWriter struct {
	texto     *Label
	cursor    *cursor
	posCursor int
	inicio    int
	Cadena    string
	tiempo    float64
	enter     func()
}

func NewTypeWriter(s string, w float64, h float64, atlas *text.Atlas, prop *Css) *TypeWriter {
	tw := &TypeWriter{}
	tw.cursor = newCursor(h)
	tw.texto = NewLabel(s, w, h, atlas, entregarCss(prop, CssDefaultEntry))
	tw.Cadena = s
	tw.texto.tresPuntos = false
	return tw
}

func (tw *TypeWriter) SetPos(x float64, y float64) {
	tw.texto.X = x
	tw.texto.Y = y
	// tw.ubicarLabel()
}

func (tw *TypeWriter) moverTexto() {
	tw.texto.S = tw.Cadena[tw.inicio:]
}

// verifica que el texto dentro quede dentro de la zona del cursor y no sobrepase el ancho del entry.
// corre la posicion de inicio donde se ve el texto
func (tw *TypeWriter) corregirInicio() {
	if tw.inicio < tw.posCursor {
		txt := tw.Cadena[tw.inicio:tw.posCursor]
		for tw.texto.W < tw.texto.txt.BoundsOf(txt).W() {
			tw.inicio++
			txt = tw.Cadena[tw.inicio:tw.posCursor]
		}

	}
	if tw.inicio > tw.posCursor {
		tw.inicio = tw.posCursor
	}
	if tw.inicio == tw.posCursor {
		tw.inicio = tw.espacioIzq()
	}
	if tw.texto.S == "" && len(tw.Cadena) > 0 {
		tw.inicio = tw.espacioIzq()
	}
	tw.moverTexto()

}

func (tw *TypeWriter) dibujarCursor(target pixel.Target) {
	c := tw.Cadena[tw.inicio:tw.posCursor]
	/*correccion porque termina la cadena con un espacio y boundsOf no cuenta el
	espacio al final para calcular el ancho*/
	if len(c) > 0 {
		if c[len(c)-1] == 32 {
			c = c + " "
		}
	}
	x := tw.texto.X + tw.texto.txt.BoundsOf(c).W()
	tw.cursor.X = x
	tw.cursor.Y = tw.texto.Y
	tw.cursor.Dib(target)
}

func (tw *TypeWriter) Dib(target pixel.Target) {
	tw.texto.Dib(target)
	if focoTypeWriter == tw {
		tw.dibujarCursor(target)
	}
}

//----------------------------------------------------------------------------------

/*
entrega el indice del espacio que esta a la izquierda del cursor, se
utiliza para hacer el control izquierda o control borrar
*/

func (tw *TypeWriter) espacioIzq() (n int) {
	for i := tw.posCursor; i > 0; {
		i-- //tiene que ir aqui porque el inicio puede ser un espacio
		n = i
		if tw.Cadena[i] == 32 {
			break
		}
	}
	return n
}

/*
entrega el indice del espacio que esta a la derecha del cursor, se
utiliza para hacer el control derecha o control suprimir
*/
func (tw *TypeWriter) espacioDer() (n int) {
	n = tw.posCursor
	for i := tw.posCursor; i < len(tw.Cadena)-1; {
		i++ //tiene que ir aqui porque el inicio puede ser un espacio
		n = i
		if tw.Cadena[i] == 32 {
			break
		} else {
			/*esta linea es para que el cursor ocupe la posicion del caracter que
			viene al finalizar la cadena*/
			if n == len(tw.Cadena)-1 {
				n = len(tw.Cadena)
			}
		}
	}
	return n
}

func (tw *TypeWriter) borrar(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyBackspace) {
		tw.borrarCaracter()
	}
	if win.Pressed(pixelgl.KeyBackspace) {
		tw.tiempo += dt
		if tw.tiempo > 0.4 {
			tw.borrarCaracter()
			tw.tiempo = 0.33
		}
	}
	if win.JustReleased(pixelgl.KeyBackspace) {
		tw.tiempo = 0
	}

	if win.JustPressed(pixelgl.KeyDelete) {
		if tw.posCursor < len(tw.Cadena) {
			tw.posCursor += 1
			tw.borrarCaracter()
		}
	}
	if win.Pressed(pixelgl.KeyDelete) {
		tw.tiempo += dt
		if tw.tiempo > 0.4 {
			if tw.posCursor < len(tw.Cadena) {
				tw.posCursor += 1
				tw.borrarCaracter()
			}
			tw.tiempo = 0.33
		}
	}
	if win.Pressed(pixelgl.KeyLeftControl) || win.Pressed(pixelgl.KeyRightControl) {
		if win.JustPressed(pixelgl.KeyBackspace) {
			i := tw.espacioIzq()
			tw.Cadena = tw.Cadena[:i] + tw.Cadena[tw.posCursor:]
			tw.posCursor = i
		}
		if win.JustPressed(pixelgl.KeyDelete) {
			i := tw.espacioDer()
			tw.Cadena = tw.Cadena[:tw.posCursor] + tw.Cadena[i:]
		}
	}
}

// simplemente borra un caracter de la cadena de texto
func (tw *TypeWriter) borrarCaracter() {
	if tw.posCursor > 0 {
		tw.Cadena = tw.Cadena[:tw.posCursor-1] + tw.Cadena[tw.posCursor:]
		tw.posCursor -= 1
	}
}

func (tw *TypeWriter) Escribir(win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyLeftControl) || win.Pressed(pixelgl.KeyRightControl) {
		if win.JustPressed(pixelgl.KeyV) {
			text := clipboard.Read(clipboard.FmtText)
			tw.Cadena = tw.Cadena[:tw.posCursor] + string(text) + tw.Cadena[tw.posCursor:]
			tw.posCursor += len(text)
		}
	} else {
		for _, r := range win.Typed() {
			s := string(byte(r))
			tw.Cadena = tw.Cadena[:tw.posCursor] + s + tw.Cadena[tw.posCursor:]
			tw.posCursor += 1
		}
	}
	if win.JustPressed(pixelgl.KeyEnter) || win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyKPEnter) {
		mouse.foco = nil
		focoTypeWriter = nil
		tw.inicio = 0
		tw.posCursor = 0
		tw.corregirInicio()
		if tw.enter != nil {
			tw.enter()
		}
	}

	tw.texto.S = tw.Cadena
	tw.moverCursor(win)
	tw.borrar(win)
	tw.corregirInicio()
}

/*
mueve el cursor con las flechas izquierda, derecha, ctrl+izquierda y
ctrl+derecha
*/
func (tw *TypeWriter) moverCursor(win *pixelgl.Window) {
	//------------------------------------------------------------------------
	//saltar cursor
	if win.Pressed(pixelgl.KeyLeftControl) || win.Pressed(pixelgl.KeyRightControl) {
		//--------------------------------------------------------------------
		//saltar cursor a la izquierda
		if win.JustPressed(pixelgl.KeyLeft) {
			tw.posCursor = tw.espacioIzq()
		}
		if win.Pressed(pixelgl.KeyLeft) {
			tw.tiempo += dt
			if tw.tiempo > 0.5 {
				if tw.posCursor > 0 {
					tw.posCursor = tw.espacioIzq()
				}
				tw.tiempo = 0.45
			}
		}
		if win.JustReleased(pixelgl.KeyLeft) {
			tw.tiempo = 0
		}
		//--------------------------------------------------------------------
		//saltar cursor a la derecha
		if win.JustPressed(pixelgl.KeyRight) {
			tw.posCursor = tw.espacioDer()
		}
		if win.Pressed(pixelgl.KeyRight) {
			tw.tiempo += dt
			if tw.tiempo > 0.5 {
				if tw.posCursor < len(tw.Cadena) {
					tw.posCursor = tw.espacioDer()
				}
				tw.tiempo = 0.45
			}
		}
		if win.JustReleased(pixelgl.KeyRight) {
			tw.tiempo = 0
		}
	} else {
		//--------------------------------------------------------------------
		//mover a la izquierda
		if win.JustPressed(pixelgl.KeyLeft) && tw.posCursor > 0 {
			tw.posCursor--
		}
		if win.Pressed(pixelgl.KeyLeft) {
			tw.tiempo += dt
			if tw.tiempo > 0.5 {
				if tw.posCursor > 0 {
					tw.posCursor -= 1
				}
				tw.tiempo = 0.45
			}
		}
		if win.JustReleased(pixelgl.KeyLeft) {
			tw.tiempo = 0
		}

		//--------------------------------------------------------------------
		//mover a la derecha
		if win.JustPressed(pixelgl.KeyRight) && tw.posCursor < len(tw.Cadena) {
			tw.posCursor++
		}

		if win.Pressed(pixelgl.KeyRight) {
			tw.tiempo += dt
			if tw.tiempo > 0.5 {
				if tw.posCursor < len(tw.Cadena) {
					tw.posCursor += 1
				}
				tw.tiempo = 0.45
			}
		}
		if win.JustReleased(pixelgl.KeyRight) {
			tw.tiempo = 0
		}
		if win.JustPressed(pixelgl.KeyEnd) {
			tw.posCursor = len(tw.Cadena)
		}
		if win.JustPressed(pixelgl.KeyHome) {
			tw.posCursor = 0
			tw.inicio = 0
		}

	}
}
