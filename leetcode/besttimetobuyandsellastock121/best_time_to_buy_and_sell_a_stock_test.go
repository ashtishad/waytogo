package besttimetobuyandsellastock121

import (
	"reflect"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		result int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{5}, 0},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2, 3, 4, 5}, 4},          // Increasing prices
		{[]int{8, 7, 6, 5, 4, 3, 2, 1}, 0}, // No profit, decreasing prices
		{[]int{2, 2, 2, 2, 2}, 0},          // No profit, constant prices
		{[]int{5, 1, 8, 6, 2, 10}, 9},      // Buy on day 2, sell on day 6
		{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 4}, // Buy on day 4, sell on day 7
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 7}, // Buy on day 0, sell on day 7
	}

	for _, test := range tests {
		result := maxProfit(test.prices)
		if !reflect.DeepEqual(result, test.result) {
			t.Errorf("For prices=%v, expected %v but got %v", test.prices, test.result, result)
		}
	}
}
