package main

import (
	"fmt"
)

func isValid(s string) bool {
	st := []byte{}
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			st = append(st, byte(c))
		case ')':
			if len(st) == 0 || st[len(st)-1] != '(' {
				return false
			}
			st = st[:len(st)-1]
		case ']':
			if len(st) == 0 || st[len(st)-1] != '[' {
				return false
			}
			st = st[:len(st)-1]
		case '}':
			if len(st) == 0 || st[len(st)-1] != '{' {
				return false
			}
			st = st[:len(st)-1]
		}
	}
	if len(st) != 0 {
		return false
	}
	return true
}

func main() {
	// s := "l12"
	// s[0] = 'f'
	// fmt.Println(s)
	// fmt.Println(strconv.Atoi(string(s[1])))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))

}
