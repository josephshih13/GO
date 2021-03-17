func countBits(num int) []int {
    ret := []int{0}
    if num == 0{
        return ret
    }
    
    k := 1
    for i := 1; i <= num ; i++{
        if i == k * 2{
            k *= 2
        }
        ret = append(ret,ret[i-k]+1)
    }
    return ret
}