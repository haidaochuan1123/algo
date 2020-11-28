package list

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item 通用类型数据
type Item generic.Type

// Node 链表结构体
type Node struct {
	Value Item
	Next  *Node
}

// LinkedList 链表
type LinkedList struct {
	header *Node
	size   int
	lock   sync.RWMutex
}

// NewLinkedNode 返回链表结构体
func NewLinkedNode(value Item) *Node {
	return &Node{
		Value: value,
	}
}

// NewLinkedNodeWithNext 返回链表结构体
func NewLinkedNodeWithNext(value Item, next *Node) *Node {
	return &Node{
		Value: value,
		Next:  next,
	}
}

// IsEmpty 判断单链表是否为空
func (l *LinkedList) IsEmpty() bool {
	return l.header == nil
}

// Len 返回单链表长度
func (l *LinkedList) Len() int {
	return l.size
}

// Append 向链表最后添加元素
func (l *LinkedList) Append(value Item) {
	l.lock.Lock()
	defer l.lock.Unlock()
	newNode := NewLinkedNode(value)

	curNode := l.header

	// 空链表
	if curNode == nil {
		l.header = newNode
		l.size++
		return
	}

	for {
		if curNode.Next == nil {
			break
		}

		curNode = curNode.Next
	}

	curNode.Next = newNode
	l.size++

}

// Insert 向链表指定位置添加元素
func (l *LinkedList) Insert(value Item, index int) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	if index < 0 || index > l.Len() {
		return fmt.Errorf("Index %d out of bonuds", index)
	}

	newNode := NewLinkedNode(value)

	// 向链表头插入元素
	if index == 0 {
		newNode.Next = l.header
		l.header = newNode
		l.size++
		return nil
	}

	// 找到需要插入的位置的前一个元素
	preNode := l.header
	preIndex := 0
	for preIndex < index-2 {
		preIndex++
		preNode = preNode.Next
	}

	// 在查找的位置执行插入
	newNode.Next = preNode.Next
	preNode.Next = newNode
	l.size++
	return nil
}

// Find 从链表中查找指定的元素的位置
func (l *LinkedList) Find(value Item) (int, bool) {
	// l.lock.RLock()
	// defer l.lock.RUnlock()

	curNode := l.header
	locIndex := 0
	for {
		if curNode.Value == value {
			return locIndex, true
		}
		if curNode.Next == nil {
			return -1, false
		}

		curNode = curNode.Next
		locIndex++
	}
}

// RemoveNode 删除单链表节点
func (l *LinkedList) RemoveNode(value Item) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	curNode := l.header
	if curNode == nil {
		return fmt.Errorf("empty list")
	} else if curNode.Value == value {
		l.header = l.header.Next
		l.size--
	}
	for {
		if curNode.Next.Value == value {
			curNode.Next = curNode.Next.Next
			l.size--
			return nil
		}
		curNode = curNode.Next
	}

	// return fmt.Errorf("not found value")
}

// Traverse 单链表反转
func (l *LinkedList) Traverse() *LinkedList {
	traversedList := &LinkedList{}
	curNode := l.header
	for {
		if curNode == nil {
			break
		}
		traversedList.Insert(curNode.Value, 0)
		curNode = curNode.Next
	}
	return traversedList
}

// String 打印出当前单链表
func (l *LinkedList) String() {
	if l.IsEmpty() {
		fmt.Printf("This is an empty list %d\n", l.Len())
		return
	}

	curNode := l.header
	fmt.Println("-------LinkList Element-------")

	index := 0
	for {
		fmt.Printf("Node %d, values: %v \n", index, curNode.Value)
		index++
		curNode = curNode.Next
		if curNode == nil {
			break
		}

	}

	fmt.Printf("-------LinkList Len : %d------- \n", l.size)
}

// UnmarshalListBySlice  根据数组反序列化链表
func UnmarshalListBySlice(nums []interface{}) *LinkedList {
	linkedList := &LinkedList{}
	for _, v := range nums {
		linkedList.Append(v)
	}

	return linkedList
}

// UnmarshalListByRand 根据随机数反序列化链表
func UnmarshalListByRand(maxNum, len int) *LinkedList {
	linkedList := &LinkedList{}
	for i := 0; i < len; i++ {
		value := rand.Intn(maxNum)
		linkedList.Append(value)
	}

	return linkedList
}
