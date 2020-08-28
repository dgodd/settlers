package board

import (
	"math"
	"math/rand"
)

const (
	Nothing = iota
	Water
	Desert
	Wood
	Brick
	Sheep
	Wheat
	Ore
)

type Tile struct {
	Klass     int
	Number    int
	HasRobber bool
}

type Board struct {
	Tiles   [][]Tile
	Corners []XY
}

func New(tilePlaces [][]byte, numKlass []int, numbers []int) Board {
	var tiles [][]Tile
	numKlassIdx := 0
	numbersIdx := 0
	for _, r := range tilePlaces {
		var row []Tile
		for _, c := range r {
			if c == 'l' {
				tile := Tile{Klass: numKlass[numKlassIdx]}
				numKlassIdx += 1
				if tile.Klass != Desert {
					tile.Number = numbers[numbersIdx]
					numbersIdx += 1
				}
				row = append(row, tile)
			} else if c == 'w' {
				row = append(row, Tile{Klass: Water})
			} else {
				row = append(row, Tile{Klass: Nothing})
			}
		}
		tiles = append(tiles, row)
	}
	return Board{
		Tiles:   tiles,
		Corners: corners(tiles),
	}
}

func NewSimple() Board {
	tilePlaces := [][]byte{
		{'n', 'n', 'w', 'w', 'w', 'w', 'n'},
		{'n', 'w', 'l', 'l', 'l', 'w', 'n'},
		{'n', 'w', 'l', 'l', 'l', 'l', 'w'},
		{'w', 'l', 'l', 'l', 'l', 'l', 'w'},
		{'n', 'w', 'l', 'l', 'l', 'l', 'w'},
		{'n', 'w', 'l', 'l', 'l', 'w', 'n'},
		{'n', 'n', 'w', 'w', 'w', 'w', 'n'},
	}
	numKlass := []int{
		Desert,
		Wood, Wood, Wood, Wood,
		Brick, Brick, Brick,
		Sheep, Sheep, Sheep, Sheep,
		Wheat, Wheat, Wheat, Wheat,
		Ore, Ore, Ore,
	}
	shuffle(numKlass)
	numbers := []int{
		2, 3, 5, 6, 8, 10, 10, 11, 11, 12, 3, 4, 4, 5, 6, 8, 9, 9,
	}
	shuffle(numbers)

	return New(tilePlaces, numKlass, numbers)
}

func shuffle(a []int) {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}

type XY struct {
	X float64
	Y float64
}

func (xy *XY) Distance(other *XY) float64 {
	x := xy.X - other.X
	y := xy.Y - other.Y
	return math.Sqrt(x*x + y*y)
}

func corners(tiles [][]Tile) []XY {
	set := make(map[XY]bool, 0)
	for idxY, row := range tiles {
		for idxX, tile := range row {
			if tile.Number > 0 {
				x := float32(idxX)*100.0 + 50.0
				y := float32(idxY)*75.0 + 50.0
				if idxY%2 == 1 {
					x += 50
				}
				set[XY{float64(x), float64(y) - 50.0}] = true
				set[XY{float64(x), float64(y) + 50.0}] = true
				set[XY{float64(x) + 50.0, float64(y) - 25.0}] = true
				set[XY{float64(x) + 50.0, float64(y) + 25.0}] = true
				set[XY{float64(x) - 50.0, float64(y) - 25.0}] = true
				set[XY{float64(x) - 50.0, float64(y) + 25.0}] = true
			}
		}
	}
	all := make([]XY, 0, len(set))
	for xy := range set {
		all = append(all, xy)
	}
	return all
}
