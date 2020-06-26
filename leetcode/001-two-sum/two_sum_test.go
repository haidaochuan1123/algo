package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := [][]int{
		[]int{3, 4, 6},
		[]int{3, 4, 6},
		[]int{2, 7, 11, 15},
		[]int{1, 2, 3, 4, 5},
		[]int{3},
	}

	targets := []int{
		6,
		7,
		9,
		5,
		3,
	}

	results := [][]int{
		[]int{1, 2},
		[]int{0, 1},
		[]int{0, 1},
		[]int{1, 2},
		[]int{0, 1},
	}

	for idx, target := range targets {
		ret := twoSum(tests[idx], target)
		fmt.Printf("nums = %v target = %v result = %v\n",
			tests[idx],
			target,
			ret,
		)

		if ret != nil && ret[0] != results[idx][0] && ret[1] != results[idx][1] {
			t.Errorf("case %d fails: %v\n", idx, ret)
		}

	}
}
