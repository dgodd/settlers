package main

import (
	"fmt"
	"github.com/dgodd/settlers/board"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

var (
	emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
	mplusFont     font.Face
)

func init() {
	rand.Seed(time.Now().UnixNano())

	emptyImage.Fill(color.White)

	tt, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	mplusFont = truetype.NewFace(tt, &truetype.Options{
		Size:    20,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}

var Colors = []color.RGBA{
	{0x0, 0x0, 0x0, 0x00},    // Nothing
	{0x08, 0x66, 0xA5, 0xFF}, // Water
	{0xD6, 0xCE, 0x90, 0xFF}, // Desert
	{0x14, 0x95, 0x3A, 0xFF}, // Wood
	{0xE2, 0x64, 0x29, 0xFF}, // Brick
	{0x90, 0xB6, 0x0B, 0xff}, // Sheep,
	{0xF3, 0xBA, 0x21, 0xff}, // Wheat,
	{0xA2, 0xA8, 0xA4, 0xff}, // Ore,
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
			{DstX: x, DstY: y - 48, SrcX: 1, SrcY: 1, ColorR: r, ColorG: g, ColorB: b, ColorA: a},
			{DstX: x + 48, DstY: y - 24, SrcX: 1, SrcY: 1, ColorR: r, ColorG: g, ColorB: b, ColorA: a},
			{DstX: x + 48, DstY: y + 24, SrcX: 1, SrcY: 1, ColorR: r, ColorG: g, ColorB: b, ColorA: a},
			{DstX: x, DstY: y + 48, SrcX: 1, SrcY: 1, ColorR: r, ColorG: g, ColorB: b, ColorA: a},
			{DstX: x - 48, DstY: y + 24, SrcX: 1, SrcY: 1, ColorR: r, ColorG: g, ColorB: b, ColorA: a},
			{DstX: x - 48, DstY: y - 24, SrcX: 1, SrcY: 1, ColorR: r, ColorG: g, ColorB: b, ColorA: a},
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
			x := float32(idxX)*100.0 + 50.0
			y := float32(idxY)*75.0 + 50.0
			if idxY%2 == 1 {
				x += 50
			}
			v, i := hexagon(x, y, Colors[tile.Klass])
			screen.DrawTriangles(v, i, emptyImage, nil)

			if tile.Number > 0 {
				drawNumber(screen, float64(x-15), float64(y), tile.Number)
			}
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func drawNumber(screen *ebiten.Image, x, y float64, number int) {
	common := []int{0, 0, 1, 2, 3, 4, 5, 0, 5, 4, 3, 2, 1}
	b, _ := font.BoundString(mplusFont, strconv.Itoa(number))
	w := (b.Max.X - b.Min.X).Ceil()
	h := (b.Max.Y - b.Min.Y).Ceil()
	gray := color.RGBA{0xEA, 0xE9, 0xE4, 0xFF}
	ebitenutil.DrawRect(screen, x, y, 30, 30, gray)
	textColor := color.Color(color.Black)
	if common[number] == 5 {
		textColor = color.RGBA{0xFF, 0x0, 0x0, 0xFF}
	}
	text.Draw(screen, strconv.Itoa(number), mplusFont, int(x+14.0-(float64(w)/2.0)), int(y)+h+4, textColor)

	dots := ""
	for i := 0; i < common[number]; i++ {
		dots += "."
	}
	b, _ = font.BoundString(mplusFont, dots)
	w = (b.Max.X - b.Min.X).Ceil()
	text.Draw(screen, dots, mplusFont, int(x+14.0-(float64(w)/2.0)), int(y)+h+10, textColor)
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
