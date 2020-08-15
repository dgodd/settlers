package board

import (
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
	Klass int
	Number int
	HasRobber bool
}

type Board struct {
	Tiles [][]Tile
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
		Tiles: tiles,
	}
}

func NewSimple() Board {
	tilePlaces := [][]byte{
		[]byte{'n', 'n', 'w', 'w', 'w', 'w', 'n'},
		[]byte{'n', 'w', 'l', 'l', 'l', 'w', 'n'},
		[]byte{'n', 'w', 'l', 'l', 'l', 'l', 'w'},
		[]byte{'w', 'l', 'l', 'l', 'l', 'l', 'w'},
		[]byte{'n', 'w', 'l', 'l', 'l', 'l', 'w'},
		[]byte{'n', 'w', 'l', 'l', 'l', 'w', 'n'},
		[]byte{'n', 'n', 'w', 'w', 'w', 'w', 'n'},
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
