package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Point struct {
	x, y float64
}

const (
	screenWidth  = 640
	screenHeight = 480
)

func DrawLineDDA(img *ebiten.Image, x1, y1, x2, y2 float64, c color.Color) {
	if math.Abs(x2-x1) >= math.Abs(y2-y1) {
		if x2 < x1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(y2-y1) / float64(x2-x1)
		for x, y := x1, float64(y1)+0.5; x <= x2; x, y = x+1, y+k {
			img.Set(int(x), int(y), c)
		}
	} else {
		if y2 < y1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(x2-x1) / float64(y2-y1)
		for x, y := float64(x1)+0.5, y1; y <= y2; x, y = x+k, y+1 {
			img.Set(int(x), int(y), c)
		}
	}
}

type Game struct {
	d          float64
	pos1, pos2 Point
}

func (g *Game) Update() error {
	g.d++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ang := math.Pi * g.d / 180
	g.pos2.x -= g.pos1.x
	g.pos2.y -= g.pos1.y
	g.pos2.x = g.pos2.x*math.Cos(ang) + g.pos2.y*math.Sin(ang)
	g.pos2.y = -g.pos2.x*math.Sin(ang) + g.pos2.y*math.Cos(ang)
	g.pos2.x += g.pos1.x
	g.pos2.y += g.pos1.y
	DrawLineDDA(screen, g.pos1.x, g.pos1.y, g.pos2.x, g.pos2.y, color.RGBA{R: 227, G: 76, B: 235, A: 1})
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := ebiten.RunGame(&Game{pos1: Point{x: 200, y: 200}, pos2: Point{x: 400, y: 400}}); err != nil {
		log.Fatal(err)
	}
}
