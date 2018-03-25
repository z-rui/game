package tictactoe

// Move represents a position on the board
type Move struct {
	I, J uint8
}

var invalidMove = Move{N, N}

// Strings converts a Move to the string representation
func (m Move) String() string {
	return string([]byte{byte(m.I) + 'A', byte(m.J) + '1'})
}

// Valid tells if the move is a valid position (not out-of-bound)
// on the board.
func (m Move) Valid() bool {
	return 0 <= m.I && m.I < N && 0 <= m.J && m.J < N
}

// Allowed tells if the move is allowed according to the game's rule.
func (m Move) Allowed(s *State) bool {
	return s.Board[m.I][m.J] == Empty
}
