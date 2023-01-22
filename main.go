package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Wwidth  = 700
	Wheight = 700
)

type Game struct {
	width, height int
	angle         float64
}

func (g *Game) Line(img *ebiten.Image, x1, y1, x2, y2 float64) {
	if y2 < y1 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	k := float64(x2-x1) / float64(y2-y1)
	for x, y := float64(x1)+0.5, y1; y <= y2; x, y = x+k, y+1 {
		xr := float64(x1) + float64(x-x1)*math.Cos(g.angle) - (y-float64(y1))*math.Sin(g.angle)
		yr := float64(y1) + float64(x-x1)*math.Sin(g.angle) + (y-float64(y1))*math.Cos(g.angle)
		img.Set(int(xr), int(yr), color.RGBA{210, 39, 48, 255})
	}

}
func (g *Game) Draw(screen *ebiten.Image) {
	//screen.Fill(color.White)
	g.Line(screen, 100, 350, 350, 300)

}

func (g *Game) Update() error {
	g.angle += 0.03
	return nil

}
func (g *Game) Layout(int, int) (w, h int) {
	return g.width, g.height
}

func main() {
	ebiten.SetWindowSize(Wwidth, Wheight)
	if err := ebiten.RunGame(NewGame(Wwidth, Wheight)); err != nil {
		log.Fatal(err)
	}
}
func NewGame(width, height int) *Game {
	return &Game{width: width, height: height}
}
