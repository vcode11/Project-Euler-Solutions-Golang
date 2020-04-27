package main

import (
	"fmt"
)
func gcd(a int, b int) int{
	for b != 0{
		t:=b
		b = a%b
		a = t
	}
	return a
}
func Lcm(a int, b int) int{
	return (a/gcd(a,b))*b
}
func main(){
	lcm := 1
	for i := 1; i < 21; i++{
		lcm = Lcm(lcm,i)
	}
	fmt.Println(lcm)

}
