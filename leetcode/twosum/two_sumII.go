package twosum

import "slices"

// 167. Two Sum II - Input Array Is Sorted
// Time: O(nlogn) Space: O(1)
// Exactly one solution, index1 < index2, added by one

// Input: numbers = [2,7,11,15], target = 9. Output: [1,2]
// Caution: {[]int{0, 0, 3, 4}, 0, []int{1, 2}}

func twoSumBinarySearch(nums []int, target int) []int {
	for i := range nums {
		idx, ok := slices.BinarySearch(nums, target-nums[i])

		if ok {
			switch {
			case i == idx:
				continue
			case i > idx:
				return []int{idx + 1, i + 1}
			default:
				return []int{i + 1, idx + 1}
			}
		}
	}

	return []int{-1, -1}
}
