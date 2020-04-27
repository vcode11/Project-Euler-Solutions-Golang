package main

import "fmt"
import "strconv"

func next_permutation(perm []int) bool{
	j:=len(perm)-1;
	f:=false
	for j > 0 {
		if perm[j] > perm[j-1]{
			f = true
			break
		}
		j--
	}
	if f {
		k:=len(perm)-1
		for ; k >= 0; k--{
			if perm[k] > perm[j-1]{
				break
			}
		}
		tmp:=perm[j-1]
		perm[j-1] = perm[k]
		perm[k] = tmp
		k = len(perm)-1
		for i := j; i < k; i++{
			tmp:=perm[i]
			perm[i] = perm[k]
			perm[k] = tmp
			k--;
		}
	}
	return f
}
func main(){
	var n, size int
	fmt.Scanf("%d %d", &n, &size)
	perm := make([]int, size)
	for i := 0; i < size; i++{
		fmt.Scanf("%d",&perm[i])
	}
	for i:=1; i < n; i++{
		next_permutation(perm);
	}
	for i := 0; i < size; i++{
		fmt.Printf(strconv.Itoa(perm[i]))
	}
	fmt.Printf("\n")
}
