package main

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		test := [9]bool{}
		test2 := [9]bool{}
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if test[board[i][j]-'1'] {
					return false
				}
				test[board[i][j]-'1'] = true
			}
			if board[j][i] != '.' {
				if test2[board[j][i]-'1'] {
					return false
				}
				test2[board[j][i]-'1'] = true
			}
		}
	}
	for dx := 0; dx < 3; dx++ {
		for dy := 0; dy < 3; dy++ {
			//square
			test := [9]bool{}
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					xx, yy := dx*3+x, dy*3+y
					if board[xx][yy] != '.' {
						if test[board[xx][yy]-'1'] {
							return false
						}
						test[board[xx][yy]-'1'] = true
					}
				}
			}
		}
	}
	return true
}

func main() {

}
