package main

import (
	"fmt"
	"github.com/z-rui/game"
	"github.com/z-rui/game/tictactoe"
)

type CpuPlayer struct {
	name  string
	level uint
}

func (p *CpuPlayer) Name() string {
	return p.name
}

func (p *CpuPlayer) Next(s *tictactoe.State) *tictactoe.State {
	var next game.State
	findMin := s.Turn == tictactoe.X
	if *verboseSearch {
		var eval game.Evaluation
		for _, t := range s.Next() {
			ns, e := game.MinMax(t, p.level, !findMin)
			fmt.Printf("Move %v: value = %v", t.(*tictactoe.State).LastMove, e)
			if ns != nil {
				fmt.Printf(", opponent = %v", ns.(*tictactoe.State).LastMove)
			}
			fmt.Println()
			if next == nil || (findMin && e < eval || !findMin && e > eval) {
				next, eval = t, e
			}
		}
	} else {
		next, _ = game.MinMax(s, p.level, findMin)
	}
	if next == nil {
		return nil
	}
	return next.(*tictactoe.State)
}
