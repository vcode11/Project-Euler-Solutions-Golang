package main

import "fmt"
import "math"

func main(){
	var primes [1000000] int;
	for i:=2; i < len(primes); i++{
		if(primes[i] == 0){
			for j:=i; j < len(primes); j+=i{
				primes[j]=i
			}
		}
	}
	var n int;
	fmt.Scanf("%d",&n)
	/*if n < len(primes){
		fmt.Println(primes[n])
	}*/
	var max_prime_divisor int = n
	for n%2 == 0{
		max_prime_divisor = 2
		n/=2
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i+=2{
		if n%i == 0{
			max_prime_divisor = i;
			for n%i == 0{
				n/=i
			}
		}
	}
	if n != 1{
		max_prime_divisor = n
	}
	fmt.Println(max_prime_divisor)
}
