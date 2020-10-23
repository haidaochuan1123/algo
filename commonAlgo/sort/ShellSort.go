package sort

// ShellSort  希尔排序算法，基于插入排序
func ShellSort(sortList []int) []int {
	// 控制分步步长
	sortListLen := len(sortList)
	for step := sortListLen / 2; step > 0; step /= 2 {
		// 针对step内的元素进行插入排序
		for i := step; i < sortListLen; i += step {
			for j := i - step; j >= 0 && sortList[j] > sortList[j+step]; j -= step {
				sortList[j], sortList[j+step] = sortList[j+step], sortList[j]
			}
		}
	}

	return sortList
}
