package tictactoe

import "github.com/z-rui/game"

// N is the board size of the Tic-Tac-Toe game
const N = 3

// Cell represents a cell of the board.
// It has three states: Empty, O and X.
type Cell uint8

const (
	Empty Cell = iota
	O
	X
)

// String converts a cell to the string representation.
func (c Cell) String() string {
	switch c {
	case O:
		return "O"
	case X:
		return "X"
	default:
		return " "
	}
}

// State represents the current state of the game.
type State struct {
	Board    [N][N]Cell
	LastMove Move
	Turn     Cell // must be O or X
}

// NewState returns a new state at the start of the game.
func NewState() *State {
	s := new(State)
	s.LastMove = invalidMove
	s.Turn = O
	return s
}

// Dim returns the dimension of the board
func (s *State) Dim() (int, int) {
	return N, N
}

// Get returns the string representation at (i, j)
func (s *State) Get(i, j int) string {
	return s.Board[i][j].String()
}

// Eval returns the evaluation of the current state.
func (s *State) Eval() (eval game.Evaluation) {
	if s.LastMove == invalidMove {
		return 0
	}
	won := s.match(0, 1)
	for k := -1; k <= 1; k++ {
		won = won || s.match(1, k)
	}
	if won {
		switch s.Turn {
		case O:
			return game.Lost
		case X:
			return game.Won
		}
	}
	return 0
}

// IsEnd tells if the game has ended.
func (s *State) IsEnd() bool {
	switch s.Eval() {
	case game.Won, game.Lost:
		return true
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if s.Board[i][j] == Empty {
				return false
			}
		}
	}
	return true
}

func (s *State) match(di, dj int) bool {
	i, j := int(s.LastMove.I), int(s.LastMove.J)
	cell := s.Board[i][j]
	for {
		i1, j1 := i-di, j-dj
		if 0 <= i1 && i1 < N && 0 <= j1 && j1 < N && s.Board[i1][j1] == cell {
			i, j = i1, j1
		} else {
			break
		}
	}
	n := 0
	for {
		n++
		i1, j1 := i+di, j+dj
		if 0 <= i1 && i1 < N && 0 <= j1 && j1 < N && s.Board[i1][j1] == cell {
			i, j = i1, j1
		} else {
			break
		}
	}
	return n == N
}

// Next returns all possible next states.
func (s *State) Next() (nxt []game.State) {
	if s.IsEnd() {
		return
	}
	nxt = make([]game.State, 0, 4)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if t := s.Move(Move{uint8(i), uint8(j)}); t != nil {
				nxt = append(nxt, t)
			}
		}
	}
	return
}

// Move returns the next state based on the move.
// It returns nil if the move is not allowed.
func (s *State) Move(m Move) (t *State) {
	if !m.Valid() || !m.Allowed(s) {
		return
	}
	t = new(State)
	t.Board = s.Board
	t.Board[m.I][m.J] = s.Turn
	t.LastMove = m
	t.Turn = s.Turn ^ (O ^ X)
	return t
}
