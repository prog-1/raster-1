package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

//---------------------------Declaration--------------------------------

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	//all global variables
	width, height int
	angle         float64
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	g.angle += 0.01 //increasing line angle
	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawLineDDA(screen, screenWidth/2, screenHeight/2, screenWidth/2+150, screenHeight/2, color.RGBA{100, 200, 230, 255}) //Drawing line
}

//-------------------------Draw-Functions-----------------------------

//Line Drawing function
func (g *Game) DrawLineDDA(img *ebiten.Image, x1, y1, x2, y2 int, c color.Color) {
	if x2 < x1 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	k := float64(y2-y1) / float64(x2-x1)
	for x, y := x1, float64(y1)+0.5; x <= x2; x, y = x+1, y+k {
		//-----------------------------------------------------
		// Coordinate rotation logic
		xRot := float64(x1) + float64(x-x1)*math.Cos(g.angle) - (y-float64(y1))*math.Sin(g.angle)
		yRot := float64(y1) + float64(x-x1)*math.Sin(g.angle) + (y-float64(y1))*math.Cos(g.angle)
		//-----------------------------------------------------
		img.Set(int(xRot), int(yRot), c) //with rotation
		//img.Set(x, int(y), c) //without rotation
	}
}

//---------------------------Main-------------------------------------

func (g *Game) Layout(inWidth, inHeight int) (outWidth, outHeight int) {
	return g.width, g.height
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Line")
	g := NewGame(screenWidth, screenHeight)   //creating game instance
	if err := ebiten.RunGame(g); err != nil { //running game
		log.Fatal(err)
	}
}

//New game instance creation
func NewGame(width, height int) *Game {
	return &Game{width: width, height: height}
}
