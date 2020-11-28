package lur

import (
	"fmt"
)

// DLinkNode  双向列表
type DLinkNode struct {
	Key      int
	Value    int
	PreNode  *DLinkNode
	NextNode *DLinkNode
}

// LRUCache LRUCache
type LRUCache struct {
	// Cache容量
	Capacity int
	HashMap  map[int]*DLinkNode
	head     *DLinkNode
	tail     *DLinkNode
}

// NewLRUCache 构造LRUCache
func NewLRUCache(capacity int) LRUCache {
	lurCache := LRUCache{Capacity: capacity}
	lurCache.HashMap = make(map[int]*DLinkNode, capacity)
	return lurCache
}

// Get 获取LURCache中列的值，获取后刷新LRUCache，将值刷新到链表头
func (l *LRUCache) Get(key int) int {
	if v, ok := l.HashMap[key]; ok {
		l.refreshNode(v)
		return v.Value
	}

	return -1
}

// Put 向LURCache的链表头插入新节点，如超过Cache最大容量，淘汰一个表末尾最老的节点
func (l *LRUCache) Put(key, value int) {
	if v, ok := l.HashMap[key]; !ok {
		if len(l.HashMap) >= l.Capacity {
			oldestKey := l.removeNode(l.tail)
			delete(l.HashMap, oldestKey)
		}
		node := DLinkNode{Key: key, Value: value}
		l.addNode(&node)
		l.HashMap[key] = &node
	} else {
		v.Value = value
		l.refreshNode(v)
	}
}

// Size 获取当前Cache列表使用容量
func (l *LRUCache) Size() int {
	return len(l.HashMap)
}

// 根据节点刷新LURCache
func (l *LRUCache) refreshNode(node *DLinkNode) {
	if node == l.head {
		return
	}

	l.removeNode(node)
	l.addNode(node)
}

// 从双向列表中删除节点
func (l *LRUCache) removeNode(node *DLinkNode) int {
	if node == l.tail {
		l.tail = l.tail.PreNode
		l.tail.NextNode = nil
	} else if node == l.head {
		l.head = l.head.NextNode
		l.head.PreNode = nil
	} else {
		node.PreNode.NextNode = node.NextNode
		node.NextNode.PreNode = node.PreNode
	}

	return node.Key
}

// 添加新节点向双向列表头尾添加新值
func (l *LRUCache) addNode(node *DLinkNode) {
	if l.head != nil {
		l.head.PreNode = node
		node.PreNode = nil
		node.NextNode = l.head
	}
	l.head = node
	if l.tail == nil {
		l.tail = node
	}
}

// 获取LURCache当前列表
func (l *LRUCache) printCache() {
	node := l.head

	for node.NextNode != nil {
		fmt.Printf("key %d = [value] %d \n", node.Key, node.Value)
		node = node.NextNode
	}

	fmt.Printf("key %d = [value] %d \n", l.tail.Key, l.tail.Value)

}
