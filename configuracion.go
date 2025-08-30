package sui

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"golang.org/x/image/font"
)

type Position int
type Direction int
type Display int

const (
	Left Position = iota
	Right
	Center
	Top
	Buttom
	correrCentro
	correrDerecha

	Horizontal Direction = iota
	Vertical

	Ajustar Display = iota
	Expand
	Expand_W
	Expand_H
	None
)

var (
	Bar_width   float64 = 10
	Line_height float64 = 20
	Check_width float64 = 16
	TextSize    float64 = 12
	Cell_height float64 = 20
)

func init() {
	fmt.Println("modulo de configuracion... ok")
}

type Basic struct {
	Border_Width       float64
	Border_Radius      float64
	Background         color.Color
	Foreground         color.Color
	Backgroung_Img     pixel.Picture
	border_Color       color.Color
	Active_color       color.Color
	Normal_color       color.Color
	Over_color         color.Color
	Press_color        color.Color
	Disable_color      color.Color
	Bar_color          color.Color
	Text_over_color    color.Color
	Text_press_color   color.Color
	Text_disable_color color.Color
	Text_color_label   color.Color
	Text_color         color.Color
}

type Css struct {
	*Basic
	Aling         Position
	Arrangement   Direction
	Margin        float64
	Margin_left   float64
	Margin_Right  float64
	Margin_Top    float64
	Margin_Bottom float64
	Text_size     float64
	Text_Font     font.Face
	cursor_color  color.Color
	Padding       float64
	Resize        bool
	Spacing       float64
	Expand        bool
	Expand_width  bool
	Expand_height bool
	Min           float64
	Display       Display
	Text_Aling    Position
}

func entregarCss(valor *Css, predeterminado *Css) (entregado *Css) {
	if valor != nil {
		entregado = valor
	} else {
		copia := predeterminado
		entregado = copia
	}
	if predeterminado == CssDefaultSlider {
		fmt.Println(entregado.Text_Aling)
	}
	return entregado
}

func entregarColor(valor color.Color, predeterminado color.Color) (entregado color.Color) {
	if valor != nil {
		entregado = valor
	} else {
		entregado = predeterminado
	}
	return entregado
}

