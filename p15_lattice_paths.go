package main

import "fmt"

func main(){
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	var dp[1000][1000]int
	for i:=0; i <= m; i++{
		dp[0][i] = 1
	}
	for i:=0; i <= n; i++{
		dp[i][0] = 1
	}
	dp[0][0] = 0
	for i := 1; i <= n; i++{
		for j:=1; j <= m; j++{
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	fmt.Println(dp[n][m])
}
