//sum of digits of 100 factorial
package main

import "fmt"

func main(){
	var num [200] int
	num[0] = 1
	length := 1
	var n int
	fmt.Scanf("%d", &n)
	for i:=2; i <= n; i++{
		carry := 0
		for j:=0; j < length; j++{
			num[j]*=i
			num[j]+=carry
			carry=(num[j]-num[j]%10)/10
			num[j]%=10
		}
		for carry > 0{
			num[length] = carry%10
			carry = (carry-carry%10)/10
			length++
		}
	}
	sum := 0
	for i:= length-1; i >=0; i--{
		fmt.Printf(string(num[i]+48))
		sum += num[i]
	}
	fmt.Printf("\n")
	fmt.Println("Digit sum is", sum)
}
