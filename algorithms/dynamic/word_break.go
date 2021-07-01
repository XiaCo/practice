package dynamic

// https://leetcode-cn.com/problems/word-break/

// dp[i] 表示s[0,i]是否可以被分割成全部在字典中的子串
// dp[i] == dp[j] && s[j:i] in wordDict
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	m := make(map[string]bool, len(wordDict))
	for _, item := range wordDict {
		m[item] = true
	}
	dp[0] = true
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			val := m[s[j:i]]
			dp[i] = dp[j] && val
			if dp[i] {
				break
			}
		}
	}
	return dp[len(s)]
}
