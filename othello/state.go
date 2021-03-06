package othello

import "github.com/z-rui/game"

// N is the board size of the Othello game.
const N = 8

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
	Board          [N][N]Cell
	countO, countX uint8
	LastMove       Move
	Turn           Cell // must be O or X
}

// NewState returns a new state at the start of the game.
func NewState() *State {
	s := new(State)
	const (
		x = N/2 - 1
		y = x + 1
	)
	s.Board[x][x] = O
	s.Board[x][y] = X
	s.Board[y][x] = X
	s.Board[y][y] = O
	s.countO = 2
	s.countX = 2
	s.LastMove = invalidMove
	s.Turn = O
	return s
}

// clone clones a state.
func (s *State) clone() *State {
	t := new(State)
	t.Board = s.Board
	t.countO = s.countO
	t.countX = s.countX
	t.Turn = s.Turn
	return t
}

// Pass returns a new state after a player passes.
func (s *State) Pass() *State {
	t := s.clone()
	t.LastMove = invalidMove
	t.Turn ^= O ^ X
	return t
}

// Count returns the counts of O's and X's on the board.
func (s *State) Count() (o uint8, x uint8) {
	o = s.countO
	x = s.countX
	return
}

// Dim returns the dimension of the board
func (s *State) Dim() (rows int, cols int) {
	rows = N
	cols = N
	return
}

// Get returns the string representation at (i, j)
func (s *State) Get(i, j int) string {
	return s.Board[i][j].String()
}

// valueMap assigns a value to each cell of the board
var valueMap [N][N]int8

// valueMap generation
func init() {
	i, j := 0, 0
	for _, v := range [...]int8{
		/*1:*/ 99, -8, 8, 6,
		/*2: */ -24, -4, -3,
		/*3:        */ 7, 4,
		/*4:           */ 0,
	} {
		valueMap[i][j] = v
		valueMap[i][N-j-1] = v
		valueMap[j][i] = v
		valueMap[j][N-i-1] = v
		valueMap[N-i-1][j] = v
		valueMap[N-i-1][N-j-1] = v
		valueMap[N-j-1][i] = v
		valueMap[N-j-1][N-i-1] = v
		j++
		if j == N/2 {
			i++
			j = i
		}
	}
}

// Eval returns the evaluation of the current state.
func (s *State) Eval() (eval game.Evaluation) {
	if s.IsEnd() {
		switch {
		case s.countO > s.countX:
			eval = game.Won
		case s.countO < s.countX:
			eval = game.Lost
		}
	} else {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				switch s.Board[i][j] {
				case O:
					eval += game.Evaluation(valueMap[i][j])
				case X:
					eval -= game.Evaluation(valueMap[i][j])
				}
			}
		}
	}
	return
}

// IsEnd tells if the game has ended.
func (s *State) IsEnd() bool {
	// BUG() IsEnd won't return true if the state
	// is generated by passing twice.
	return s.countO == 0 || s.countX == 0 || s.countO+s.countX == N*N
}

// MustPass tells if the current user must pass.
func (s *State) MustPass() bool {
	var i, j uint8
	for i = 0; i < N; i++ {
		for j = 0; j < N; j++ {
			if (Move{i, j}).Allowed(s) {
				return false
			}
		}
	}
	return true
}

// Next returns all possible next states.
func (s *State) Next() (nxt []game.State) {
	nxt = make([]game.State, 0, 4)
	for _, mv := range validMoves {
		if t := s.Move(mv); t != nil {
			nxt = append(nxt, t)
		}
	}
	if len(nxt) == 0 && !s.IsEnd() && s.LastMove != invalidMove {
		nxt = append(nxt, s.Pass())
	}
	return
}

// Move returns the next state based on the move.
// It returns nil if the move is not allowed.
func (s *State) Move(m Move) (t *State) {
	i, j := int(m.I), int(m.J)
	if !m.Valid() || s.Board[i][j] != Empty {
		return nil
	}
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}
			n := s.match(i, j, di, dj)
			if n > 0 {
				// CoW 8x8 board can save time
				if t == nil {
					t = s.clone()
					t.Board[i][j] = s.Turn
				}
				t.flip(i, j, di, dj, n)
			}
		}
	}
	if t != nil {
		if t.Turn == O {
			t.countO++
		} else {
			t.countX++
		}
		t.LastMove = m
		t.Turn ^= O ^ X // switch player
	}
	return
}

// match finds how many discs will be reversed in the given direction.
func (s *State) match(i, j, di, dj int) int {
	n := 0
	for {
		i += di
		j += dj
		if !(0 <= i && i < N && 0 <= j && j < N) {
			return 0
		}
		switch s.Board[i][j] {
		case Empty:
			return 0
		case s.Turn:
			return n
		}
		n++
	}
}

// flip flips the given amount of discs in the given direction.
func (s *State) flip(i, j, di, dj, n int) {
	if s.Turn == O {
		s.countO += uint8(n)
		s.countX -= uint8(n)
	} else {
		s.countO -= uint8(n)
		s.countX += uint8(n)
	}
	for n > 0 {
		i += di
		j += dj
		s.Board[i][j] = s.Turn
		n--
	}
}
