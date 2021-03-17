func dfs(x,y int, used,valid [][]bool, matrix [][]int){
    dx := []int{0,0,1,-1}
    dy := []int{1,-1,0,0}
    used[x][y], valid[x][y] = true,true
    m,n := len(matrix),len(matrix[0])
    for i:=0;i<4;i++{
        xx:= x + dx[i]
        yy:= y + dy[i]
        if xx >= 0 && xx < m && yy>=0 && yy < n && !used[xx][yy] && matrix[xx][yy] >= matrix[x][y]{
            dfs(xx,yy,used,valid,matrix)
        }
    }
    
}

func make2dbool(m,n int)[][]bool{
    ret := [][]bool{}
    for i:=0;i<m;i++{
        tmp:=make([]bool,n)
        ret = append(ret,tmp)
            
    }
    return ret
}

func pacificAtlantic(matrix [][]int) [][]int {
    if len(matrix)==0{
        return [][]int{}
    }
    m,n := len(matrix),len(matrix[0])
    
    valid1 := make2dbool(m,n)
    used := make2dbool(m,n)
    
    for i :=0 ; i < m;i++{
        valid1[i][0] = true
    }
    for i :=0 ; i < n;i++{
        valid1[0][i] = true
    }
    
    update := true
    
    for update{
        update = false
        for i:=0 ; i < m;i++{
            for j:=0;j<n;j++{
                if valid1[i][j] && !used[i][j]{
                    update = true
                    dfs(i,j,used,valid1,matrix)
                }
            }
        }
    }
    
    
    
    valid2 := make2dbool(m,n)
    used = make2dbool(m,n)
    for i :=0 ; i < m;i++{
        valid2[i][n-1] = true
    }
    for i :=0 ; i < n;i++{
        valid2[m-1][i] = true
    }
    
    update = true
    
    for update{
        update = false
        for i:=0 ; i < m;i++{
            for j:=0;j<n;j++{
                if valid2[i][j] && !used[i][j]{
                    update = true
                    dfs(i,j,used,valid2,matrix)
                }
            }
        }
    }
    
    ret := [][]int{}
    
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if valid1[i][j] && valid2[i][j]{
                tmp := []int{i,j}
                ret = append(ret,tmp)
            }
        }
    }
    
    return ret
}
