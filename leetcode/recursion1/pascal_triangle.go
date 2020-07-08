package recursion1

func getRow(rowIndex int) []int {
	return getRowRecursive(rowIndex)
}

func getRowRecursive(rowNum int) []int {
	if rowNum == 0 {
		return []int{1}
	}
	prev := getRowRecursive(rowNum - 1)
	cur := make([]int, rowNum+1)

	for i := 0; i < len(cur); i++ {
		if i == 0 || i == len(cur)-1 {
			cur[i] = 1
		} else {
			cur[i] = prev[i-1] + prev[i]
		}
	}
	return cur
}
