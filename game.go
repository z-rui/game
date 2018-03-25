// Package game provides a standard framework for a two-player game
// and a MinMax algorithm for finding optimal moves.
package game

import "math"

// Evaluation is the number measuring the state of the game.
// It is positive if the state is advantageous to a chosen player
// (defined by the concrete State) and negative otherwise.
// The absolute value measures the likeliness for a player to win.
// If zero, the players are equally likely to win.
type Evaluation int32

const (
	// Lost refers to the chosen player will definitely lose.
	Lost Evaluation = math.MinInt32
	// Won refers to the chosen player will definitely win.
	Won  Evaluation = math.MaxInt32
)

// State represents an abstract state of the game.
type State interface {
	// Eval returns the evaluation of the current state.
	Eval() Evaluation
	// Next returns all possible states from the current state.
	Next() []State
}

// MinMax is the algorithm to find an optimal move for a current state.
// It finds the next state who will result in a minimum/maximum
// evaluation after certain iterations.
func MinMax(s State, iterations uint, findMin bool) (next State, eval Evaluation) {
	if findMin {
		return min(s, iterations, Lost)
	} else {
		return max(s, iterations, Won)
	}
}

func min(s State, iterations uint, limit Evaluation) (next State, eval Evaluation) {
	nxt := s.Next()
	if iterations == 0 || len(nxt) == 0 {
		eval = s.Eval()
		return
	}
	eval = Won
	for _, t := range nxt {
		_, e := max(t, iterations-1, eval)
		if next == nil || e < eval {
			next = t
			eval = e
			if e <= limit {
				break
			}
		}
	}
	return
}

func max(s State, iterations uint, limit Evaluation) (next State, eval Evaluation) {
	nxt := s.Next()
	if iterations == 0 || len(nxt) == 0 {
		eval = s.Eval()
		return
	}
	eval = Lost
	for _, t := range nxt {
		_, e := min(t, iterations-1, eval)
		if next == nil || e > eval {
			next = t
			eval = e
			if e >= limit {
				break
			}
		}
	}
	return
}
