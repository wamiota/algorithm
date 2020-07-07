package recursion1

func reverseString(s []byte) {
	recursiveReverseString(s, 0, len(s)-1)
	twoPointersReverseString(s)
}

func recursiveReverseString(s []byte, left int, right int) {
	if left >= right {
		return
	}

	temp := s[left]
	s[left] = s[right]
	s[right] = temp

	recursiveReverseString(s, left+1, right-1)
}

func twoPointersReverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
	}
}
