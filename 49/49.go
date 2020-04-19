package main

import (
	"fmt"
)

func strsummary(str string) [26]int {
	cnt := [26]int{}
	for _, r := range str {
		cnt[r-'a']++
	}
	// fmt.Println(cnt)
	return cnt
}

func groupAnagrams(strs []string) [][]string {
	group := make(map[[26]int][]string)
	for _, str := range strs {
		tmp := strsummary(str)
		_, exist := group[tmp]
		if exist {
			group[tmp] = append(group[tmp], str)
		} else {
			group[tmp] = []string{str}
		}
	}
	// fmt.Println(group)
	ans := [][]string{}
	for _, v := range group {
		ans = append(ans, v)
	}
	// return [][]string{}
	return ans
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
