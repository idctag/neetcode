package arr

func ValidSudoku(board [][]byte) bool {
	rows := make([]int, 9)
	cols := make([]int, 9)
	squares := make([]int, 9)

	for c := range cols {
		for r := range rows {
			if board[c][r] == '.' {
				continue
			}
			val := board[c][r] - '1'
			bit := 1 << val
			sidx := (r/3)*3 + (c / 3)

			if rows[r]&bit != 0 || cols[c]&bit != 0 || squares[sidx]&bit != 0 {
				return false
			}

			rows[r] |= bit
			cols[c] |= bit
			squares[sidx] |= bit
		}
	}
	return true
}
