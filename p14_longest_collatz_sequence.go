package main

func collatz(n int64, dp []int64) int64{
	if n < 0{
		println(n)
		return 0
	}
	if int(n) < len(dp){
		if dp[n] != 0 {
			return dp[n]
		} else if n%2 == 0{
			dp[n] = collatz(n/2, dp)+1
		} else{
			dp[n] = collatz(3*n+1, dp) + 1
		}
		return dp[n]
	} else {
		if n%2 == 0{
			return collatz(n/2, dp)+1
		} else {
			return collatz(3*n+1, dp)+1
		}
	}
}

func main(){
	dp := make([]int64,10)
	dp[1] = 1
	m, n  := int64(0), int64(1)
	for i := int64(2); i < 1000001; i++{
		val:=collatz(i,dp)
		if val > m{
			m = val
			n = i
		}
	}
	println("Longest Collatz length is ", m, "For the value ", n)
}
