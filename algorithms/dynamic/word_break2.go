package dynamic

import "strings"

// https://leetcode-cn.com/problems/word-break-ii/

func wordBreak2(s string, wordDict []string) []string {
	var res []string
	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if m[s[j:i]] && dp[j] {
				dp[i] = true
				break
			}
		}
	}

	if !dp[len(s)] {
		return nil
	}

	var find func(str string, words []string)
	find = func(str string, words []string) {
		if str == "" {
			res = append(res, strings.Join(words, " "))
			return
		}
		for i := 1; i < len(str)+1; i++ {
			if m[str[:i]] {
				find(str[i:], append(words, str[:i]))
			}
		}
	}
	find(s, nil)
	return res
}
