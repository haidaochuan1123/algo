package sort

import (
	"fmt"
	"testing"
)

// InsertionAlgo 插入排序算法
func TestInsertionAlgo(t *testing.T) {
	var sortList = []int{111, 11, 2, 555, 555, 555, 1, 2313, 44, 5435, 436, 54, 6, 32}

	out := InsertionAlgo(sortList)
	fmt.Printf("sortedList := %v", out)
	sortListLen := len(sortList)
	outLen := len(out)
	if sortListLen != outLen {
		t.Errorf("排序错误，返回值与初始值长度不同")
	}
	for i := range out {
		if i == outLen-1 {
			break
		}

		if out[i] > out[i+1] {
			t.Errorf("返回列表排序不符合规则")
		}
	}
}
