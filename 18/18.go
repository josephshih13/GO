package main

import (
	"fmt"
	"sort"
)

func mapfind(m map[int]int, tar int) bool {
	cnt, prs := m[tar]
	return prs && cnt > 0
}

func maxint(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minint(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func fourSum(nums []int, target int) [][]int {
	ans := [][]int{}
	// if len(nums) < 3 {
	// 	fmt.Println("???")
	// 	return ans
	// }
	z := make(map[int]int)
	for _, i := range nums {
		_, prs := z[i]
		if !prs {
			z[i] = 1
		} else {
			z[i]++
		}
	}
	// tt, prs := z[0]
	// for k := range z {
	// 	if prs && k < 0 {
	// 		_, tmp := z[-k]
	// 		if tmp {
	// 			ans = append(ans, []int{k, 0, -k})
	// 		}
	// 	}
	// 	if k%2 == 0 && k != 0 {
	// 		fmt.Println(k)
	// 		t, tmp := z[(-k)/2]
	// 		if tmp && t >= 2 {
	// 			ans = append(ans, []int{k, -k / 2, -k / 2})
	// 		}
	// 	}
	// }
	// if tt >= 3 {
	// 	ans = append(ans, []int{0, 0, 0})
	// }
	set := map[[4]int]bool{}
	for k1 := range z {
		z[k1]--
		for k2 := range z {
			if z[k2] == 0 {
				continue
			}
			z[k2]--
			for k3 := range z {
				if z[k3] == 0 {
					continue
				}
				z[k3]--
				k4 := target - k1 - k2 - k3
				if mapfind(z, k4) {
					tmp := []int{k1, k2, k3, k4}
					// tmp := []int{1, 2, 3}
					sort.Ints(tmp)
					var arr [4]int
					copy(arr[:], tmp)
					set[arr] = true
				}
				z[k3]++
			}
			// k3 := -k1 - k2
			// if mapfind(z, k3) {
			// 	// ans = append(ans, []int{k1, k2, k3})
			// 	// fmt.Println(k1, k2, k3)
			// 	c1, c3 := maxint(k1, maxint(k2, k3)), minint(k1, minint(k2, k3))
			// 	c2 := k1 + k2 + k3 - c1 - c3
			// 	set[[3]int{c1, c2, c3}] = true
			// }
			z[k2]++
		}
		z[k1]++
	}
	// fmt.Println(set)
	// var test []int
	// fmt.Println("???")
	for k := range set {
		// fmt.Println("TEST:", test)
		// fmt.Println(reflect.TypeOf(k[:]))
		// fmt.Println(k[:])
		// fmt.Println("b:", ans)
		tmp := make([]int, 4)
		copy(tmp, k[:])
		// fmt.Println("tmp:", tmp)
		ans = append(ans, tmp)
		// fmt.Println("af", ans)
		// test = tmp
		// fmt.Println("TEST2:", test)
	}
	// ll := [][3]int{[3]int{1, 0, -1}, [3]int{2, -1, -1}}
	// fmt.Println(ll)
	// for _, k := range ll {
	// 	// fmt.Println(k)
	// 	fmt.Println("TEST:", test)
	// 	fmt.Println(reflect.TypeOf(k[:]))
	// 	fmt.Println(k[:])
	// 	fmt.Println("b:", ans)
	// 	ans = append(ans, k[:])
	// 	fmt.Println("af", ans)
	// 	test = k[:]
	// 	fmt.Println("TEST2:", test)
	// }

	return ans
}

func main() {
	//test("  -1233mfgmfm")
	// fmt.Println(myAtoi("words and 987"))
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
