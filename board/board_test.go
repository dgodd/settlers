package board_test

import (
	"github.com/dgodd/settlers/board"
	"math"
	"testing"
)

func TestNewSimple(t *testing.T) {
	board := board.NewSimple()
	if board.Corners == nil {
		t.Fatal()
	}
}

func TestDistance(t *testing.T) {
	xy1 := board.XY{X: 100.0, Y: 100.0}
	xy2 := board.XY{X: 110.0, Y: 110.0}
	d := xy1.Distance(&xy2)
	if math.Abs(d-14.142) > 0.001 {
		t.Fatalf("Got %v, expected %v", d, 14.142)
	}
}
