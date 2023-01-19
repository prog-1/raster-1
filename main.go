package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type game struct{}

func (*game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (*game) Update() error                             { return nil }
func (*game) Draw(screen *ebiten.Image) {
	DrawLineDDA(screen, 320, 240, 300, 50, color.RGBA{255, 1, 1, 255})
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(&game{}); err != nil {
		log.Fatal(err)
	}
}

// DrawLineDDA rasterizes a line using Digital Differential Analyzer algorithm.
func DrawLineDDA(img *ebiten.Image, x1, y1, x2, y2 int, c color.Color) {
	if x2 < x1 {
		x1, x2 = x2, x1
		y1, y2 = y2, y1
	}
	k := float64(y2-y1) / float64(x2-x1)
	if k <= 1 {
		for x, y := x1, float64(y1)+0.5; x <= x2; x, y = x+1, y+k {
			img.Set(x, int(y), c)
		}
	} else {
		k := float64(x2-x1) / float64(y2-y1)
		for x, y := float64(x1), float64(y1)+0.5; y <= float64(y2); x, y = x+k, y+1 {
			img.Set(int(x), int(y), c)
		}
	}
}
