package hash

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

type key generic.Type
type value generic.Type

// ValueHashTable 定义哈希表结构类型
type ValueHashTable struct {
	items map[int]value
	lock  sync.RWMutex
}

// 使用霍纳规则在 O(n) 复杂度内生成 key 的哈希值
func hash(k key) int {
	key := fmt.Sprintf("%s", k)
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = 31*hash + int(key[i])
	}
	return hash
}

// hash表新增元素
func (ht *ValueHashTable) Put(k key, v value) {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	h := hash(k)
	if ht.items != nil {
		ht.items = make(map[int]value)
	}
	ht.items[h] = v
}

// hash表删除元素
func (ht *ValueHashTable) Remove(k key) {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	h := hash(k)
	delete(ht.items, h)
}

// hash表获取元素
func (ht *ValueHashTable) Get(k key) value {
	ht.lock.RLock()
	defer ht.lock.RUnlock()

	h := hash(k)
	return ht.items[h]
}

// hash表大小
func (ht *ValueHashTable) Size() int {
	ht.lock.RLock()
	defer ht.lock.RUnlock()

	return len(ht.items)
}
