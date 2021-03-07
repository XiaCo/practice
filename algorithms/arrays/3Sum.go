package arrays

// https://leetcode-cn.com/problems/3sum/

// 返回三数和为0
func ThreeSum(nums []int) [][]int {
	result := make([][]int, 0)
	nums = QuickSort(nums)
	for index, num := range nums {
		if index != 0 && nums[index-1] == num { // 去重
			continue
		}
		otherNum := -num
		indexLeft := index + 1
		indexRight := len(nums) - 1
		for indexLeft < indexRight {
			numLeft := nums[indexLeft]
			numRight := nums[indexRight]
			if indexLeft != index+1 && nums[indexLeft-1] == numLeft { // 去重
				indexLeft++
				continue
			}
			if sum := numLeft + numRight; sum < otherNum {
				indexLeft++
			} else if sum > otherNum {
				indexRight--
			} else {
				result = append(result, []int{num, numLeft, numRight})
				indexLeft++
			}
		}
	}
	return result
}

// 返回两数和为13的组合
func TwoSum(nums []int) [][]int {
	// 排序
	// 双指针
	var result [][]int
	if len(nums) < 2 {
		return result
	}
	nums = QuickSort(nums)

	indexLeft := 0
	indexRight := len(nums) - 1
	for indexLeft < indexRight {
		numLeft := nums[indexLeft]
		numRight := nums[indexRight]
		if indexLeft != 0 && nums[indexLeft-1] == numLeft { // 去重
			indexLeft++
			continue
		}
		if sum := numLeft + numRight; sum < 13 {
			indexLeft++
		} else if sum > 13 {
			indexRight--
		} else {
			result = append(result, []int{numLeft, numRight})
			indexLeft++
		}
	}
	return result
}

// 快速排序
func QuickSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	data[head] = mid
	QuickSort(data[:head])
	QuickSort(data[head+1:])
	return data
}

// 使用map
func TwoSumByMap(nums []int) [][]int {
	m := make(map[int]int, len(nums))
	var result [][]int
	for index, num := range nums {
		if _, selfExist := m[num]; selfExist { // 去重
			continue
		}
		if _, otherExist := m[13-num]; otherExist {
			result = append(result, []int{num, 13 - num})
		}
		m[num] = index
	}
	return result
}
