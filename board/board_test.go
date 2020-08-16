package board_test

import (
	"fmt"
	"github.com/dgodd/settlers/board"
	"testing"
)

func TestNewSimple(t *testing.T) {
	board := board.NewSimple()

	fmt.Printf("%+v\n", board)
	t.Fatal()
}
