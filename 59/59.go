package main

import "fmt"

// func nextdir(x,y int, matrix,used [][]int) int{
// 	xx,yy := len(matrix), len(matrix[0])
// 	if y +1 < yy && !used[x][y+1]{
// 		return 0
// 	} else if x+1 < xx && !used[x+1][y]{
// 		return 1
// 	} else if y -1 >=0 && !used[x][y-1]{
// 		return 2
// 	} else if x-1 >=0 && !used[x-1][y]{
// 		return 3
// 	}
// 	return -1
// }

// func goright(x,y int, matrix,used [][]int, ans []int) int,int{
// 	xx,yy := len(matrix), len(matrix[0])
// 	for y < yy{
// 		used[x][y] = true
// 		ans = append(ans,matrix[x][y])
// 		y++
// 	}
// 	return x,y
// }

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	x, y := 0, 0
	dx, dy := [4]int{0, 1, 0, -1}, [4]int{1, 0, -1, 0}
	m, n := len(matrix), len(matrix[0])
	used := [][]bool{}
	for i := 0; i < m; i++ {
		tmp := make([]bool, n)
		used = append(used, tmp)
	}
	used[0][0] = true
	dir := 0
	if n == 1 {
		dir = 1
	}
	ans := []int{matrix[0][0]}
	fail := 0
	for fail < 4 {
		update := false
		xx, yy := x+dx[dir], y+dy[dir]
		for xx >= 0 && xx < m && yy >= 0 && yy < n && !used[xx][yy] {
			// fmt.Println(xx, yy)
			update, used[xx][yy] = true, true
			x, y = xx, yy
			ans = append(ans, matrix[xx][yy])
			xx, yy = xx+dx[dir], yy+dy[dir]
		}
		dir = (dir + 1) % 4
		if !update {
			fail++
		} else {
			fail = 0
		}
	}
	return ans
}

func generateMatrix(n int) [][]int {
	ret := [][]int{}
	used := [][]bool{}
	for i := 0; i < n; i++ {
		ret = append(ret, make([]int, n))
		used = append(used, make([]bool, n))
	}
	x, y := 0, 0
	k := 2
	used[x][y] = true
	ret[x][y] = 1
	dir, fail := 0, 0
	dx, dy := [4]int{0, 1, 0, -1}, [4]int{1, 0, -1, 0}
	for fail < 4 {
		update := false
		xx, yy := x+dx[dir], y+dy[dir]
		for xx >= 0 && xx < n && yy >= 0 && yy < n && !used[xx][yy] {
			// fmt.Println(xx, yy)
			update, used[xx][yy] = true, true
			ret[xx][yy] = k
			x, y = xx, yy
			// ret[x][y] = k
			k++
			xx, yy = xx+dx[dir], yy+dy[dir]
		}
		dir = (dir + 1) % 4
		if !update {
			fail++
		} else {
			fail = 0
		}
	}
	return ret
}

func main() {

	fmt.Println(generateMatrix(2))
	// fmt.Println(spiralOrder(a2))
}
