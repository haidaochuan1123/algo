package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	L1 := []int{1, 2, 3, 4}
	L2 := []int{1, 2, 3, 4}
	l1 := NewListNode(L1)
	l2 := NewListNode(L2)

	result := addTwoNumbers(l1, l2)
	for result.Next != nil {
		fmt.Printf("%d -> ", result.Val)
		result = result.Next
	}
	fmt.Printf("%d \n", result.Val)
}
