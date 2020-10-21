package sort

// QuickSort 快速排序算法
func QuickSort(sortList []int, start, last int) []int {
	if start < last {
		i, j := start, last
		key := sortList[(start+last)/2]
		for i <= j {
			for sortList[i] < key {
				i++
			}
			for sortList[j] > key {
				j--
			}

			if i <= j {
				sortList[i], sortList[j] = sortList[j], sortList[i]
				i++
				j--
			}
		}

		if start < j {
			QuickSort(sortList, start, j)
		}

		if last > i {
			QuickSort(sortList, i, last)
		}

	}

	return sortList
}
