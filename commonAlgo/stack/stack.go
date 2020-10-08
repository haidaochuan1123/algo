package stack

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type ItemStack struct {
	items []Item
	lock  sync.RWMutex
}

// 创建一个空栈
func NewItemStack() *ItemStack {
	i := []Item{}
	return &ItemStack{items: i}
}

// 入栈，添加到栈尾
func (s *ItemStack) Push(i Item) {
	s.lock.Lock()
	s.items = append(s.items, i)
	s.lock.Unlock()
}

// 出栈，后入先出
func (s *ItemStack) Pop() *Item {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.lock.Unlock()
	return &item
}

// 判断栈是否为空
func (s *ItemStack) IsEmpty() bool {
	return len(s.items) == 0
}

// 判断栈长度
func (s *ItemStack) Size() int {
	return len(s.items)
}
