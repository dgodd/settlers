package board_test

import (
	"testing"
	"github.com/dgodd/settlers/board"
	"fmt"
)

func TestNewSimple(t *testing.T) {
	board := board.NewSimple()

	fmt.Printf("%+v\n", board)
	t.Fatal()
}
