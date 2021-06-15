package dynamic

// https://leetcode-cn.com/problems/longest-palindromic-substring/

// "babad" -> "bab"

// 暴力
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	left := 0
	right := 1
	for i := 0; i < len(s)-1; i++ {
		for j := len(s); j > i; j-- {
			if isStringPal(s[i:j]) && (j-i) > (right-left) {
				left = i
				right = j
				break
			}
		}
	}
	return s[left:right]
}

func isStringPal(s string) bool {
	if len(s) <= 1 {
		return true
	}
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// dp
// s[i:j] => s[i+1:j-1]
func longestPalindromeDP(s string) string {
	if len(s) <= 1 {
		return s
	}
	matrix := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		matrix[i] = make([]bool, len(s))
		matrix[i][i] = true
	}

	max, index := 1, 0
	for j := 1; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if j-i < 2 {
				matrix[i][j] = s[i] == s[j]
			} else {
				matrix[i][j] = matrix[i+1][j-1] && (s[i] == s[j])
			}
			curLen := j - i + 1
			if matrix[i][j] && max < curLen {
				max = curLen
				index = i
			}
		}
	}
	return s[index : max+index]
}
