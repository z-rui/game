package tictactoe

import (
	"github.com/z-rui/game"
	"testing"
)

const E = Empty

func TestEval(t *testing.T) {
	var s *State
	s = NewState()
	if s.Eval() != 0 {
		t.Errorf("new state not evaluated 0: %v", s.Eval())
	}
	s.Board = [N][N]Cell{
		{O, X, O},
		{X, O, X},
		{X, O, X},
	}
	s.LastMove = Move{0, 0}
	s.Turn = X
	if s.Eval() != 0 {
		t.Errorf("draw game not evaluated 0")
	}
	s.Board = [N][N]Cell{
		{O, X, O},
		{X, O, X},
		{X, O, X},
	}
	s.LastMove = Move{0, 0}
	s.Turn = X
	if s.Eval() != 0 {
		t.Errorf("draw game not evaluated 0")
	}
	s.Board = [N][N]Cell{
		{X, O, O},
		{O, O, X},
		{X, O, X},
	}
	s.LastMove = Move{1, 1}
	s.Turn = X
	if e := s.Eval(); e != game.Won {
		t.Errorf("won game not evaluated Won: %v", e)
	}
	s.Board = [N][N]Cell{
		{O, X, O},
		{X, X, X},
		{O, X, O},
	}
	s.LastMove = Move{1, 1}
	s.Turn = O
	if e := s.Eval(); e != game.Lost {
		t.Errorf("lost game not evaluated Lost: %v", e)
	}
}
