package sort

// SelectionSort  选择排序算法
func SelectionSort(sortList []int) []int {
	listLength := len(sortList)
	var min int

	// 从左往右遍历数组，找出右边最小的数，移到遍历开始位置
	for i := 0; i < listLength-1; i++ {
		min = i
		for j := listLength - 1; j > i; j-- {
			if sortList[j] < sortList[min] {
				min = j
			}

		}

		sortList[i], sortList[min] = sortList[min], sortList[i]
	}

	return sortList
}
