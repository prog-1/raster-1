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

type Point struct {
	x, y float64
}

//Implement a program that rotates a vector. The zero point must be in the middle of the window. You must not use any built-in rotation functionality. Just animate an angle and recompute point position using cos/sin functions

func DrawLineDDA(screen *ebiten.Image, p0, p1 Point, color color.Color) {
	dx := math.Abs(p1.x - p0.x)
	dy := math.Abs(p1.y - p0.y)
	if p1.x < p0.x {
		p0, p1 = p1, p0
	}
	if dx >= dy {
		k := dy / dx
		for x, y := p0.x, p0.y+0.5; x <= p1.x; x, y = x+1, y+k {
			screen.Set(int(x), int(y), color)
		}
	} else {
		k := dx / dy
		for x, y := p0.x+0.5, p0.y; y <= p1.y; x, y = x+k, y+1 {
			screen.Set(int(x), int(y), color)
		}
	}

}

type game struct {
	angle  float64
	p0, p1 Point
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	g.angle += math.Pi / 180
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.p0.x = screenWidth / 2
	g.p0.y = screenHeight / 2
	g.p1.x = (math.Cos(g.angle) * (-math.Sin(g.angle))) * g.p0.x
	g.p1.y = math.Sin(g.angle) * math.Cos(g.angle) * g.p0.y
	DrawLineDDA(screen, g.p0, g.p1, color.RGBA{255, 0, 0, 255})
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(&game{}); err != nil {
		log.Fatal(err)
	}
}
