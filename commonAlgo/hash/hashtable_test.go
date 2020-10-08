package hash

import (
	"fmt"
	"testing"
)

// 构造hashtable
func generateHashTable(len, start int) *ValueHashTable {
	ht := ValueHashTable{}
	for i := start; i <= len; i++ {
		ht.Put(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}

	return &ht
}

func TestHashTable(t *testing.T) {
	ht := generateHashTable(20, 1)
	size := ht.Size()
	if size != 20 {
		t.Errorf("hash表个数不对,获取到 %d", size)
	}

	ht.Put("key3", "value4")

	ht.Remove("key2")

	if "valsue3" != ht.Get("key3") {
		t.Errorf("hash表元素获取出错,获取到 %s", ht.Get("key3"))
	}

	for i := 0; i < ht.Size(); i++ {
		fmt.Printf("数组[%d]: key=%d: value=%s \n", i, hash(i), ht.items[hash(i)])
	}
}
