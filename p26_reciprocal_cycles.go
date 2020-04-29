package main

import "fmt"

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}

func main(){
	cycle := 0
	num := 0
	for d := 1; d < 1000; d++{
		remainder := make([]int, 0)
		n := 1
		for true{
			cnt := 0
			for n < d{
				n*=10
				cnt++
			}
			cnt--
			for ;cnt > 0 ; cnt--{
				remainder = append(remainder, 0)
			}
			new_rem := n%d
			if new_rem == 0{
				break
			}
			f := false
			for i, rem := range remainder{
				if rem == new_rem{
					cycle = max(cycle, len(remainder)-i)
					//fmt.Println(d,len(remainder)-i)
					if cycle == len(remainder) - i{
						num = d
					}
					f = true
					break
				}
			}
			if f{
				break
			}
			remainder = append(remainder, new_rem)
			n = new_rem
		}
	}
	fmt.Println(cycle, "for the number", num)
}
