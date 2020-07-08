package recursion1

func fib(N int) int {
	if N == 0 {
		return 0
	}
	memo := make([]int, N+1)
	memo[0], memo[1] = 0, 1

	return fibMemorization(N, memo)
}

func fibMemorization(i int, memo []int) int {
	if memo[i] != 0 || i == 0 {
		return memo[i]
	}
	return fibMemorization(i-1, memo) + fibMemorization(i-2, memo)
}
