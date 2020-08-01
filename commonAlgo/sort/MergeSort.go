package sort

import "fmt"

// MergeSort 归并排序算法
func MergeSort(sortList []int) []int {
	lenSortList := len(sortList)

	// 如果待排序列表为1，结束递归退出
	if lenSortList <= 1 {
		return sortList
	}

	middle := lenSortList / 2
	leftList := MergeSort(sortList[:middle])
	rightList := MergeSort(sortList[middle:])
	sortList = merge(leftList, rightList)
	fmt.Printf("排序结果为:%v \n", sortList)
	return sortList
}

func merge(leftList, rightList []int) []int {

	sortedList := make([]int, 0, len(leftList)+len(rightList))

	for len(leftList) > 0 || len(rightList) > 0 {
		// 分割后左右列表个数可能不一致
		if len(leftList) == 0 {
			return append(sortedList, rightList...)
		}
		if len(rightList) == 0 {
			return append(sortedList, leftList...)
		}

		// 取出左右两个列表第一个元素对比，按顺序插入新列表
		// 并将插入新列表的元素从原列表删除
		if leftList[0] < rightList[0] {
			sortedList = append(sortedList, leftList[0])

			leftList = leftList[1:]
		} else {
			sortedList = append(sortedList, rightList[0])

			rightList = rightList[1:]
		}
	}

	return sortedList
}
