package main

import "fmt"

func exist(board [][]byte, word string) bool {
	if len(board)==0||len(board[0])==0{
		return false
	}
	used:=[][]bool{}
	for i:=0;i<len(board);i++{
		tmp:=make([]bool,len(board[0]))
		used=append(used,tmp)
	}
	var dfs func (int,int,int) bool
	dfs = func (x,y,idx int) bool{
		//fmt.Println(x,y,idx)
		if idx == len(word){
			//fmt.Println("YR")
			return true
		}
		dx,dy := []int{1,0,-1,0},[]int{0,-1,0,1}
		for i:=0;i<4;i++{
			xx,yy:= x+dx[i],y+dy[i]
			if xx<0 || xx>=len(board) || yy<0||yy>=len(board[0]){
				continue
			}
			if !used[xx][yy] && board[xx][yy] == word[idx]{
				used[xx][yy]=true
				if dfs(xx,yy,idx+1){
					//fmt.Println("YY")
					return true
				}
				used[xx][yy]=false
			}
		}
		//fmt.Println("N0")
		return false
	}
	for i:=0;i<len(board);i++{
		for j:=0;j<len(board[0]);j++{
			if board[i][j] == word[0]{
				used[i][j]=true
				if dfs(i,j,1){
					return true
				}
				used[i][j]=false
			}
		}
	}
	return false
}

func main() {
	a:=[][]byte{
		{'A','B','C','E'},
		{'S','F','C','S'},
		{'A','D','E','E'},
	}
	fmt.Println(exist(a,"ABCCED"))
}