var (
	FondoPrincipal      = grafito
	AllCircularEdges    = [4]bool{true, true, true, true}
	LeftCircularEdges   = [4]bool{true, false, true, false}
	RightCircularEdges  = [4]bool{false, true, false, true}
	TopCircularEdges    = [4]bool{true, true, false, false}
	ButtomCircularEdges = [4]bool{false, false, true, true}
	NotCircularEdges    = [4]bool{false, false, false, false}
	CornerButtonLEdge   = [4]bool{false, false, true, false}
	CornerButtonREdge   = [4]bool{false, false, false, true}
	CornerTopLEdge      = [4]bool{true, false, false, false}
	CornerTopREdge      = [4]bool{false, true, false, false}

	Blanco         = color.RGBA{255, 255, 255, 255}
	gris           = color.RGBA{180, 180, 180, 255}
	gris_claro     = color.RGBA{205, 205, 205, 255}
	gris_oscuro    = color.RGBA{135, 135, 135, 255}
	gris_oscuro2   = color.RGBA{120, 120, 120, 255}
	grafito        = color.RGBA{60, 60, 60, 255}
	grafito_oscuro = color.RGBA{30, 30, 30, 255}
	grafito_claro  = color.RGBA{100, 100, 100, 255}
	negro          = color.RGBA{0, 0, 0, 255}
	transparente   = color.RGBA{0, 0, 0, 0}
	naranja        = color.RGBA{255, 127, 0, 255}

	Basic_Rect_Claro = &Basic{
		Border_Width:  1,
		Border_Radius: 5,
		Background:    gris,
		border_Color:  negro,
		Normal_color:  gris_claro,
		Over_color:    Blanco,
		Press_color:   naranja,
		Active_color:  naranja,
		// Text_Aling:       Center,
		Text_color:       negro,
		Text_color_label: negro,
		Bar_color:        gris_oscuro,
	}

	Basic_Rect_ClaroTrans = &Basic{
		Border_Width:  1,
		Border_Radius: 5,
		Background:    gris,
		border_Color:  transparente,
		Normal_color:  transparente,
		Over_color:    transparente,
		Press_color:   naranja,
		Active_color:  naranja,
		// Text_Aling:       Center,
		Text_color:       negro,
		Text_color_label: negro,
		Bar_color:        gris_oscuro,
	}

	Basic_Rect_medio = &Basic{
		Border_Width:     1,
		Border_Radius:    5,
		Background:       gris_oscuro2,
		Foreground:       gris_claro,
		Bar_color:        grafito,
		border_Color:     negro,
		Normal_color:     grafito_claro,
		Over_color:       grafito,
		Press_color:      naranja,
		Active_color:     naranja,
		Text_color:       negro,
		Text_color_label: negro,
	}

	Basic_Rect_Oscuro = &Basic{
		Border_Width:     1,
		Border_Radius:    5,
		border_Color:     negro,
		Background:       grafito,
		Foreground:       grafito_oscuro,
		Bar_color:        naranja,
		Normal_color:     grafito,
		Over_color:       grafito_oscuro,
		Press_color:      grafito_claro,
		Active_color:     naranja,
		Text_color:       Blanco,
		Text_color_label: Blanco,
	}

	CssDefaultPanel = &Css{
		Basic:       Basic_Rect_medio,
		Spacing:     5,
		Margin:      5,
		Padding:     5,
		Resize:      true,
		Expand:      true,
		Margin_left: 40,
	}
	CssDefaultBoton = &Css{
		Basic:      Basic_Rect_Claro,
		Text_Aling: Center,
	}
	CssDefaultBotonTrans = &Css{
		Basic: Basic_Rect_ClaroTrans,
	}
	CssDefaultSpin = &Css{
		Basic:      Basic_Rect_Claro,
		Text_Aling: Center,
	}
	CssDefaultEntry = &Css{
		Basic: Basic_Rect_Oscuro,
		// Text_Aling: Center,
	}
	CssDefaultSlider = &Css{
		Basic:      Basic_Rect_Oscuro,
		Text_Aling: Center,
	}
	CssDefaultMenuItem = &Css{
		Basic: Basic_Rect_Oscuro,
	}
	CssDefaultBoxSize = &Css{
		Margin:  10,
		Spacing: 5,
	}

	CssDefaultLabel = &Css{
		Basic: Basic_Rect_Claro,
	}
	CssDefaultLabelButton = &Css{
		Basic: Basic_Rect_Claro,
		// Text_Aling: Center,
		// Text_color: colornames.White,
	}
	CssDefaultSpaceBar = &Css{
		Basic: Basic_Rect_medio,
	}
	CssDefaultEtiqueta = &Css{
		Basic:        Basic_Rect_Claro,
		Margin:       10,
		Expand_width: true,
	}
	CssDefaultComboBox = &Css{
		Basic: Basic_Rect_Oscuro,
	}
	CssDefaultListBox = &Css{
		Basic:        Basic_Rect_medio,
		Margin_left:  0,
		Margin_Right: 5,
		Margin_Top:   5,
		Spacing:      1,
	}
	CssDefaultCheckButton = &Css{
		Basic:   Basic_Rect_Oscuro,
		Spacing: 5,
	}
	CssDefaultToogleButton = &Css{
		Basic:      Basic_Rect_Claro,
		Aling:      Left,
		Text_Aling: Center,
	}
	CssDefaultItemButtom = &Css{
		Basic: Basic_Rect_Claro,
	}
	CssdefaultArea = &Css{
		Basic:  Basic_Rect_Claro,
		Margin: 5,
		Expand: true,
	}
	CssDefaultIconButton = &Css{
		Basic: Basic_Rect_Oscuro,
	}
)

/*
Entrega la distancia hacia los bordes segun la direccion en que sea necesario, elige
el margen y la distancia especifica del borde y elige la mayor
*/
func distMargen(direccion Position, conf *Css) (valor float64) {
	var margen float64
	switch direccion {
	case Left:
		margen = conf.Margin_left
	case Right:
		margen = conf.Margin_Right
	case Top:
		margen = conf.Margin_Top
	case Buttom:
		margen = conf.Margin_Bottom
	}
	if conf.Margin > margen {
		valor = conf.Margin
	} else {
		valor = margen
	}
	return valor
}
