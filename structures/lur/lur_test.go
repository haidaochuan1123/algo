package lur

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	lur := NewLRUCache(5)
	for i := 1; i <= 20; i++ {
		fmt.Printf("Cache 当前大小为%d,总容量为%d \n", lur.Size(), lur.Capacity)
		lur.Put(i, i+100)
	}
	lur.printCache()
	fmt.Println(lur.Get(5))
	fmt.Println(lur.Get(19))
	lur.Put(1, 101)
	lur.Put(17, 17)
	lur.printCache()
}
