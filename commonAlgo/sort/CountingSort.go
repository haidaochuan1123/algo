package sort

// CountingSort 计数排序算法
func CountingSort(sortList []int) []int {
	var (
		maxValue    = sortList[0]
		sortedIndex = 0
	)

	for i := range sortList {
		if maxValue < sortList[i] {
			maxValue = sortList[i]
		}
	}
	// 分配长度为最大数值长度的数组
	bucket := make([]int, maxValue+1)

	for i := range sortList {
		bucket[sortList[i]]++
	}

	// 遍历新数组，其下标为排序数据的值
	for j := range bucket {
		for bucket[j] > 0 {
			sortList[sortedIndex] = j
			sortedIndex++
			bucket[j]--
		}
	}

	return sortList
}
