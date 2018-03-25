package othello

import "sort"

// Move represents a position on the board
type Move struct {
	I, J uint8
}

// Strings converts a Move to the string representation
func (m Move) String() string {
	if m == invalidMove {
		return "(pass)"
	}
	return string([]byte{byte(m.I) + 'A', byte(m.J) + '1'})
}

// invalidMove represents an invalid move;
// also for representing a pass.
var invalidMove = Move{N, N}

// validMoves contains all valid moves,
// ordered by the value assigned in valueMap
var validMoves [N * N]Move

// validMoves generation
func init() {
	var i, j, k uint8
	for i = 0; i < N; i++ {
		for j = 0; j < N; j++ {
			validMoves[k] = Move{i, j}
			k++
		}
	}
	sort.Slice(validMoves[:], func(i, j int) bool {
		x := validMoves[i]
		y := validMoves[j]
		return x.value() > y.value()
	})
}

func (m Move) value() int8 {
	return valueMap[m.I][m.J]
}

// Valid tells if the move is a valid position (not out-of-bound)
// on the board.
func (m Move) Valid() bool {
	return 0 <= m.I && m.I < N && 0 <= m.J && m.J < N
}

// Allowed tells if the move is allowed according to the game's rule.
func (m Move) Allowed(s *State) bool {
	i, j := int(m.I), int(m.J)
	if s.Board[i][j] != Empty {
		return false
	}
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}
			if s.match(i, j, di, dj) > 0 {
				return true
			}
		}
	}
	return false
}
