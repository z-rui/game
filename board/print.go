package board

import (
	"io"
	"bufio"
)

type Board interface {
	// Get gets the string representation at (i, j)
	Get(i, j int) string
	// Dim returns the dimension of the board
	Dim() (rows int, cols int)
}

// UnicodeBox and AsciiBox are box-drawing charset
// that can be passed to Print.
var (
	UnicodeBox = [][]rune{
		[]rune("╔═╤╗"),
		[]rune("╟─┼╢)"),
		[]rune("╚═╧╝"),
		[]rune("║│"),
	}
	AsciiBox = [][]rune{
		[]rune("+-++"),
		[]rune("+-++"),
		[]rune("+-++"),
		[]rune("||"),
	}
)

func printRule(w *bufio.Writer, cols int, boxDrawing []rune) {
	w.WriteRune(' ')
	w.WriteRune(boxDrawing[0])
	i := 0
	for {
		w.WriteRune(boxDrawing[1])
		i++
		if i == cols {
			break
		}
		w.WriteRune(boxDrawing[2])
	}
	w.WriteRune(boxDrawing[3])
	w.WriteRune('\n')
}

// Print prints the board to Writer.
// Pass UnicodeBox or AsciiBox as the third argument
// to use different style.
func Print(writer io.Writer, b Board, boxDrawing [][]rune) {
	rows, cols := b.Dim()

	//BUG: larges boards (more than columns) are not supported yet.
	if cols > 9 {
		panic("large Board not supported yet")
	}
	w := bufio.NewWriter(writer)
	w.WriteRune(' ')
	for i := 0; i < cols; i++ {
		w.WriteRune(' ')
		w.WriteRune(rune('1' + i))
	}
	w.WriteRune('\n')

	printRule(w, cols, boxDrawing[0])
	i := 0
	for {
		w.WriteRune(rune('A' + i))
		w.WriteRune(boxDrawing[3][0])
		j := 0
		for {
			w.WriteString(b.Get(i, j))
			j++
			if j == cols {
				break
			}
			w.WriteRune(boxDrawing[3][1])
		}
		w.WriteRune(boxDrawing[3][0])
		w.WriteRune('\n')
		i++
		if i == rows {
			break
		}
		printRule(w, cols, boxDrawing[1])
	}
	printRule(w, cols, boxDrawing[2])
	w.Flush()
}
