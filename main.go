package main

import (
	"image/color"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

// Point is a struct for representing 2D vectors.
type Point struct {
	x, y float64
}

type Line struct {
	pos1, pos2 Point
	angle      float64
	color      color.RGBA
}

// NewLine initializes and returns a new Line instance.
func NewLine(x1, y1, x2, y2 float64) *Line {
	return &Line{
		pos1: Point{x: x1, y: y1},
		pos2: Point{x: x2, y: y2},
		color: color.RGBA{
			R: 0xff,
			G: 0xff,
			B: 0xff,
			A: 0xff,
		},
	}
}

func (l *Line) Update() {
	l.angle += 0.01
}

// DrawLineDDA rasterizes a line using Digital Differential Analyzer algorithm.
func (l *Line) DrawLineDDA(img *ebiten.Image, x1, y1, x2, y2 float64, c color.Color) {
	if x2 < x1 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	k := math.Abs((y2 - y1) / (x2 - x1))
	if k <= 1 {
		for x, y := x1, y1+0.5; x <= x2; x, y = x+1, y+k {
			xr := x2 + (x-x2)*math.Cos(l.angle) - (y-y2)*math.Sin(l.angle)
			yr := y2 + (x-x2)*math.Sin(l.angle) + (y-y2)*math.Cos(l.angle)
			img.Set(int(xr), int(yr), c)
		}
	} else {
		k := math.Abs((x2 - x1) / (y2 - y1))
		for x, y := x1+0.5, y1; y <= y2; x, y = x+k, y+1 {
			xr := x2 + (x-x2)*math.Cos(l.angle) - (y-y2)*math.Sin(l.angle)
			yr := y2 + (x-x2)*math.Sin(l.angle) + (y-y2)*math.Cos(l.angle)
			img.Set(int(xr), int(yr), c)
		}
	}
}

func (l *Line) Draw(screen *ebiten.Image) {
	//ebitenutil.DrawLine(screen, 200, 100, 500, 400, l.color)
	l.DrawLineDDA(screen, l.pos1.x, l.pos1.y, l.pos2.x, l.pos2.y, l.color)
}

// Game is a game instance.
type Game struct {
	width, height int
	line          *Line
	// last is a timestamp when Update was called last time.
	last time.Time
}

// NewGame returns a new Game instance.
func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		line:   NewLine(float64(width/2), float64(height/2), 130, 130),
		last:   time.Now(),
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

// Update updates a game state.
func (g *Game) Update() error {
	t := time.Now()
	//dt := float64(t.Sub(g.last).Milliseconds())
	g.last = t
	g.line.Update()
	return nil
}

// Draw renders a game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	g.line.Draw(screen)
}

func main() {
	//rand.Seed(time.Now().UnixNano())
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame(screenWidth, screenHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
