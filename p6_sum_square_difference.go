package main

import "fmt"

func main(){
	var n int64
	fmt.Scanf("%d",&n)
	var sum1 , sum2 int64 = 0,0
	for  i := int64(1); i <= n; i++{
		sum1+=i*i
		sum2+=i
	}
	fmt.Println(sum2*sum2-sum1)
}
