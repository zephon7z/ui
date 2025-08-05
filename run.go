package sui

import (
	"fmt"

	// "image/color"

	"github.com/faiface/pixel/pixelgl"
)

var (
	W              float64 = 100.
	H              float64 = 100.
	mouse                  = NewPuntero()
	div_main       *Div
	SizerPrincipal *BoxSize
	lienzo         *pixelgl.Canvas
	focoTypeWriter *TypeWriter
	FocoItems      []*MenuItem
	dt             float64
)

func init() {
	fmt.Println("iniciando run ... ok")
	//hay que crear primero una div main primero para cambiar el eje de dibujado y a este se agregan los div sin padre
	SizerPrincipal = NewBoxSizer(Horizontal, nil)
	SizerPrincipal.SetSize(4000, 4000)
}

// dibuja y hace funcionar a los widgets
func RunApp(win *pixelgl.Window, deltaTiempo float64) {
	dt = deltaTiempo
	w := win.Bounds().W()
	h := win.Bounds().H()
	verfificarDimension(w, h)
	SizerPrincipal.Dib(win)
	mouse.Detectar(win)
	if len(FocoItems) > 0 {
		//ejecuta las acciones de los items
		accionarItems()
	} else {
		//ejecuta las acciones de los widgets en los sizer de la ventana principal
		SizerPrincipal.Accionar(mouse.P)
	}

	//hace funcionar el capturador de las entradas de teclado en los entry y similares
	if focoTypeWriter != nil {
		focoTypeWriter.Escribir(win)
	}
	ctx.Clear()

	//se dibujan los items
	for _, obj := range FocoItems {
		obj.Dib(win)
	}
	// ctx.Draw(win)
	div_main.Dib(win)
}

func accionarItems() {
	//cierra los menus o items emergentes
	if mouse.Soltar && mouse.foco == nil {
		cerrar := true
		for _, obj := range FocoItems {
			if obj.CollideP(mouse.P) {
				cerrar = false
			}
		}
		if cerrar {
			FocoItems = nil
		}
	}
	if FocoItems != nil {
		//hace funcionar los widgets dentro de los items
		FocoItems[len(FocoItems)-1].Accionar(mouse.P)
	}
}

func verfificarDimension(w float64, h float64) {
	if SizerPrincipal.W != w || SizerPrincipal.H != h {
		SizerPrincipal.SetSize(w, h)
	}
}
