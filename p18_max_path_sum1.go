package main

import "fmt"

var n int
var dp [101][101] int
func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}
func max_path_sum(i, j int, arr [][]int) int{
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	if i == n-1 {
		return arr[i][j]
	}
	dp[i][j] = arr[i][j] + max(max_path_sum(i+1,j, arr), max_path_sum(i+1,j+1,arr))
	return dp[i][j]
}
func main(){
	var arr [][]int
	fmt.Println("Enter number of rows:")
	fmt.Scanf("%d",&n)
	for i := 0; i < n; i++{
		var slice []int
		for  j:=0; j <= i; j++{
			var x int
			fmt.Scanf("%d", &x)
			slice = append(slice, x)
			dp[i][j] = -1
		}
		arr = append(arr, slice)
	}
	fmt.Println(max_path_sum(0,0,arr))
}
