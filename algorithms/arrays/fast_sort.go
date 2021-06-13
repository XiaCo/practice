package arrays

func fastSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	baseNum := nums[0]
	var left = 0
	var right = len(nums) - 1
	for i := 1; i <= right; {
		if nums[i] > baseNum {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		} else {
			nums[left], nums[i] = nums[i], nums[left]
			left++
			i++
		}

	}
	fastSort(nums[:left])
	fastSort(nums[left+1:])
}
