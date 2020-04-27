package main

import "fmt"
import "strconv"

func main(){
	var n int64 = 1000000
	var factorial [11] int64
	factorial[0] = 1
	for i:=int64(1); i < 11; i++{
		factorial[i] = i*factorial[i-1]
	}
	digits := map[int]bool{
		0:true,
		1:true,
		2:true,
		3:true,
		4:true,
		5:true,
		6:true,
		7:true,
		8:true,
		9:true,
	}
	var perm [10]int
	for i := 0; i < 10; i++{
		ways_to_permute:=factorial[10-i-1]
		k := int64(1)
		for j := 0; j < 10; j++{
			if(digits[j]){
				if k*ways_to_permute >= n{
					perm[i] = j
					n-=(k-1)*ways_to_permute
					digits[j] = false
					break
				}
				k++
			}
		}
	}
	for i := 0; i < 10; i++{
		fmt.Printf(strconv.Itoa(perm[i]))
	}
	fmt.Printf("\n")
}
