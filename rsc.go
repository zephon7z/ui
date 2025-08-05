package sui

import (
	"fmt"

	"image"
	_ "image/png"
	"os"

	"io/ioutil"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func CheckErrorPanic(mensaje string, err error) {
	if err != nil {
		fmt.Println(mensaje)
		panic(err)
	}
}

func CargarImagen(ruta string) pixel.Picture {
	archivo, err := os.Open(ruta)
	defer archivo.Close()
	CheckErrorPanic("no se pudo abrir el archivo de imagen:"+ruta, err)
	imagen, _, err := image.Decode(archivo)
	CheckErrorPanic("no se pudo decodificar el archivo de imagen: "+ruta, err)
	return pixel.PictureDataFromImage(imagen)
}

func CargarFuente(ruta string, alto float64) font.Face {
	archivo, err := os.Open(ruta)
	CheckErrorPanic("no se pudo abrir el archivo de fuente ttf", err)
	bytes, err := ioutil.ReadAll(archivo)
	CheckErrorPanic("no se pudo decodificar la fuente", err)
	fuente, err := truetype.Parse(bytes)
	CheckErrorPanic("no se pudo analizar la fuente", err)
	return truetype.NewFace(fuente, &truetype.Options{
		Size:              alto,
		GlyphCacheEntries: 1,
	})
}

var (
	imgWidgets pixel.Picture
	atlas      map[int]*text.Atlas
	txt        map[int]*text.Text
	diseno     map[string]pixel.Rect
)

type Class struct {
	Icon       pixel.Rect
	Normal     pixel.Rect
	Sobre      pixel.Rect
	Presionado pixel.Rect
	Activo     pixel.Rect
	Apagado    pixel.Rect
}

type Conf struct {
	Fuente     string
	Tamanos    []int
	Img        string
	BtnNormal  pixel.Rect
	BtnSobre   pixel.Rect
	BtnPress   pixel.Rect
	BtnActivo  pixel.Rect
	BtnApagado pixel.Rect
}

func Iniciar(conf Conf) {
	atlas = map[int]*text.Atlas{
		10: text.NewAtlas(CargarFuente("report.ttf", 15), text.ASCII),
	}
	txt = map[int]*text.Text{
		10: text.New(pixel.ZV, atlas[10]),
	}
	diseno = map[string]pixel.Rect{
		"btn normal":     conf.BtnNormal,
		"btn sobre":      conf.BtnSobre,
		"btn presionado": conf.BtnPress,
		"btn activo":     conf.BtnActivo,
		"btn Apagado":    conf.BtnApagado,
	}
	imgWidgets = CargarImagen(conf.Img)

}
