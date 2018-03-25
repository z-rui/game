package othello

import "testing"

const E = Empty

func TestHasValidMoves(t *testing.T) {
	s := State{
		Board: [N][N]Cell{
			{O, O, O, O, O, O, O, O},
			{O, O, X, X, X, X, X, X},
			{O, O, O, O, O, O, X, X},
			{O, O, O, O, O, O, X, X},
			{O, O, O, O, O, O, X, X},
			{O, O, O, O, O, O, X, X},
			{O, O, O, O, O, O, O, X},
			{O, X, O, E, O, O, E, O},
		},
		LastMove: Move{6, 7},
		Turn:     O,
	}
	if !s.MustPass() {
		t.Errorf("This should not have valid moves: %v", s)
	}
	s.Board = [N][N]Cell{
		{E, O, O, O, O, E, E},
		{E, E, O, O, O, O, E},
		{O, O, O, O, O, O, O},
		{E, O, O, O, O, O, O},
		{E, X, X, X, X, O, O},
		{X, X, X, X, O, O, O},
		{X, X, X, O, O, O, E},
		{X, X, O, E, E, E, O},
	}
	s.LastMove = Move{7, 2}
	s.Turn = O
	if s.MustPass() {
		t.Errorf("This should have valid moves: %v", s)
	}
}

func TestFlip(t *testing.T) {
	s := NewState()
	n := s.match(2, 4, 1, 0)
	if n != 1 {
		t.Errorf("Should match 1, got %d", n)
	}
	s.Board[2][4] = O
	s.flip(2, 4, 1, 0, n)
	reference := [N][N]Cell{
		{E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E},
		{E, E, E, E, O, E, E, E},
		{E, E, E, O, O, E, E, E},
		{E, E, E, X, O, E, E, E},
		{E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E},
	}
	if s.Board != reference {
		t.Errorf("Board mismatch")
	}
}

func TestMove(t *testing.T) {
	s := NewState()
	s = s.Move(Move{2, 4})
	reference := [N][N]Cell{
		{E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E},
		{E, E, E, E, O, E, E, E},
		{E, E, E, O, O, E, E, E},
		{E, E, E, X, O, E, E, E},
		{E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E},
		{E, E, E, E, E, E, E, E},
	}
	if s.Board != reference {
		t.Errorf("Board mismatch")
	}
}

func TestValueMap(t *testing.T) {
	reference := [N][N]int8{
		{99, -8, 8, 6, 6, 8, -8, 99},
		{-8, -24, -4, -3, -3, -4, -24, -8},
		{8, -4, 7, 4, 4, 7, -4, 8},
		{6, -3, 4, 0, 0, 4, -3, 6},
		{6, -3, 4, 0, 0, 4, -3, 6},
		{8, -4, 7, 4, 4, 7, -4, 8},
		{-8, -24, -4, -3, -3, -4, -24, -8},
		{99, -8, 8, 6, 6, 8, -8, 99},
	}
	if valueMap != reference {
		t.Errorf("wrong valueMap: %v", valueMap)
	}
}
