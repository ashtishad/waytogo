package containsduplicate217

import (
	"testing"
)

// Given an integer array nums, return true if any value appears "at least twice" in the array,
// and return false if every element is distinct.

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums   []int
		result bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{[]int{}, false},    // Edge case: empty array
		{[]int{1}, false},   // Edge case: single element array
		{[]int{1, 1}, true}, // Edge case: duplicate elements
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := containsDuplicate(test.nums)
			if result != test.result {
				t.Errorf("Expected %v, but got %v for input %v", test.result, result, test.nums)
			}
		})
	}
}
