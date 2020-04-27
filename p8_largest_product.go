package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)
func max(a int, b int) int{
	if a > b{
		return a;
	}
	return b;
}
func main(){
	data, err := ioutil.ReadFile("p8.txt")
	if err != nil{
		panic(err)
	}
	str:= string(data)
	str=strings.Replace(str,"\n","",-1)
	max_prod:=1
	for i := 0; i < len(str)-13; i++{
		prod:=1
		for j:=0; j < 13; j++{
			prod*=int(str[i+j])-48
		}
		if prod > max_prod{
			max_prod = prod
		}
	}
	fmt.Println(max_prod)
}
