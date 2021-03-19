import "sort"

type interval struct {
    st,ed  int
}

type inters []interval

func (a inters) Len() int           { return len(a) }
func (a inters) Less(i, j int) bool { return a[i].ed < a[j].ed }
func (a inters) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func eraseOverlapIntervals(intervals [][]int) int {
    var li inters 
    
    for _,v := range(intervals){
        tmp := interval{v[0],v[1]}
        li = append(li,tmp)
    }
    sort.Sort(li)
    
    k := -1000000
    ret :=0
    
    for _,v := range(li){
        if v.st >= k{
            ret++
            k = v.ed
        }
    }
    return len(intervals) - ret
    
}