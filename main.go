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
	if math.Abs(p1.x-p0.x) >= math.Abs(p1.y-p0.y) {
		if p0.x > p1.x {
			p0, p1 = p1, p0
		}
		y := p0.y
		for x := p0.x; x <= p1.x; x++ {
			screen.Set(int(x), int(y), color)
			y += (p1.y - p0.y) / (p1.x - p0.x)
		}
	} else {
		if p0.y > p1.y {
			p0, p1 = p1, p0
		}
		x := p0.x
		for y := p0.y; y <= p1.y; y++ {
			screen.Set(int(x), int(y), color)
			x += (p1.x - p0.x) / (p1.y - p0.y)
		}
	}
}

type game struct {
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	DrawLineDDA(screen, Point{0, 0}, Point{screenWidth, screenHeight}, color.RGBA{255, 0, 0, 255})
	DrawLineDDA(screen, Point{screenWidth, 0}, Point{0, screenHeight}, color.RGBA{255, 0, 0, 255})
	DrawLineDDA(screen, Point{screenWidth / 2, 0}, Point{screenWidth / 2, screenHeight}, color.RGBA{255, 0, 0, 255})
	DrawLineDDA(screen, Point{0, screenHeight / 2}, Point{screenWidth, screenHeight / 2}, color.RGBA{255, 0, 0, 255})
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(&game{}); err != nil {
		log.Fatal(err)
	}
}
