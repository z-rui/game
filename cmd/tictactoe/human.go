package main

import (
	"fmt"
	"github.com/z-rui/game/tictactoe"
	"log"
	"unicode"
)

type HumanPlayer struct {
	name string
}

func (p *HumanPlayer) Name() string {
	return p.name
}

func askPlaying() tictactoe.Cell {
	for {
		fmt.Print("Do you want to play as O or X? ")
		answer, err := stdin.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		if err == nil && len(answer) >= 2 {
			switch unicode.ToUpper(rune(answer[0])) {
			case 'O':
				return tictactoe.O
			case 'X':
				return tictactoe.X
			}
		}
		fmt.Println("Sorry, but that does not make sense.")
	}
}

func (p *HumanPlayer) Next(s *tictactoe.State) (t *tictactoe.State) {
	if s.IsEnd() {
		return nil
	}
	for {
		fmt.Print("Where do you want to go? ")
		coord, err := stdin.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		if len(coord) >= 3 {
			var m tictactoe.Move
			m.I = uint8(unicode.ToUpper(rune(coord[0])) - 'A')
			m.J = uint8(coord[1] - '1')
			t = s.Move(m)
			if t != nil {
				return
			}
		}
		fmt.Println("Sorry, but that does not make sense.")
	}
	return
}
