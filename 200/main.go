func trav(x,y int, used [][]bool,grid [][]byte) {
    dx := []int{0,0,1,-1}
    dy := []int{1,-1,0,0}
    used[x][y]=true
    for i:=0;i<4;i++{
        xx := x + dx[i]
        yy := y + dy[i]
        if xx >= 0 && yy >= 0 && xx < len(grid) && yy < len(grid[0]) && grid[xx][yy] == '1' && !used[xx][yy]{
            trav(xx,yy,used,grid)
        }
    }
}

func numIslands(grid [][]byte) int {
    used := [][]bool{}
    for _,_ = range(grid){
        tmp := make([]bool,len(grid[0]))
        used = append(used, tmp)    
    }
    
    ret :=0
    
    for i:=0;i<len(grid);i++{
        for j:=0; j<len(grid[0]);j++{
            if !used[i][j] && grid[i][j] == '1'{
                trav(i,j,used,grid)
                ret++
            }
        }
    }
    return ret
}