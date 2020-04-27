package main

import "fmt"

func main(){
	fib := make([]int64,100000)
	fib[0] = 0
	fib[1] = 1
	var sum int64 = 0
	for i:=2; i < len(fib); i++{
		fib[i]=fib[i-2]+fib[i-1]
		if fib[i] > 4000000{
			fmt.Println(i)
			break
		}
		if fib[i]%2 == 0{
			sum+=fib[i]
		}
	}
	fmt.Println(sum)
}
