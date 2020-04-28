package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	sp := strings.Split(path,"/")
	//fmt.Println(sp)
	st := []string{}
	for i:= range sp{
		if sp[i] == "" || sp[i] == "."{
			continue
		}
		if sp[i] == ".."{
			if len(st)>0{
				st = st[:len(st)-1]
			}
		} else{
			st = append(st,sp[i])
		}
	}
	var sb strings.Builder
	//sb.WriteByte('/')
	for i:=range st{
		sb.WriteByte('/')
		sb.WriteString(st[i])

	}
	if sb.String() == ""{
		return "/"
	}
	return sb.String()
}

func main() {
	fmt.Println(simplifyPath("/a/../../b/../c//.//"))
}
