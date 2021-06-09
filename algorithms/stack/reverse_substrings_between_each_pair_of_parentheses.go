package stack

// https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses/

// 输入：s = "(ed(et(oc))el)"
// 输出："leetcode"
// 解释：先反转子字符串 "oc" ，接着反转 "etco" ，然后反转整个字符串。
func reverseParentheses(s string) string {
	sub := [][]byte{{}}
	for index := range s {
		b := s[index]
		if b == '(' {
			sub = append(sub, make([]byte, 0))
		} else if b == ')' {
			tailSub := sub[len(sub)-1]
			sub = sub[:len(sub)-1]
			reverseBytes(tailSub)
			sub[len(sub)-1] = append(sub[len(sub)-1], tailSub...)
		} else {
			sub[len(sub)-1] = append(sub[len(sub)-1], b)
		}
	}
	return string(sub[0])
}

func reverseBytes(buf []byte) {
	for i := 0; i < len(buf)/2; i++ {
		buf[i], buf[len(buf)-1-i] = buf[len(buf)-1-i], buf[i]
	}
}
