func missingNumber(nums []int) int {
    ma := map[int]int{}
    for _,k:= range(nums){
        ma[k] =1
    }
    for i:=0;i<=len(nums);i++{
        _,prs := ma[i]
        if !prs{
            return i
        }
    }
    return -1
}