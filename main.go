package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type (
	point struct {
		x, y float64
	}

	Game struct {
		p1, p2 point
	}
)

var c = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}

const (
	winTitle            = "raster"
	winWidth, winHeight = 500, 500
)

func DrawLineDDA(img *ebiten.Image, x1, y1, x2, y2 float64, c color.Color) {
	if math.Abs(x2-x1) <= math.Abs(y2-y1) {
		if y2 < y1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(x2-x1) / float64(y2-y1)
		for x, y := float64(x1)+0.5, y1; y <= y2; x, y = x+k, y+1 {
			img.Set(int(x), int(y), c)
		}
	} else {
		if x2 < x1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(y2-y1) / float64(x2-x1)
		for x, y := x1, float64(y1)+0.5; x <= x2; x, y = x+1, y+k {
			img.Set(int(x), int(y), c)
		}
	}

}
func (g *Game) Draw(screen *ebiten.Image) {
	DrawLineDDA(screen, g.p1.x, g.p1.y, g.p2.x, g.p2.y, c)

}

func (g *Game) rotation() {
	g.p1.x = g.p1.x*math.Cos(0.03) - g.p1.y*math.Sin(0.03) + (250 - 250*math.Cos(0.03) + 250*math.Sin(0.03))
	g.p1.y = g.p1.x*math.Sin(0.03) + g.p1.y*math.Cos(0.03) + (-250*math.Sin(0.03) + 250 - 250*math.Cos(0.03))
	g.p2.x = g.p2.x*math.Cos(0.03) - g.p2.y*math.Sin(0.03) + (250 - 250*math.Cos(0.03) + 250*math.Sin(0.03))
	g.p2.y = g.p2.x*math.Sin(0.03) + g.p2.y*math.Cos(0.03) + (-250*math.Sin(0.03) + 250 - 250*math.Cos(0.03))
}
func (g *Game) Update() error {
	g.rotation()
	return nil

}

func (g *Game) Layout(int, int) (w, h int) { return winWidth, winHeight }

func main() {
	ebiten.SetWindowTitle(winTitle)
	ebiten.SetWindowSize(winWidth, winHeight)
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(&Game{p1: point{x: 100, y: 100}, p2: point{x: 400, y: 400}}); err != nil {
		log.Fatal(err)
	}
}
