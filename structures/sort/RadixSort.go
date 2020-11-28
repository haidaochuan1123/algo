package sort

// RadixSort 基数排序算法
func RadixSort(sortList []int) []int {
	maxDigit := sortList[0]

	for i := range sortList {
		if sortList[i] >= maxDigit {
			maxDigit = sortList[i]
		}
	}

	num := 1
	for ; maxDigit > 0; maxDigit = maxDigit / 10 {
		bucket := make([][]int, 10)
		for i := range sortList {
			n := sortList[i] / num % 10
			bucket[n] = append(bucket[n], sortList[i])
		}

		sortedIndex := 0
		for i := range bucket {
			for j := range bucket[i] {
				sortList[sortedIndex] = bucket[i][j]
				sortedIndex++
			}
		}

		num *= 10
	}

	return sortList
}
