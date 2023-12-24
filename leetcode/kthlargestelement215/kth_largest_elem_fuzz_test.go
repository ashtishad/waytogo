package kthlargestelement215

import (
	"sort"
	"testing"
	"testing/quick"
)

func fuzzTest(fn func([]int, int) int) func([]int, int) bool {
	return func(nums []int, k int) bool {
		// Ensure k is within the valid range
		if k < 1 || k > len(nums) {
			return true
		}

		// Copy the input to avoid modifying the original slice
		input := make([]int, len(nums))
		copy(input, nums)

		// Sort the copy using the standard library's sort package
		sort.Ints(input)

		// Get the expected result by finding the kth largest element in the sorted copy
		expected := input[len(nums)-k]

		// Run the tested function
		result := fn(nums, k)

		// Compare the result with the expected value
		return result == expected
	}
}

func TestFuzzFindKthLargest(t *testing.T) {
	f := fuzzTest(findKthLargest)
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
