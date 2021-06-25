package dynamic

import "strconv"

// https://leetcode-cn.com/problems/decode-ways/

// 输入：s = "226"
// 输出：3
// 解释：它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6)

// s[i-1:i+1]
// [0] => return 0
// (0,9] => dp[i] = dp[i-1]
// 10,20 => dp[i] = dp[i-2]
// [11,19],[21,26] => dp[i] = dp[i-1] + dp[i-2]
// 30,40,50...90 => return 0
// [27,99] => dp[i] = dp[i-1]

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s))
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	dp[0] = 1
	num, _ := strconv.Atoi(s[0:2])
	if num >= 0 && num <= 9 {
		return 0
	} else if num == 10 || num == 20 {
		dp[1] = 1
	} else if (num >= 11 && num <= 19) || (num >= 21 && num <= 26) {
		dp[1] = 2
	} else if s[0]-'0' > 2 && s[1] == '0' {
		return 0
	} else {
		dp[1] = 1
	}

	for i := 2; i < len(s); i++ {
		num, _ := strconv.Atoi(s[i-1 : i+1])
		if num == 0 {
			return 0
		} else if num > 0 && num <= 9 {
			dp[i] = dp[i-1]
		} else if num == 10 || num == 20 {
			dp[i] = dp[i-2]
		} else if (num >= 11 && num <= 19) || (num >= 21 && num <= 26) {
			dp[i] = dp[i-1] + dp[i-2]
		} else if s[i-1]-'0' > 2 && s[i] == '0' {
			return 0
		} else {
			dp[i] = dp[i-1]
		}

	}
	return dp[len(s)-1]
}
