package arrays

import "sort"

// https://leetcode-cn.com/problems/4sum/

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}
	var result [][]int
	sort.Ints(nums)
	for index1, num1 := range nums {
		if index1 != 0 && nums[index1-1] == num1 {
			continue
		}
		threeSum := target - num1
		for index2 := index1 + 1; index2 < len(nums)-2; index2++ {
			num2 := nums[index2]
			if index2 != index1+1 && nums[index2-1] == num2 {
				continue
			}
			twoSum := threeSum - num2
			left := index2 + 1
			right := len(nums) - 1
			for left < right {
				if left != index2+1 && nums[left-1] == nums[left] {
					left++
					continue
				}
				if sum := nums[left] + nums[right]; sum < twoSum {
					left++
				} else if sum > twoSum {
					right--
				} else {
					result = append(result, append([]int{num1, num2, nums[left], nums[right]}))
					left++
				}
			}
		}
	}
	return result
}
