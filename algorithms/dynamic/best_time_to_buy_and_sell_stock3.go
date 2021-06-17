package dynamic

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/

// 输入：prices = [3,3,5,0,0,3,1,4]
// 输出：6
// 到结束，可能有5种状态:
// dp1. 没有交易过
// dp2. 买过1次
// dp3. 买卖过1次
// dp4. 买卖过1次，又买了1次
// dp5. 买卖过2次
// 求5种状态的最大值，即：
// dp1 = 0
// dp2 = max{-prices[index], dp2}
// dp3 = max{dp3, dp2+prices[index]}
// dp4 = max{dp4, dp3-prices[index]}
// dp5 = max{dp5, dp4+prices[index]}
func maxProfit3DP(prices []int) int {
	dp1 := 0
	dp2 := -prices[0]
	dp3, dp4, dp5 := -100000, -100000, -100000
	for i := 1; i < len(prices); i++ {
		dp2 = maxNum(dp2, dp1-prices[i])
		dp3 = maxNum(dp3, dp2+prices[i])
		dp4 = maxNum(dp4, dp3-prices[i])
		dp5 = maxNum(dp5, dp4+prices[i])
	}
	return maxNum(dp1, dp5)
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
