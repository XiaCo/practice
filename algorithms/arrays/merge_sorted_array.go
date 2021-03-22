package arrays

// https://leetcode-cn.com/problems/merge-sorted-array/

func Merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n - 1; i >= 0; i-- {
		if m == 0 && n != 0 {
			copy(nums1[:n], nums2[:n])
			return
		} else if n == 0 && m != 0 {
			return
		}
		if nums1[m-1] > nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
	}
}
