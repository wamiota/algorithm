package recursion1

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	memo := make([]int, n)
	memo[0] = 1
	memo[1] = 1
	return climbStairsMemorization(2, n, memo)
}

func climbStairsMemorization(i int, n int, memo []int) int {
	if i == n {
		return memo[n-1] + memo[n-2]
	}
	memo[i] = memo[i-1] + memo[i-2]
	return climbStairsMemorization(i+1, n, memo)
}
