package kthlargestelement215

import "slices"

// Approach 3: Counting Sort
// Time: O(n)
// Space: O(n)

func findKthLargest(nums []int, k int) int {
	min, max := 10000, -10000 // from constraints: -104 <= nums[i] <= 104

	min, max = slices.Min(nums), slices.Max(nums)

	n := max - min + 1

	counter := make([]int, n)

	for _, num := range nums {
		counter[num-min]++
	}

	for i := n - 1; i >= 0; i-- {
		k -= counter[i]
		if k <= 0 {
			return i + min
		}
	}

	return -1
}

// Algo
// It employs counting sort to count the frequency of each element in the array.
// The algorithm then iterates through the counting array from the highest index to the lowest,
// decrementing k until it becomes less than or equal to 0, and returns the corresponding value in the original array.
