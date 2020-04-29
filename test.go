package main
import (
"fmt"
"sort"
)

type Pair struct {
Key   string
Value int
}

type ByKey []Pair

func (s ByKey) Len() int {
return len(s)
}

func (s ByKey) Swap(i, j int) {
s[i], s[j] = s[j], s[i]
}

func (s ByKey) Less(i, j int) bool {
return s[i].Key < s[j].Key
}

func main() {
pairs := ByKey{{"a", 1}, {"b", 0}, {"c", 2}}
// Sort by Key
sort.Sort(pairs)
fmt.Println(pairs) // [{a 1} {b 0} {c 2}]
}