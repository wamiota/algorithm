package recursion2

func sortArray(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	pivod := len(nums) / 2

	left := mergeSort(nums[0:pivod])
	right := mergeSort(nums[pivod:len(nums)])

	return merge(left, right)
}

func merge(leftNums []int, rightNums []int) []int {
	res := make([]int, len(leftNums)+len(rightNums))
	leftCur, rightCur, resCur := 0, 0, 0

	for leftCur < len(leftNums) && rightCur < len(rightNums) {
		if leftNums[leftCur] < rightNums[rightCur] {
			res[resCur] = leftNums[leftCur]
			resCur += 1
			leftCur += 1
		} else {
			res[resCur] = rightNums[rightCur]
			resCur += 1
			rightCur += 1
		}
	}

	for leftCur < len(leftNums) {
		res[resCur] = leftNums[leftCur]
		resCur += 1
		leftCur += 1
	}

	for rightCur < len(rightNums) {
		res[resCur] = rightNums[rightCur]
		resCur += 1
		rightCur += 1
	}

	return res
}
