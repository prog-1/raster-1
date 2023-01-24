package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Point struct {
	x, y int
}

type Game struct {
	width, height int
	pos           Point
	angle         float64
}

// DrawLineDDA rasterizes a line using Digital Differential Analyzer algorithm.
func DrawLineDDA(img *ebiten.Image, x1, y1, x2, y2 int, c color.Color) {
	if math.Abs(float64(x2-x1)) >= math.Abs(float64(y2-y1)) {
		if x2 < x1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(y2-y1) / float64(x2-x1)
		for x, y := x1, float64(y1)+0.5; x <= x2; x, y = x+1, y+k {
			img.Set(x, int(y), c)
		}
	} else {
		if y2 < y1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := float64(x2-x1) / float64(y2-y1)
		for x, y := float64(x1)+0.5, y1; y <= y2; x, y = x+k, y+1 {
			img.Set(int(x), y, c)
		}
	}
}

func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		pos:    Point{x: 5 * width / 8, y: height / 2},
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.angle += math.Pi / 180
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawLineDDA(screen, g.width/2, g.height/2, int(float64(g.pos.x-g.width/2)*math.Cos(g.angle)-float64(g.pos.y-g.height/2)*math.Sin(g.angle)+float64(g.width/2)), int(float64(g.pos.y-g.height/2)*math.Cos(g.angle)+float64(g.pos.x-g.width/2)*math.Sin(g.angle)+float64(g.height/2)), color.White)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
