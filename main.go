package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Line struct {
	X1, Y1    int
	Magnitude int
	Degrees   float64
	color.Color
}

func ToRadians(Degrees float64) float64 {
	return Degrees * math.Pi / float64(180)
}

type game struct {
	l Line
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	g.l.Degrees += 1
	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	x := float64(g.l.Magnitude) * math.Cos(ToRadians(g.l.Degrees))
	y := float64(g.l.Magnitude) * math.Sin(ToRadians(g.l.Degrees))
	x2, y2 := x+float64(g.l.X1), y+float64(g.l.Y1)

	DrawLineDDA(screen, g.l.X1, g.l.Y1, int(x2), int(y2), g.l.Color)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := game{Line{320, 240, 100, 0, color.RGBA{255, 1, 1, 255}}}
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}

// DrawLineDDA rasterizes a line using Digital Differential Analyzer algorithm.
func DrawLineDDA(img *ebiten.Image, x1, y1, x2, y2 int, c color.Color) {
	if x2 < x1 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	k := float64(y2-y1) / float64(x2-x1)
	if k <= 1 {
		for x, y := x1, float64(y1)+0.5; x <= x2; x, y = x+1, y+k {
			img.Set(x, int(y), c)
		}
	} else {
		k := float64(x2-x1) / float64(y2-y1)
		for x, y := float64(x1), float64(y1)+0.5; y <= float64(y2); x, y = x+k, y+1 {
			img.Set(int(x), int(y), c)
		}
	}
}
