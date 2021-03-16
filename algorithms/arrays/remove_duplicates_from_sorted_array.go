package arrays

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
// 双指针

func RemoveDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var indexSlow, indexFast int
	indexSlow = 1
	indexFast = 1
	for indexFast < len(nums) {
		nums[indexSlow] = nums[indexFast]
		indexFast++
		if nums[indexSlow] != nums[indexSlow-1] {
			indexSlow++
		}
	}
	return indexSlow
}
