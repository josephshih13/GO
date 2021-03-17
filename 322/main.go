func minint(a,b int)int{
    if a<b{
        return a
    }
    return b
}

func coinChange(coins []int, amount int) int {
    BIG := 1000000000
    dp := [][]int{}
    for _,_ = range(coins){
        tmp := make([]int,amount+1)
        for i := range(tmp){
            tmp[i] = BIG
        }
        dp  = append(dp, tmp)
    }
    for i := range(coins){
        dp[i][0] = 0
        for j:= 0 ; j <= amount; j++{
            if j - coins[i] >= 0 && i != 0{
                dp[i][j] = minint(dp[i][j-coins[i]]+1,dp[i-1][j])
            }else if i != 0{
                dp[i][j] = dp[i-1][j]
            }else if j - coins[i] >= 0{
                dp[i][j] = dp[i][j-coins[i]]+1
            }
        }
    }
    if dp[len(coins)-1][amount] >= BIG{
        return -1
    }
    return dp[len(coins)-1][amount]
}