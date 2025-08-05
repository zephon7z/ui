package sui

import (
	"fmt"
	"image"
	"image/color"

	// dw "image/draw"

	// . "math"

	"github.com/faiface/pixel"
	"golang.org/x/image/draw"
)

func init() {
	fmt.Println("iniciando imgrect...")
}

func AnyToInt(val any) (n int) {
	switch v := val.(type) {
	case float64:
		n = int(v)
	case int:
		n = v
	}
	return n
}

func CrearImgRect(xw any, xh any, xr any, colorFondo color.Color, colorBorde color.Color, esq [4]bool) *pixel.PictureData {
	w := AnyToInt(xw)
	h := AnyToInt(xh)
	r := AnyToInt(xr)

	si := esq[0]
	sd := esq[1]
	ii := esq[2]
	id := esq[3]
	imgBorde := image.NewRGBA(image.Rect(0, 0, w, h))
	imgFondo := image.NewRGBA(image.Rect(0, 0, w, h))
	//------------------------------------------------------------------------
	//bordes redondeados
	var ptsInf, ptsSup, a, b, c, d map[int]int
	if sd { //esquina superior derecha
		a = circuloBresenham(w-r-1, r, r, imgBorde, colorBorde, 1)
	}
	if si { //esquina superior izquierda
		b = circuloBresenham(r, r, r, imgBorde, colorBorde, 2)
	}
	if ii { //esquina inferior izquierda
		c = circuloBresenham(r, h-r-1, r, imgBorde, colorBorde, 3)
	}
	if id { //esquina inferior derecha
		d = circuloBresenham(w-r-1, h-r-1, r, imgBorde, colorBorde, 4)
	}
	ptsInf = Unir(a, b)
	ptsSup = Unir(c, d)
	//------------------------------------------------------------------------
	//bordes lineales
	for x := 0; x < w; x++ {
		if _, ok := ptsInf[x]; !ok {
			imgBorde.Set(x, 0, colorBorde)
		}
		if _, ok := ptsSup[x]; !ok {
			imgBorde.Set(x, h-1, colorBorde)
		}
	}
	infI, supI := supInf(0, h, ptsInf, ptsSup)
	infD, supD := supInf(w-1, h, ptsInf, ptsSup)
	for y := 0; y < h; y++ {
		if y > infI && y < supI {
			imgBorde.Set(0, y, colorBorde)
		}
		if y > infD && y < supD {
			imgBorde.Set(w-1, y, colorBorde)
		}
	}
	//------------------------------------------------------------------------
	//relleno del rectangulo
	for x := 0; x < w; x++ {
		inf, sup := supInf(x, h, ptsInf, ptsSup)
		for y := 0; y < h; y++ {
			if y > inf && y < sup {
				imgFondo.Set(x, y, colorFondo)
			}
		}
	}
	//------------------------------------------------------------------------
	//superposicion de imagen
	draw.Draw(imgFondo, imgFondo.Bounds(), imgBorde, image.ZP, draw.Over)
	//borde del rectangulo

	return pixel.PictureDataFromImage(imgFondo)
}

func Unir(d1 map[int]int, d2 map[int]int) (dic map[int]int) {
	if d1 != nil && d2 != nil {
		for k, v := range d2 {
			d1[k] = v
		}
		dic = d1
	} else if d1 != nil {
		dic = d1
	} else if d2 != nil {
		dic = d2
	}
	return dic
}

func supInf(x int, h int, ptsInf map[int]int, ptsSup map[int]int) (int, int) {
	var inf, sup int
	if v, ok := ptsInf[x]; ok {
		inf = v
	} else {
		inf = 0
	}
	if v, ok := ptsSup[x]; ok {
		sup = v
	} else {
		sup = h - 1
	}
	return inf, sup
}

func circuloBresenham(x int, y int, r int, img *image.RGBA, co color.Color, cuadrante int) map[int]int {
	dic := map[int]int{}
	tx := 0
	ty := r
	d := 3 - 2*r
	for tx <= ty {
		switch cuadrante {
		case 1:
			img.Set(x+ty, y-tx, co)
			img.Set(x+tx, y-ty, co)
			dic[x+ty] = y - tx
			dic[x+tx] = y - ty
		case 2:
			img.Set(x-tx, y-ty, co)
			img.Set(x-ty, y-tx, co)
			dic[x-tx] = y - ty
			dic[x-ty] = y - tx
		case 3:
			img.Set(x-ty, y+tx, co)
			img.Set(x-tx, y+ty, co)
			dic[x-ty] = y + tx
			dic[x-tx] = y + ty
		case 4:
			img.Set(x+tx, y+ty, co)
			img.Set(x+ty, y+tx, co)
			dic[x+tx] = y + ty
			dic[x+ty] = y + tx
		}
		if d < 0 {
			d += 4*tx + 6
		} else {
			d += 4*(tx-ty) + 10
			ty--
		}
		tx++
	}
	return dic
}
