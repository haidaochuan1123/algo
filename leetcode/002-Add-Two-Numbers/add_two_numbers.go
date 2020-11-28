package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(input []int) *ListNode {
	listLen := len(input)
	var curNode *ListNode
	for i := listLen; i >= 1; i-- {
		node := &ListNode{
			Val:  input[i-1],
			Next: curNode,
		}

		curNode = node
	}

	return curNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{Val: 0}
	n1, n2, carry, curNode := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		// 处理l1、l2 长度不一致问题
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}

		// 处理相加进1问题
		curNode.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		curNode = curNode.Next
		carry = (n1 + n2 + carry) / 10
	}

	return head.Next
}
