package main

import "fmt"

func main(){
	var a, b [1001] int
	var  len_b int  = 1
	a[0] = 0
	b[0] = 1
	i := 1
	for len_b != 1000{
		c := b
		carry := 0
		for k := 0; k < len_b; k++{
			b[k]+=a[k]+carry
			carry = (b[k]-b[k]%10)/10
			b[k]%=10
		}
		for carry != 0{
			b[len_b] = carry%10;
			carry = (carry-carry%10)/10
			len_b++
		}
		a = c
		i++
	}
	fmt.Println(i,"th fibonacci number has 1000 digits.")
}
