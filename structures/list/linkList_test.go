package list

import (
	"testing"
)

var list LinkedList

func TestLindedList(t *testing.T) {
	intArray := []interface{}{1, 2, 3, 4, 5}
	strArray := []string{"a", "b", "c", "d"}

	linkedList := UnmarshalListBySlice(intArray)
	linkedList.String()
	linkedList.RemoveNode(1)
	linkedList.String()

	// // Traverse(linkedList).PrintList()

	for _, v := range strArray {
		linkedList.Insert(v, 0)
	}

	linkedList.String()
}

func TestRemove(t *testing.T) {
	if !list.IsEmpty() {
		t.Errorf("Linked list should be empty")
	}

	list.Append("first")

	if size := list.Len(); size != 1 {
		t.Errorf("Wrong count, expected 1 but got %d", size)
	}
	list.Append("two")
	list.Append("three")
	list.Insert("zero", 0)

	list.RemoveNode("two")
	list.String()

	if size := list.Len(); size != 3 {
		t.Errorf("Wrong count, expected 3 but got %d", size)
	}

}

func TestTraverse(t *testing.T) {
	linkedList := UnmarshalListByRand(100, 10)
	linkedList.String()

	traversedList := linkedList.Traverse()
	traversedList.String()
	traversedDoubleList := linkedList.Traverse().Traverse()

	linkedNode, traversedNode := linkedList.header, traversedDoubleList.header
	for i := 0; i < linkedList.size; i++ {
		if linkedNode.Value != traversedNode.Value {
			t.Errorf("Wrong Traverse")
		}
	}
}
