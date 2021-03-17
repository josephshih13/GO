func maxint(a,b int) int{
    if a>b{
        return a
    }
    return b
}

func rob1(nums []int) int {
    if len(nums) == 0{
        return 0
    }
    if len(nums) < 3{
        ret := 0
        for _, v:=range nums{
            ret = maxint(v,ret)
        }
        return ret
    }
    dp := make([]int,len(nums))
    dp[0],dp[1] = nums[0], nums[1]
    dp[2] = nums[0] + nums[2]
    ret := maxint(dp[2],dp[1])
    for i := 3; i < len(nums);i++{
        dp[i] = maxint(dp[i-2],dp[i-3]) + nums[i]
        ret = maxint(ret,dp[i])
    }
    return ret
}

func rob(nums []int) int{
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1{
        return nums[0]
    }
    return maxint(rob1(nums[1:]), rob1(nums[:len(nums)-1]))
}