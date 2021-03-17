package arrays

// https://leetcode-cn.com/problems/rotate-array/

func RotateByCopy(nums []int, k int) {
	mirror := make([]int, len(nums))
	copy(mirror, nums)
	deviation := len(nums) - k
	for index := range nums { // 将deviation替换为k就是向左偏移
		indexMirror := (deviation + index) % len(nums)
		nums[index] = mirror[indexMirror]
	}
}

func swap(nums []int) {
	mid := len(nums) / 2
	for i := 0; i < mid; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}

func RotateBySwap(nums []int, k int) {
	// deviation := len(nums) - k
	// 将下面的k改为deviation就是向左偏移
	swap(nums)
	swap(nums[k:])
	swap(nums[:k])
}
