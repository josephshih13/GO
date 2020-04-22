package main

import (
	"fmt"
	"regexp"
)

//`^[+-]?\d+$|^[+-]?\d+.\d+$`
func isNumber(s string) bool {
	pat := "^\\s*[+-]?\\d+\\s*$|^\\s*[+-]?\\d*\\.\\d+$\\s*|^\\s*[+-]?\\d*\\.\\d+e[+-]?\\d+\\s*$|^\\s*[+-]?\\d+e[+-]?\\d+\\s*$"
	matched, _ := regexp.MatchString(pat, s)
	return matched
}

func main() {

	fmt.Println(isNumber("0"))
	fmt.Println(isNumber("0.1"))
	fmt.Println(isNumber("abc"))
	fmt.Println(isNumber("1 a"))
	fmt.Println(isNumber("2e10"))
	fmt.Println(isNumber("-90e3"))
	fmt.Println(isNumber("1e"))
	fmt.Println(isNumber("e3"))
	fmt.Println(isNumber("6e-1"))
	fmt.Println(isNumber("99e2.5"))
	fmt.Println(isNumber("53.5e93"))
	fmt.Println(isNumber("--6"))
	fmt.Println(isNumber("-+3"))
	fmt.Println(isNumber("95a54e53"))
}
