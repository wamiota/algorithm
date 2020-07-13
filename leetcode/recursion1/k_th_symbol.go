package recursion1

func kthGrammar(N int, K int) int {
	if K == 1 {
		return 0
	}
	formar := kthGrammar(N, (K+1)/2)
	if formar == 0 {
		return (K + 1) % 2
	} else {
		return K % 2
	}
}
