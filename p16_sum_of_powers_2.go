package main

import "fmt"

func max(a , b int) int{
	if a > b{
		return a
	}
	return b
}
func main(){
	var num [400]int
	num[0] = 1
	var n int
	fmt.Scanf("%d", &n)
	num_len := 1
	for i := 0; i < n; i++{
		carry:=0
		j:=0
		for j < 400 && j < num_len{
			num[j]*=2
			num[j]+=carry
			carry=0
			carry+=(num[j]-num[j]%10)/10
			num[j]%=10
			j++
		}
		for carry > 0 {
			num[j] = carry%10
			carry = (carry-carry%10)/10
			j++
		}
		num_len = max(num_len, j)
	}
	sum:=0
	for i:= num_len-1; i >= 0; i--{
		sum+=num[i]
		fmt.Printf(string(num[i]+48))
	}
	fmt.Println("\n", "Sum of digits is ", sum)
}
