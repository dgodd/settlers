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
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

var (
	emptyImage, _                                                                               = ebiten.NewImage(16, 16, ebiten.FilterDefault)
	woodImage, brickImage, sheepImage, wheatImage, oreImage, houseImage, roadImage, circleImage *ebiten.Image
	mplusFont                                                                                   font.Face
	houseXY, roadXY                                                                             board.XY
)

func init() {
	rand.Seed(time.Now().UnixNano())

	houseXY = board.XY{X: 28, Y: screenHeight - 28.0}
	roadXY = board.XY{X: 68, Y: screenHeight - 28.0}

	emptyImage.Fill(color.White)
	{
		f, err := os.Open("images/wood.png")
		if err != nil {
			log.Fatal(err)
		}
		img, _, err := image.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		woodImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		f.Close()
	}
	{
		f, err := os.Open("images/brick.png")
		if err != nil {
			log.Fatal(err)
		}
		img, _, err := image.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		brickImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		f.Close()
	}
	{
		f, err := os.Open("images/sheep.png")
		if err != nil {
			log.Fatal(err)
		}
		img, _, err := image.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		sheepImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		f.Close()
	}
	{
		f, err := os.Open("images/wheat.png")
		if err != nil {
			log.Fatal(err)
		}
		img, _, err := image.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		wheatImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		f.Close()
	}
	{
		f, err := os.Open("images/ore.png")
		if err != nil {
			log.Fatal(err)
		}
		img, _, err := image.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		oreImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		f.Close()
	}
	{
		f, err := os.Open("images/house.png")
		if err != nil {
			log.Fatal(err)
		}
		img, _, err := image.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		houseImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		f.Close()
	}
	{
		f, err := os.Open("images/road.png")
		if err != nil {
			log.Fatal(err)
		}
		img, _, err := image.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		roadImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		f.Close()
	}
	{
		f, err := os.Open("images/circle.png")
		if err != nil {
			log.Fatal(err)
		}
		img, _, err := image.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		circleImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		f.Close()
	}

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
	Board            board.Board
	LeftMouseClicked bool
	MouseXY          board.XY
	PlaceHouse       bool
	PlaceRoad        bool
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
	x, y := ebiten.CursorPosition()
	g.MouseXY.X, g.MouseXY.Y = float64(x), float64(y)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.LeftMouseClicked = true
	} else if g.LeftMouseClicked {
		g.LeftMouseClicked = false
		fmt.Println("You pressed the 'LEFT' mouse button.", g.MouseXY)
		if corner := g.Board.FindCorner(g.MouseXY); corner != nil {
			fmt.Println("You selected:", corner)
			town := g.Board.Towns[*corner]
			if town.Owner == 0 || town.Owner == 1 { // TODO: Which owner am I? Hardcode "1" here.
				town.Owner = 1
				if town.Type < 2 {
					town.Type += 1
				}
				g.Board.Towns[*corner] = town
				fmt.Println("You selected:", town)
			} else {
				// TODO: In game alert
				fmt.Println("ERROR: Can't steal someones town")
			}
		} else if g.MouseXY.Distance(&houseXY) <= 11.0 {
			fmt.Println("CLICKED ON HOUSE ICON")
			g.PlaceHouse = !g.PlaceHouse
			g.PlaceRoad = false
		} else if g.MouseXY.Distance(&roadXY) <= 11.0 {
			fmt.Println("CLICKED ON ROAD ICON")
			g.PlaceHouse = false
			g.PlaceRoad = !g.PlaceRoad
		}
	}
	// TODO: DO STUFF
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// cf := float64(g.count)

	screen.Fill(color.RGBA{0x77, 0x77, 0x77, 0xff})

	for idxY, row := range g.Board.Tiles {
		for idxX, tile := range row {
			x := float32(idxX)*100.0 + 50.0
			y := float32(idxY)*75.0 + 50.0
			if idxY%2 == 1 {
				x += 50
			}
			v, i := hexagon(x, y, Colors[tile.Klass])
			screen.DrawTriangles(v, i, emptyImage, nil)

			if tile.Klass == board.Wood {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(0.3, 0.3)
				op.GeoM.Translate(float64(x)-30.0, float64(y)-50.0)
				screen.DrawImage(woodImage, op)
			} else if tile.Klass == board.Brick {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(0.22, 0.22)
				op.GeoM.Translate(float64(x)-20.0, float64(y)-46.0)
				screen.DrawImage(brickImage, op)
			} else if tile.Klass == board.Sheep {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(0.2, 0.2)
				op.GeoM.Translate(float64(x)-20.0, float64(y)-44.0)
				screen.DrawImage(sheepImage, op)
			} else if tile.Klass == board.Wheat {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(0.3, 0.3)
				op.GeoM.Rotate(0.9)
				op.GeoM.Translate(float64(x)+4.0, float64(y)-50.0)
				screen.DrawImage(wheatImage, op)
			} else if tile.Klass == board.Ore {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(0.016, 0.016)
				op.GeoM.Translate(float64(x)-10.0, float64(y)-40.0)
				screen.DrawImage(oreImage, op)
			}

			if tile.Number > 0 {
				drawNumber(screen, float64(x-15), float64(y), tile.Number)
			}
		}
	}

	for _, xy := range g.Board.Corners {
		if town, ok := g.Board.Towns[xy]; ok {
			if town.Type <= 2 { // TODO: Distinguish between settlement and city
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Scale(0.34, 0.34)
				op.GeoM.Translate(xy.X-18.0, xy.Y-18.0)
				op.ColorM = ebiten.ScaleColor(1, 0, 0, 1)
				screen.DrawImage(houseImage, op)
			}
		} else {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(0.2, 0.2)
			op.ColorM = ebiten.ScaleColor(1, 1, 0, 1)
			if g.MouseXY.Distance(&xy) <= 11.0 {
				op.ColorM = ebiten.ScaleColor(0, 0, 1, 1)
			}
			op.GeoM.Translate(xy.X-11.0, xy.Y-11.0)
			screen.DrawImage(circleImage, op)
		}
	}

	{
		// DRAW HOUSE TO CHOOSE TO BUY ONE
		xy := houseXY
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.34, 0.34)
		op.GeoM.Translate(xy.X-18.0, xy.Y-18.0)
		if g.PlaceHouse {
			op.ColorM = ebiten.ScaleColor(1, 0, 0, 1)
		} else if g.MouseXY.Distance(&xy) <= 18.0 {
			op.ColorM = ebiten.ScaleColor(1, 0, 0, 0.5)
		} else {
			op.ColorM = ebiten.ScaleColor(1, 0, 0, 0.3)
		}
		screen.DrawImage(houseImage, op)
	}

	{
		// DRAW THE ROAD
		xy := roadXY
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.34, 0.34)
		op.GeoM.Translate(xy.X-18.0, xy.Y-18.0)
		if g.PlaceRoad {
			op.ColorM = ebiten.ScaleColor(1, 0, 0, 1)
		} else if g.MouseXY.Distance(&xy) <= 18.0 {
			op.ColorM = ebiten.ScaleColor(1, 0, 0, 0.5)
		} else {
			op.ColorM = ebiten.ScaleColor(1, 0, 0, 0.3)
		}
		screen.DrawImage(roadImage, op)
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
