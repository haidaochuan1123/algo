package list

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
)

// Node 链表结构体
type Node struct {
	Value interface{}
	Next  *Node
}

// LinkedList 链表
type LinkedList struct {
	header *Node
	size   int
	lock   sync.RWMutex
}

// NewNode 返回链表结构体
func NewNode(value interface{}) *Node {
	return &Node{
		Value: value,
	}
}

// NewNodeWithNext 返回链表结构体
func NewNodeWithNext(value interface{}, next *Node) *Node {
	return &Node{
		Value: value,
		Next:  next,
	}
}

// IsEmpty 判断单链表是否为空
func (l *LinkedList) IsEmpty() bool {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.header == nil
}

// Len 返回单链表长度
func (l *LinkedList) Len() int {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.size
}

// 添加元素修改单链表长度
func (l *LinkedList) sizeInc() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.size++
}

// 删除元素修改单链表长度
func (l *LinkedList) sizeDec() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.size--
}

// AddValue 向链表添加元素
// func (l *LinkedList) AddValue(value interface{}) {
// 	if l.header == nil {
// 		node := Node{value: value}
// 		l.header = &node
// 		return
// 	}

// 	item := l.header
// 	for ; item.next != nil; item = item.next {
// 		if item.value == value {

// 		}
// 	}
// }

// Append 向链表最后添加元素
func (l *LinkedList) Append(node *Node) {
	l.lock.Lock()
	defer l.lock.Unlock()

	current := l.header

	if current == nil {
		l.header = node
		l.sizeInc()
		return
	}

	for {
		if current.Next == nil {
			break
		}

		current = current.Next
	}

	current.Next = node
	l.sizeInc()
}

// Prepend 向链表头添加元素
func (l *LinkedList) Prepend(node *Node) {
	l.lock.Lock()
	defer l.lock.Unlock()
	current := l.header
	node.Next = current
	l.header = node
	l.sizeInc()
}

// Find 从链表中查找指定的元素
func (l *LinkedList) Find(value interface{}) bool {
	if l.IsEmpty() {
		fmt.Println("This is an empty list")
		return false
	}

	current := l.header
	for current.Next != nil {
		if current.Value == value {
			return true
		}

		current = current.Next
	}

	if current.Value == value {
		return true
	}

	return false
}

// RemoveNode 删除单链表节点
func (l *LinkedList) RemoveNode(value interface{}) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.IsEmpty() {
		return errors.New("This is an empty list")
	}

	current := l.header

	if current.Value == value {
		l.header = current.Next
		l.sizeDec()
		return nil
	}

	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			l.sizeDec()
			return nil
		}

		current = current.Next

	}

	if current.Value == value {
		current = nil
		l.sizeDec()
		return nil
	}

	return errors.New("not found value")
}

// Traverse 单链表反转
func Traverse(l *LinkedList) {
}

// PrintList 打印出当前单链表
func (l LinkedList) PrintList() {
	if l.IsEmpty() {
		fmt.Println("This is an empty list")
		return
	}

	current := l.header
	fmt.Println("-------LinkList Element-------")

	i := 1
	for ; ; i++ {
		if current.Next == nil {
			break
		}
		fmt.Printf("Node %d, values: %v \n", i, current.Value)
		current = current.Next
	}

	fmt.Printf("Node %d, values: %v \n", i, current.Value)
	fmt.Printf("-------LinkList Len : %d------- \n", l.size)
}

// UnmarshalListBySlice  根据数组反序列化链表
func UnmarshalListBySlice(nums []interface{}) *LinkedList {
	linkedList := &LinkedList{}
	for _, v := range nums {
		linkedList.Append(NewNode(v))
	}

	return linkedList
}

// UnmarshalListByRand 根据随机数反序列化链表
func UnmarshalListByRand(maxNum, len int) *LinkedList {
	linkedList := &LinkedList{}
	for i := 0; i < len; i++ {
		value := rand.Intn(maxNum)
		linkedList.Append(NewNode(value))
	}

	return linkedList
}
