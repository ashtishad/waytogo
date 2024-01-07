package twosum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		result []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{0, 4, 3, 0}, 0, []int{0, 3}},                // Test case with zero values.
		{[]int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},       // Test case with negative values.
		{[]int{1, 2, 3, 4, 5}, 9, []int{3, 4}},             // Test case with positive values.
		{[]int{2, 2, 2, 2}, 4, []int{0, 1}},                // Test case with repeated values.
		{[]int{0, 1, 2, 3, 4, 5}, 9, []int{4, 5}},          // Test case with ascending values.
		{[]int{5, 4, 3, 2, 1, 0}, 9, []int{0, 1}},          // Test case with descending values.
		{[]int{100, 200, 300, 400, 500}, 800, []int{2, 4}}, // Test case with large numbers.
	}

	for _, test := range tests {
		result := twoSumHashMap(test.nums, test.target)
		if !reflect.DeepEqual(result, test.result) {
			t.Errorf("For nums=%v, target=%d, expected %v but got %v", test.nums, test.target, test.result, result)
		}
	}
}

func TestTwoSumSorted(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		result []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{0, 0, 3, 4}, 0, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
		{[]int{-5, -4, -3, -2, -1}, -8, []int{1, 3}},
		{[]int{1, 2, 3, 4, 5}, 9, []int{4, 5}},
		{[]int{100, 200, 300, 400, 500}, 800, []int{3, 5}},
	}

	for _, test := range tests {
		result := twoSumBinarySearch(test.nums, test.target)
		if !reflect.DeepEqual(result, test.result) {
			t.Errorf("For nums=%v, target=%d, expected %v but got %v", test.nums, test.target, test.result, result)
		}
	}
}
