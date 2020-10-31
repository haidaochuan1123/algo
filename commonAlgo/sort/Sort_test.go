package sort

import (
	"fmt"
	"testing"
)

var sortList = []int{111, 11, 2, 555, 555, 555, 1, 2313, 44, 5435, 436, 54, 6, 32}

// BubbleSort 冒泡排序算法
func TestBubbleSort(t *testing.T) {

	out := BubbleSort(sortList)
	fmt.Printf("sortedList := %v \n", out)
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

// InsertionSort 插入排序算法
func TestInsertionSort(t *testing.T) {
	out := InsertionSort(sortList)
	fmt.Printf("sortedList := %v \n", out)
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

// MergeSort 插入排序算法
func TestMergeSort(t *testing.T) {
	out := MergeSort(sortList)
	fmt.Printf("sortedList := %v \n", out)
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

func TestQuickSort(t *testing.T) {

	out := QuickSort(sortList, 0, len(sortList)-1)
	fmt.Printf("sortedList := %v \n", out)
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

func TestSelectionSort(t *testing.T) {

	out := SelectionSort(sortList)
	fmt.Printf("sortedList := %v \n", out)
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

func TestShellSort(t *testing.T) {

	out := ShellSort(sortList)
	fmt.Printf("ShellSort := %v \n", out)
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

func TestHeapSort(t *testing.T) {

	out := HeapSort(sortList)
	fmt.Printf("HeapSort := %v \n", out)
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

func TestCountingSort(t *testing.T) {

	out := CountingSort(sortList)
	fmt.Printf("CountingSort := %v \n", out)
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

func TestBucketSort(t *testing.T) {

	out := BucketSort(sortList, 10)
	fmt.Printf("BucketSort := %v \n", out)
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

func TestRadixSort(t *testing.T) {

	out := RadixSort(sortList)
	fmt.Printf("RadixSort := %v \n", out)
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

func TestGenerateRandomNumberRange(t *testing.T) {
	arr1 := generateRandomNumberRange(10, 100, 80)
	fmt.Println(arr1)

	arr2 := generateRandomNumber(100, 100)
	fmt.Println(arr2)
}
