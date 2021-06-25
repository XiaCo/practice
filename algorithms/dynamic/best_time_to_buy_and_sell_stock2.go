package dynamic

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

//给定一个数组 prices ，其中 prices[i] 是一支给定股票第 i 天的价格。
//设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

//输入: prices = [7,1,5,3,6,4]
//输出: 7
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
//     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3
func maxProfit2(prices []int) int {
	var result int
	for i := 1; i < len(prices); i++ {
		if num := prices[i] - prices[i-1]; num > 0 {
			result += num
		}
	}
	return result
}

// f[i][cash] => max{f[i-1][cash], f[i-1][stock] + prices[i]}
// f[i][stock] => max{f[i-1][stock], f[i-1][cash] - prices[i]}
func maxProfit2DP(prices []int) int {
	dp := make([][2]int, len(prices))
	// 0下标表cash状态
	// 1下标表stock状态
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = maxNum(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = maxNum(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}
