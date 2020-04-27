package main

import "time"
import "fmt"

func max_power(a, b int) int{
	cnt:=0
	for a%b == 0{
		cnt++
		a/=b
	}
	return cnt
}
func count_divisors(a int) int{
	cnt:=2
	for i := 2; i*i <= a; i++{
		if a%i == 0{
			cnt+=2
		}
		if i*i == a{
			cnt--
		}
	}
	return cnt
}
func main(){
	start:=time.Now()
	/*var divisors [10000000]int
	divisors[1] = 1
	for i:=2; i < len(divisors); i++{
		if divisors[i] == 0 {
			for j:=2*i; j < len(divisors); j+=i{
				if divisors[j] == 0{
					divisors[j] = 1
				}
				divisors[j]*=int(max_power(j,i)+1)
			}
			divisors[i] = 2
		}
	}*/
	for i:=1; i < 1000000; i++{
		num:=(i*(i+1))/2
		//if num <  len(divisors) && divisors[num] >= 500{
		//	println((i*(i+1))/2)
		//	break
		//} 
		if count_divisors(num) >= 500{
			println(num, i)
			break
		}
	}
	fmt.Println("Executed in", time.Since(start))
}
