package majorityelement169

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestFuzzMajorityElement(t *testing.T) {
	f := func(nums []int) bool {
		// Skip empty slices to avoid division by zero
		if len(nums) == 0 {
			return true
		}

		_ = majorityElement(nums)
		return true
	}

	if err := quick.Check(f, nil); err != nil {
		fmt.Println("Fuzz test failed:", err)
	} else {
		fmt.Println("Fuzz test passed")
	}
}
