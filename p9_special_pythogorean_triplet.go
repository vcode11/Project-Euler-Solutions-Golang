package main

import "fmt"

func main(){
	for i:=1; i < 1000; i++{
		for j:= i+1; j < 1000; j++{
			k := 1000 - i - j
			if k > 0{
				if i*i + j*j == k*k{
					fmt.Println(i,j,k,i*j*k)
				}
			}
		}
	}
}
