package sui

import (
	"fmt"
	. "math"

	// "github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

var (
	ctx *imdraw.IMDraw
)

func init() {
	fmt.Println("iniciando colision...")
	ctx = imdraw.New(nil)
}

// estructura de un punto
type P struct {
	X float64
	Y float64
}

// estructura de un circulo
type C struct {
	*P
	R float64
}

// crea un nuevo circulo
func NewC(x float64, y float64, r float64) *C {
	return &C{&P{x, y}, r}
}

// estructura de un rectangulo
type R struct {
	*P
	W float64
	H float64
}

// crea un nuevo rectangulo
func NewR(x float64, y float64, w float64, h float64) *R {
	return &R{&P{x, y}, w, h}
}

func (r *R) CollideP(p *P) bool {
	if (p.X > r.X && p.X < r.X+r.W) && (p.Y > r.Y && p.Y < r.Y+r.H) {
		return true
	} else {
		return false
	}
}

func (r *R) CollideC(c *C) bool {
	if (c.X+c.R > r.X && c.X-c.R < r.X+r.W) && (c.Y+c.R > r.Y && c.Y-c.R < r.Y+r.H) {
		return true
	} else {
		dx := Abs(c.X-(r.X+r.W/2)) - r.W
		dy := Abs(c.Y-(r.Y+r.H/2)) - r.H
		dist := Sqrt(Pow(dx, 2) + Pow(dy, 2))
		if dist < c.R {
			return true
		} else {
			return false
		}
	}
}

func (r1 *R) CollideR(r2 *R) bool {
	if (r1.X+r1.W > r2.X && r2.X+r2.W > r1.X) && (r1.Y+r1.H > r2.Y && r2.Y+r2.H > r1.Y) {
		return true
	} else {
		return false
	}
}

func (c *C) CollideP(p *P) bool {
	dist := Sqrt(Pow((c.X-p.X), 2) + Pow(c.Y-p.Y, 2))
	if dist < c.R {
		return true
	} else {
		return false
	}
}

func (c1 *C) CollideC(c2 *C) bool {
	dist := Sqrt(Pow(c1.X-c2.X, 2) + Pow(c1.Y-c2.Y, 2))
	if dist < c1.R+c2.R {
		return true
	} else {
		return false
	}
}

func (c *C) CollideR(r *R) bool {
	if (c.X+c.R > r.X && c.X-c.R < r.X+r.W) && (c.Y+c.R > r.Y && c.Y-c.R < r.Y+r.H) {
		return true
	} else {
		dx := Abs(c.X-(r.X+r.W/2)) - r.W
		dy := Abs(c.Y-(r.Y+r.H/2)) - r.H
		dist := Sqrt(Pow(dx, 2) + Pow(dy, 2))
		if dist < c.R {
			return true
		} else {
			return false
		}
	}
}

func (r *R) DrawFig(ctx imdraw.IMDraw) {
	// ctx.Color = pixel.Rect
}
