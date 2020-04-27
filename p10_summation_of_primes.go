package main

import "fmt"
import "time"

func main(){
	var primes [2000001]int
	sum:=0
	start:=time.Now()
	for i:=2; i < len(primes); i++{
		if primes[i] == 0{
			sum+=i
			for j:=2*i; j < len(primes); j+=i{
				primes[j]=1;
			}
		}
	}
	fmt.Println("Sum is",sum,"Execution time ",time.Since(start))
}
