package hash

import (
	"fmt"
	"testing"
)

// 构造hashtable
func generateHashTable(len, start int) *ValueHashTable {
	ht := ValueHashTable{}
	for i := start; i < (start + len); i++ {
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

	if "value4" != ht.Get("key3") {
		t.Errorf("hash表元素获取出错,获取到 %s", ht.Get("key3"))
	}

	for i := 1; i <= ht.Size(); i++ {
		fmt.Printf("初始值[key=%d]: hash值=%d: value=%s \n",
			i,
			hash(fmt.Sprintf("key%d", i)),
			ht.items[hash(fmt.Sprintf("key%d", i))],
		)
	}
}
