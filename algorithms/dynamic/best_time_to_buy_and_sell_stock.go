package dynamic

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

// 输入：[7,1,5,3,6,4]
// 输出：5
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//      注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
func maxProfit(prices []int) int {
	var minIndex int
	var maxEarning int
	for index := range prices {
		if prices[index]-prices[minIndex] > maxEarning {
			maxEarning = prices[index] - prices[minIndex]
		}
		if prices[minIndex] > prices[index] {
			minIndex = index
		}
	}
	return maxEarning
}
