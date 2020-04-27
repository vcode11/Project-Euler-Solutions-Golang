package main

import "fmt"

func main(){
	var arr [10005]int
	arr[1] = 0
	for i := 1; i < len(arr); i++{
		for j:=2*i; j < len(arr); j+=i{
			arr[j]+=i
		}
	}
	sum := 0
	for i := 2; i <= 10000; i++{
		if arr[i]  <= 10000 && arr[arr[i]] == i && i != arr[i]{
			sum+=i
		}
	}
	fmt.Println(sum)
}
