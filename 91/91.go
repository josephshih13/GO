package main

import "fmt"

func numDecodings(s string) int {
	dp := make([]int,len(s)+1)
	dp[0] = 1
	for i:=0;i<len(s);i++{
		if s[i] == '0' {
			if i-1 >= 0 && (s[i-1] == '1' || s[i-1] == '2') {
				dp[i+1] = dp[i-1]
				continue
			} else {
				return 0
			}
		}
		dp[i+1]+=dp[i]
		if i-1>=0&& (s[i-1] == '1' || s[i-1] == '2' && s[i] <= '6'){
			dp[i+1]+=dp[i-1]
		}

	}
	return dp[len(s)]
	//if s[0] == '0'{
	//	return 0
	//}
	//if len(s) ==1 {
	//	return 1
	//}
	//dp := make([]int,len(s))
	//dp[0] = 1
	//dp[1] = 1
	//if s[1]!='0'&&(s[0] == '1' || s[0] == '2' && s[1] <='6'){
	//	dp[1]++
	//}
	//if s[1] == '0' && s[0] != '1' &&s[0]!= '2'{
	//	return 0
	//}
	//for i:=2;i<len(s);i++{
	//	dp[i] = dp[i-1]
	//	if s[i] != '0' &&(s[i-1] == '1' || s[i-1] == '2' && s[i] <='6'){
	//		dp[i] += dp[i-2]
	//	}
	//	if s[i] == '0' && s[i-1] != '1' &&s[i-1]!= '2'{
	//		return 0
	//	}
	//}
	//return dp[len(s)-1]
}

func main() {
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("226"))
}
