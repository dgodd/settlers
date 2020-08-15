package main

import (
	"fmt"
	"image/color"
	"log"
	//"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
)

func init() {
	emptyImage.Fill(color.White)
}

type Game struct {
	count int
}

func hexagon(x, y float32, clr color.RGBA) ([]ebiten.Vertex, []uint16) {
	r := float32(clr.R) / 0xff
	g := float32(clr.G) / 0xff
	b := float32(clr.B) / 0xff
	a := float32(clr.A) / 0xff

	return []ebiten.Vertex{
		{ DstX:   x, DstY:   y-50, SrcX:   1, SrcY:   1, ColorR: r, ColorG: g, ColorB: b, ColorA: a },
		{ DstX:   x+50, DstY:   y-25, SrcX:   1, SrcY:   1, ColorR: r, ColorG: g, ColorB: b, ColorA: a },
		{ DstX:   x+50, DstY:   y+25, SrcX:   1, SrcY:   1, ColorR: r, ColorG: g, ColorB: b, ColorA: a },
		{ DstX:   x, DstY:   y+50, SrcX:   1, SrcY:   1, ColorR: r, ColorG: g, ColorB: b, ColorA: a },
		{ DstX:   x-50, DstY:   y+25, SrcX:   1, SrcY:   1, ColorR: r, ColorG: g, ColorB: b, ColorA: a },
		{ DstX:   x-50, DstY:   y-25, SrcX:   1, SrcY:   1, ColorR: r, ColorG: g, ColorB: b, ColorA: a },
	}, []uint16{
		0, 1, 2,
		1, 2, 3,
		2, 3, 4,
		3, 4, 5,
		4, 5, 0,
		5, 0, 1,
		0, 2, 3,
		3, 5, 0,
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.count++
	g.count %= 240
	return nil
}

func drawWood(screen *ebiten.Image, x, y float32) {
	v, i := hexagon(x,y, color.RGBA{0x16, 0x97, 0x38, 0xff})
	screen.DrawTriangles(v, i, emptyImage, nil)
}

func drawWheat(screen *ebiten.Image, x, y float32) {
	v, i := hexagon(x,y, color.RGBA{0xF3, 0xBA, 0x21, 0xff})
	screen.DrawTriangles(v, i, emptyImage, nil)
}

func drawBrick(screen *ebiten.Image, x, y float32) {
	v, i := hexagon(x,y, color.RGBA{0xE2, 0x64, 0x29, 0xff})
	screen.DrawTriangles(v, i, emptyImage, nil)
}

func drawSheep(screen *ebiten.Image, x, y float32) {
	v, i := hexagon(x,y, color.RGBA{0x90, 0xB6, 0x0B, 0xff})
	screen.DrawTriangles(v, i, emptyImage, nil)
}

func drawOre(screen *ebiten.Image, x, y float32) {
	v, i := hexagon(x,y, color.RGBA{0xA2, 0xA8, 0xA4, 0xff})
	screen.DrawTriangles(v, i, emptyImage, nil)
}

func drawWater(screen *ebiten.Image, x, y float32) {
	v, i := hexagon(x,y, color.RGBA{0x08, 0x66, 0xA5, 0xff})
	screen.DrawTriangles(v, i, emptyImage, nil)
}

func (g *Game) Draw(screen *ebiten.Image) {
	// cf := float64(g.count)

	drawWood(screen, 200, 200)
	drawWheat(screen, 300, 200)
	drawBrick(screen, 400, 200)
	drawSheep(screen, 250, 275)
	drawOre(screen, 350, 275)
	drawWater(screen, 450, 275)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Settlers of Catan")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
