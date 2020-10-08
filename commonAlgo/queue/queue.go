package queue

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type ItemQueue struct {
	items []Item
	lock  sync.RWMutex
}

// 创建一个空队列
func NewItemQueue() *ItemQueue {
	i := []Item{}
	return &ItemQueue{items: i}
}

// 入队列，添加到队列尾
func (q *ItemQueue) EnQuene(i Item) {
	q.lock.Lock()
	q.items = append(q.items, i)
	q.lock.Unlock()
}

// 出队列，队列为先入先出
func (q *ItemQueue) DeQuene() *Item {
	q.lock.Lock()
	item := q.items[0]
	q.items = q.items[1:]
	q.lock.Unlock()
	return &item
}

// 判断队列是否为空
func (q *ItemQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// 判断队列长度
func (q *ItemQueue) Size() int {
	return len(q.items)
}
