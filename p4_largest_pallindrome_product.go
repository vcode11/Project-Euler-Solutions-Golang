package main

import "fmt"
import "strconv"

func is_palin(n int64) bool{
	s:=strconv.FormatInt(n,10)
	flag:=true
	for i := 0; i < len(s); i++{
		if s[i] != s[len(s)-i-1]{
			flag=false
		}
	}
	return flag
}

func main(){
	var num int64 = 0
	for i:=101; i < 1000; i++{
		for j:=101; j < 1000; j++{
			m:=int64(i*j)
			if is_palin(m){
				if m > num{
					num = m
				}
			}
		}
	}
	fmt.Println(num)
}
