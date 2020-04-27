package main

import "fmt"

func in_words(n int, number map[int]string) string{
	s:=""
	if w, ok := number[n]; ok{
		return w
	}
	if n/100 > 0 {
		s+=number[(n/100)]+"hundred"
		n %= 100
		if(n > 0) {
			s+="and"
			s+=in_words(n,number)
		}
	} else if n/10 > 0{
		s+=number[n-n%10]+in_words(n%10,number)
	}
	return s

}
func main(){
	number := map[int]string{
		1:"one",
		2:"two",
		3:"three",
		4:"four",
		5:"five",
		6:"six",
		7:"seven",
		8:"eight",
		9:"nine",
		10:"ten",
		11:"eleven",
		12:"twelve",
		13:"thirteen",
		14:"fourteen",
		15:"fifteen",
		16:"sixteen",
		17:"seventeen",
		18:"eighteen",
		19:"nineteen",
		20:"twenty",
		30:"thirty",
		40:"forty",
		50:"fifty",
		60:"sixty",
		70:"seventy",
		80:"eighty",
		90:"ninety",
		1000:"onethousand",
	}
	/*for true{
		var n int
		fmt.Scanf("%d", &n)
		fmt.Println(in_words(n,number))
	}*/
	cnt:=0
	for i := 1; i <= 1000; i++{
		s:=in_words(i,number)
		cnt+=len(s)
		fmt.Println(s)
	}
	fmt.Println(cnt)
}
