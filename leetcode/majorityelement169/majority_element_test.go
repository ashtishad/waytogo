package majorityelement169

import "testing"

// appears more than ⌊n / 2⌋ times
func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		output int
	}{
		{"Example 1", []int{3, 3, 2, 3, 4, 3, 2, 4, 3, 4, 4, 4, 4, 7, 3, 4, 4, 4, 4, 4}, 4},
		{"Example 2", []int{2, 2, 1, 1, 1, 2, 2}, 2},
		{"Example 3", []int{7}, 7},
		{"Example 4", []int{2, 8, 7, 2, 2, 5, 2, 3, 1, 2, 2}, 2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := majorityElement(test.nums)
			if result != test.output {
				t.Errorf("Expected %d, but got %d", test.output, result)
			}
		})
	}
}
