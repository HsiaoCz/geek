package main

// 动态规划算法

// 递归算法实现的斐波那契数列
func Fb(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}
	return Fb(n-1) + Fb(n-2)
}

// 动态规划算法实现斐波那契数列

func Fib(n int) int {
	dp := make([]int, n+2)
	dp[0] = 0
	dp[1] = 1
	if n < 2 {
		return dp[n]
	}
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
