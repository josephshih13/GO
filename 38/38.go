package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		tmp := ""
		pre, cnt := 'q', 0
		for _, c := range s {
			if pre == 'q' {
				// ss, _ := strconv.Atoi(cnt)
				cnt, pre = 1, c
			} else if pre != c {
				tmp += strconv.Itoa(cnt) + string(pre)
				cnt, pre = 1, c
			} else {
				cnt++
			}
		}
		if cnt != 0 {
			// ss, _ := strconv.Atoi(cnt)
			// tmp += ss + string(pre)
			tmp += strconv.Itoa(cnt) + string(pre)
		}
		s = tmp
	}
	return string(s)
}

func main() {
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(5))
	fmt.Println(len(countAndSay(30)))
}
