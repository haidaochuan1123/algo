package list

import "testing"

func TestLindedList(t *testing.T) {
	intArray := []interface{}{1, 2, 3, 4, 5}
	strArray := []string{"a", "b", "c", "d"}

	linkedList := UnmarshalListBySlice(intArray)

	linkedList.RemoveNode(1)
	linkedList.PrintList()

	// Traverse(linkedList).PrintList()

	for _, v := range strArray {
		linkedList.Prepend(NewNode(v))
	}

	linkedList.PrintList()
}
