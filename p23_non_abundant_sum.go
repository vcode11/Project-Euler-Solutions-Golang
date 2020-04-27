package main

import "fmt"
var dp [30000]int
func main(){
	var arr [30000]int
	arr[1] = 0
	abundant:=make([]int, 0)
	for i := 1; i < len(arr); i++{
		if arr[i] > i{
			abundant = append(abundant,i)
		}
		for j:=2*i; j < len(arr); j+=i{
			arr[j]+=i
		}
	}
	sum := 0
	fmt.Println(len(abundant))
	for i := 0; i < len(abundant); i++{
		for j := i; j < len(abundant); j++{
			if abundant[i]+abundant[j] < len(dp){
				dp[abundant[i]+abundant[j]] = 1
			}
		}
	}
	for i := 1; i <= 28123; i++{
		if dp[i] == 0{
			sum+=i
		}
	}
	fmt.Println(sum)
}

