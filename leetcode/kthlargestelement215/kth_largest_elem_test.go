package kthlargestelement215

import "testing"

func TestFindKthLargest(t *testing.T) {
	testCases := []struct {
		nums   []int
		k      int
		output int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		// Add more test cases
		{[]int{5, 4, 3, 2, 1}, 2, 4},
		{[]int{1, 1, 1, 1, 1}, 1, 1},
		{[]int{1, 2, 3, 4, 5}, 3, 3},
		{[]int{-1, -2, -3, -4, -5}, 1, -1},
		{[]int{9, 8, 7, 6, 5}, 5, 5},
		{[]int{100, 200, 300, 400, 500}, 2, 400},
		{[]int{10, 20, 30, 40, 50}, 1, 50},
		{[]int{-5, -4, -3, -2, -1}, 3, -3},
	}

	for _, tc := range testCases {
		result := findKthLargest(tc.nums, tc.k)
		if result != tc.output {
			t.Errorf("Expected %d, but got %d for nums %v and k %d", tc.output, result, tc.nums, tc.k)
		}
	}
}
