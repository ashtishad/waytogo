package kthlargestelement215

import "slices"

// Approach 1: Sorting(quick sort and return the element of idx = len - k
// Time: O(nlogn) and Space O(1)

/*

Given an integer array nums and an integer k, return the kth largest element in the array.
Note that it is the kth largest element in the sorted order, not the kth distinct element.
Which means whatever element is there in kth largest position.

Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Sorted: nums = [1,2,2,3,3,4,5,6,6], k = 4, len= 9, rIdx = 5
Output: 4

*/

func findKthLargestA1(nums []int, k int) int {
	slices.Sort(nums)

	idx := len(nums) - k

	return nums[idx]
}
