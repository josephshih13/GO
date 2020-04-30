package main

import "fmt"

func minWindow(s string, t string) string {
	ma := make(map[byte]int)
	for i:=range t{
		c:=t[i]
		_,ex := ma[c]
		if !ex{
			ma[c]=1
		}else{
			ma[c]++
		}
	}
	fmt.Println(ma)
	tar := len(t)
	st,cnt,ml :=0,0,1000000000
	var ret_st,ret_end int
	ret_st,ret_end=-1,-1
	for i:=0;i<len(s);i++{
		_,ex:=ma[s[i]]
		if ex {
			ma[s[i]]--
			if ma[s[i]]>=0{
				cnt++
			}
		}
		for cnt == tar{
			//fmt.Println(st,i)
			if ml > i-st+1{
				ml = i-st+1
				ret_st,ret_end = st,i
			}
			_,exx:=ma[s[st]]
			if exx{
				ma[s[st]]++
				if ma[s[st]]>0{
					cnt--
				}

			}
			st++
		}
	}
	if ret_st == -1{
		return ""
	}
	return string(s[ret_st:ret_end+1])
}

func main() {
	fmt.Println(minWindow("a","aa"))
}
