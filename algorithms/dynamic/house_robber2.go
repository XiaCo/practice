package dynamic

// https://leetcode-cn.com/problems/house-robber-ii/

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	n1 := _rob2(nums[1:])
	n2 := _rob2(nums[:len(nums)-1])
	return maxNum(n1, n2)
}

func _rob2(nums []int) int {
	dp := make([][2]int, len(nums))
	dp[0][0] = nums[0]
	dp[0][1] = 0
	for i := 1; i < len(nums); i++ {
		dp[i][0] = dp[i-1][1] + nums[i]
		dp[i][1] = maxNum(dp[i-1][0], dp[i-1][1])
	}
	index := len(nums) - 1
	return maxNum(dp[index][0], dp[index][1])
}
