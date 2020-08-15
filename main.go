package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"
	"github.com/dgodd/settlers/board"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

var (
	emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
)

func init() {
	rand.Seed(time.Now().UnixNano())
	emptyImage.Fill(color.White)
}

var Colors = []color.RGBA{
	color.RGBA{0x0, 0x0, 0x0, 0x00}, // Nothing
	color.RGBA{0x08, 0x66, 0xA5, 0xFF}, // Water
	color.RGBA{0xD6, 0xCE, 0x90, 0xFF}, // Desert
	color.RGBA{0x14, 0x95, 0x3A, 0xFF}, // Wood
	color.RGBA{0xE2, 0x64, 0x29, 0xFF}, // Brick
	color.RGBA{0x90, 0xB6, 0x0B, 0xff}, // Sheep,
	color.RGBA{0xF3, 0xBA, 0x21, 0xff}, // Wheat,
	color.RGBA{0xA2, 0xA8, 0xA4, 0xff}, // Ore,
}

type Game struct {
	Board board.Board
}

func hexagon(x, y float32, clr color.RGBA) ([]ebiten.Vertex, []uint16) {
	r := float32(clr.R) / 0xff
	g := float32(clr.G) / 0xff
	b := float32(clr.B) / 0xff
	a := float32(clr.A) / 0xff

	return []ebiten.Vertex{
		{ DstX:   x, DstY:   y-48, SrcX:   1, SrcY:   1, ColorR: r, ColorG: g, ColorB: b, ColorA: a },
		{ DstX:   x+48, DstY:   y-24, SrcX:   1, SrcY:   1, ColorR: r, ColorG: g, ColorB: b, ColorA: a },
		{ DstX:   x+48, DstY:   y+24, SrcX:   1, SrcY:   1, ColorR: r, ColorG: g, ColorB: b, ColorA: a },
		{ DstX:   x, DstY:   y+48, SrcX:   1, SrcY:   1, ColorR: r, ColorG: g, ColorB: b, ColorA: a },
		{ DstX:   x-48, DstY:   y+24, SrcX:   1, SrcY:   1, ColorR: r, ColorG: g, ColorB: b, ColorA: a },
		{ DstX:   x-48, DstY:   y-24, SrcX:   1, SrcY:   1, ColorR: r, ColorG: g, ColorB: b, ColorA: a },
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
	// TODO: DO STUFF
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// cf := float64(g.count)

	for idxY, row := range g.Board.Tiles {
		for idxX, tile := range row {
			x := float32(idxX) * 100.0 + 50.0
			y := float32(idxY) * 75.0 + 50.0
			if (idxY % 2 == 1) {
				x += 50
			}
			v, i := hexagon(x, y, Colors[tile.Klass])
			screen.DrawTriangles(v, i, emptyImage, nil)
		} 
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Settlers of Catan")

	game := &Game{
		Board: board.NewSimple(),
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
