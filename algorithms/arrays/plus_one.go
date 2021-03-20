package arrays

import (
	"math"
)

// https://leetcode-cn.com/problems/plus-one/

func PlusOne(digits []int) []int {
	var result = make([]int, len(digits)+1)
	var plus bool = true
	for i := len(digits) - 1; i >= 0; i-- {
		var num int
		if plus {
			num = digits[i] + 1
		} else {
			num = digits[i]
		}
		if num == 10 {
			plus = true
			result[i+1] = 0
		} else {
			plus = false
			result[i+1] = num
		}
	}
	if plus {
		result[0] = 1
		return result
	}
	return result[1:]
}

// 位数大的时候，无法使用此方法
func PlusOneByAdd(digits []int) []int {
	var sum int
	var length = len(digits) - 1
	for i := length; i >= 0; i-- {
		sum += digits[i] * int(math.Pow10(length-i))
	}

	sum += 1

	// 考虑进位
	var result = make([]int, len(digits)+1)
	index := len(digits)
	for sum != 0 {
		num := sum % 10
		result[index] = num
		sum = (sum - num) / 10
		index--
	}
	return result[index+1:]
}
