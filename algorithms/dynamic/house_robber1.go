package dynamic

// https://leetcode-cn.com/problems/house-robber/

// dp[i][state]
// nums[-1] todaySteal
// dp[i][todaySteal] = dp[i-1][untodaySteal]+nums[i]
// nums[-1] untodaySteal
// d[i][untodaySteal] = max{dp[i-1][todaySteal], dp[i-1][untodaySteal]}
func rob(nums []int) int {
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
