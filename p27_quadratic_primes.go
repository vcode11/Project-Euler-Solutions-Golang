package main

import "fmt"

func main(){
	var primes [1000005] int
	for i := 2; i < len(primes); i++{
		for j := 2*i; j < len(primes); j+=i{
			primes[j] = 1
		}
	}
	primes[0] = 1
	primes[1] = 1
	var length, best_a, best_b int = 0, -1, -1
	for a:= -999; a < 1000; a++{
		for b := -1000; b <= 1000; b++{
			for n := 0; n*n+a*n+b > 1 && primes[n*n + a*n + b] == 0;n++{
				if n > length{
					length = n
					best_a = a
					best_b = b
				}
				//fmt.Println(a,b,n)
			}
		}
	}
	fmt.Printf("n^2+%dn+%d produces %d consecutive primes\n", best_a, best_b, length)
}
