package besttimetobuyandsellastock121

// brute-force
// Time: O(n^2) and Space: O(1)
// Leetcode -> Time Limit Exceeded

// Input: prices = [7,1,5,3,6,4]
// Output: 5 (buy on day 2 and sell on day 5)

func maxProfitIBF(prices []int) int {
	var mp int

	n := len(prices)

	if n == 1 {
		return mp
	}

	for i := 0; i < n-1; i++ { // i can buy till day before last day
		for j := i + 1; j < n; j++ {
			cp := prices[j] - prices[i] // current profit = selling - buying price

			if cp > mp {
				mp = cp
			}
		}
	}

	return mp
}
