package main

import "fmt"

func countSubstringsFromMiddle(l []int, middle_idx int) int {
	ret := 0
	left, right := middle_idx, middle_idx
	if l[middle_idx] == 0 {
		left--
		right++
	}
	for left >= 0 && right < len(l) && l[left] == l[right] {
		ret++
		left -= 2
		right += 2
	}
	return ret
}

func countSubstrings(s string) int {
	l := []int{0}

	for _, c := range s {
		l = append(l, int(c), 0)
	}

	ret := 0

	for i, _ := range l {
		ret += countSubstringsFromMiddle(l, i)
	}

	return ret
}

func main() {
	fmt.Println("123")
	fmt.Println(countSubstrings("aaa"))
	fmt.Println(countSubstrings("abc"))
}
