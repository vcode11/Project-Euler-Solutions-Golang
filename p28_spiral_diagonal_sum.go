package main

import "fmt"

func main(){
	sum := 1
	k := 2
	cnt := 0
	for i := 3; i <= 1001*1001; i+=k{
		sum+=i;
		cnt++
		if cnt%4 == 0 {
			k+=2
		}
	}
	fmt.Println(sum)
}
