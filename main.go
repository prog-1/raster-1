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

// Point is a struct for representing 2D vectors.
type Point struct {
	x, y float64
}

type Line struct {
	pos1, pos2     Point
	radians        float64
	startx, starty float64
	color          color.RGBA
}

// These 2 functions are needed so that the initial length of the line always remains the same

func StartingXPos(x1, y1, x2, y2, radians float64) float64 {
	if x1 == x2 && y1 != y2 { // if the line is vertical
		return StartingYPos(x1, y1, x1, y2, radians)
	}
	return (x2 - x1) / math.Cos(radians)
}

func StartingYPos(x1, y1, x2, y2, radians float64) float64 {
	if y1 == y2 && x1 != x2 { // if the line is horizontal
		return StartingXPos(x1, y1, x2, y2, radians)
	}
	return (y2 - y1) / math.Sin(radians)
}

// NewLine initializes and returns a new Line instance.
func NewLine(x1, y1, x2, y2 float64) *Line {
	return &Line{
		pos1:    Point{x: x1, y: y1},
		pos2:    Point{x: x2, y: y2},
		radians: math.Atan((y2 - y1) / (x2 - x1)), // math.Atan(k)  k выражено из уравнения прямой, проходящей через две точки
		startx:  StartingXPos(x1, y1, x2, y2, math.Atan((y2-y1)/(x2-x1))),
		starty:  StartingYPos(x1, y1, x2, y2, math.Atan((y2-y1)/(x2-x1))),
		color: color.RGBA{
			R: 0xff,
			G: 0xff,
			B: 0xff,
			A: 0xff,
		},
	}
}

// DrawLineDDA rasterizes a line using Digital Differential Analyzer algorithm.
func DrawLineDDA(img *ebiten.Image, x1, y1, x2, y2 float64, c color.Color) {
	if math.Abs(y2-y1) <= math.Abs(x2-x1) {
		if x2 < x1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := (y2 - y1) / (x2 - x1)
		for x, y := x1, y1+0.5; x <= x2; x, y = x+1, y+k {
			img.Set(int(x), int(y), c)
		}
	} else {
		if y2 < y1 {
			x1, x2 = x2, x1
			y1, y2 = y2, y1
		}
		k := (x2 - x1) / (y2 - y1)
		for x, y := x1+0.5, y1; y <= y2; x, y = x+k, y+1 {
			img.Set(int(x), int(y), c)
		}
	}
}

func (l *Line) Update() {
	x := math.Cos(l.radians) * l.startx
	y := math.Sin(l.radians) * l.starty
	l.pos2.x = l.pos1.x + x
	l.pos2.y = l.pos1.y + y
	l.radians += math.Pi / 180
}

func (l *Line) Draw(screen *ebiten.Image) {
	DrawLineDDA(screen, l.pos1.x, l.pos1.y, l.pos2.x, l.pos2.y, l.color)
}

// Game is a game instance.
type Game struct {
	width, height int
	line          *Line
}

// NewGame returns a new Game instance.
func NewGame(width, height int) *Game {
	return &Game{
		width:  width,
		height: height,
		line:   NewLine(float64(width/2), float64(height/2), 130, 130),
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

// Update updates a game state.
func (g *Game) Update() error {
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
