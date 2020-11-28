package sort

// HeapSort 堆排序算法
func HeapSort(sortList []int) []int {
	for i := 0; i < len(sortList); i++ {
		minHeap(sortList[i:])
	}

	return sortList
}

func minHeap(input []int) {
	inputLen := len(input)
	if inputLen <= 1 {
		return
	}

	for i := inputLen/2 - 1; i >= 0; i-- {
		if 2*i+1 <= inputLen-1 && input[2*i+1] < input[i] {
			input[2*i+1], input[i] = input[i], input[2*i+1]
		}
		if 2*i+2 <= inputLen-1 && input[2*i+2] < input[i] {
			input[2*i+2], input[i] = input[i], input[2*i+2]
		}
	}
}
