package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//---------------------------Declaration--------------------------------

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	width, height int
}

//---------------------------Update-------------------------------------

func (g *Game) Update() error {
	return nil
}

//---------------------------Draw-------------------------------------

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawLineDDA(screen, screenWidth/2, screenHeight/2, screenWidth/2+100, screenHeight/2, color.RGBA{100, 200, 230, 255}) //Line drawing
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
		img.Set(x, int(y), c) //without rotation
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
