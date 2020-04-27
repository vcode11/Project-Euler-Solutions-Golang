package main

import "fmt"
import "time"

func main(){
	primes:=make([]int,1000000)
	cnt:=0
	n:=10001
	start:=time.Now()
	for i:=2; i < len(primes); i++{
		if primes[i] == 0{
			cnt++
			if(cnt == n){
				fmt.Println(i);
				break;
			}
			for j:=2*i; j < len(primes); j+=i{
				primes[j]=1;
			}
		}
	}
	fmt.Println("Execution Time:",time.Since(start))
}
