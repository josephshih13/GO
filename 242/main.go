package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt1, cnt2 := map[rune]int{}, map[rune]int{}
	for _, c := range s {
		if _, ok := cnt1[c]; ok {
			cnt1[c]++
		} else {
			cnt1[c] = 1
		}
	}

	for _, c := range t {
		if _, ok := cnt2[c]; ok {
			cnt2[c]++
		} else {
			cnt2[c] = 1
		}
	}

	for k, v := range cnt1 {
		if val, ok := cnt2[k]; !ok || val != v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("123")
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))

}
