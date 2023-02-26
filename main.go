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

func DrawLineDDA(screen *ebiten.Image, p0x, p0y, p1x, p1y float64, color color.Color) {
	if math.Abs(p1x-p0x) >= math.Abs(p1y-p0y) {
		if p0x > p1x {
			p0x, p1x = p1x, p0x
			p0y, p1y = p1y, p0y
		}
		y := p0y
		for x := p0x; x <= p1x; x++ {
			screen.Set(int(x), int(y), color)
			y += (p1y - p0y) / (p1x - p0x)
		}
	} else {
		if p0y > p1y {
			p0x, p1x = p1x, p0x
			p0y, p1y = p1y, p0y
		}
		x := p0x
		for y := p0y; y <= p1y; y++ {
			screen.Set(int(x), int(y), color)
			x += (p1x - p0x) / (p1y - p0y)
		}
	}
}

type game struct {
	angle  float64
	p0, p1 Point
}

func RotateLine(p0, p1 Point, angle float64) Point {
	var c Point
	c.x = ((p1.x-p0.x)*math.Cos(angle) - (p1.y-p0.y)*math.Sin(angle))
	c.y = ((p1.x-p0.x)*math.Sin(angle) + (p1.y-p0.y)*math.Cos(angle))
	p1.x, p1.y = c.x, c.y
	return p1
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {
	g.angle += math.Pi / 180
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.p0 = Point{screenWidth / 2, screenHeight / 2}
	g.p1 = RotateLine(g.p0, Point{0, 0}, g.angle)
	DrawLineDDA(screen, 0, 0, screenWidth, screenHeight, color.RGBA{255, 0, 0, 255})
	DrawLineDDA(screen, screenWidth, 0, 0, screenHeight, color.RGBA{255, 0, 0, 255})
	DrawLineDDA(screen, screenWidth/2, 0, screenWidth/2, screenHeight, color.RGBA{255, 0, 0, 255})
	DrawLineDDA(screen, 0, screenHeight/2, screenWidth, screenHeight/2, color.RGBA{255, 0, 0, 255})
	DrawLineDDA(screen, g.p0.x, g.p0.y, g.p0.x+g.p1.x, g.p0.y+g.p1.y, color.RGBA{0, 255, 0, 255})
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(&game{}); err != nil {
		log.Fatal(err)
	}
}
