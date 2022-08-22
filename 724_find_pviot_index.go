func calculateRightSum(nums []int) []int {
	res := make([]int, len(nums))
	copy(res, nums)
	for i := 1; i < len(nums); i++ {
		res[i] = res[i] + res[i-1]
	}

	return res
}

func calculateLeftSum(nums []int) []int {
    res := make([]int, len(nums))
	copy(res, nums)
	for i := len(nums) - 2; i >= 0; i-- {
		res[i] = res[i] + res[i+1]
	}
	return res
}

func pivotIndex(nums []int) int {
    if len(nums) == 1 {
        return 0
    }
	rightSum := calculateRightSum(nums)
	leftSum := calculateLeftSum(nums)
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			if leftSum[i+1] == 0 {
				return i
			}
			continue
		}
		if i == len(nums)-1 {
			if rightSum[i-1] == 0 {
				return i
			}
			break
		}
		if rightSum[i-1] == leftSum[i+1] {
			return i
		}
	}
	return -1
}
