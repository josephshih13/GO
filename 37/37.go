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

func solveSudoku(board [][]byte) {
	z := [9][9][9]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 0; j++ {
			if board[i][j] != '.' {
				tar := board[i][j] - '0' - 1
				for k := 0; k < 9; k++ {
					z[i][k][tar] = true
					z[k][j][tar] = true
				}
				dx, dy := i/3, j/3
				for x := 0; x < 3; x++ {
					for y := 0; y < 3; y++ {
						z[dx+x][dy+y][tar] = true
					}
				}
			}
		}
	}
	update := true
	for update {
		update = false
		for i := 0; i < 9; i++ {
			if update {
				break
			}
			for j := 0; j < 9; j++ {
				if board[i][j] == '.' {
					ans := -1
					for k := 0; k < 9; k++ {
						if z[i][j][k] == false {
							if ans != -1 {
								ans = -1
								break
							}
							ans = k
						}
					}
					if ans != -1 {
						board[i][j] = byte('0' + ans + 1)
						tar := ans
						for kk := 0; kk < 9; kk++ {
							z[i][kk][tar] = true
							z[kk][j][tar] = true
						}
						dx, dy := i/3, j/3
						for x := 0; x < 3; x++ {
							for y := 0; y < 3; y++ {
								z[dx+x][dy+y][tar] = true
							}
						}
						update = true
						break
					}
				}
			}
		}
	}
}

func main() {

}
