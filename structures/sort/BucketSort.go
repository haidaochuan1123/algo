package sort

// BucketSort 桶排序算法
// bucketSize 定义桶大小
func BucketSort(sortList []int, bucketSize int) []int {
	var (
		minValue    = sortList[0]
		maxValue    = sortList[0]
		sortedIndex = 0
	)

	for i := range sortList {
		if minValue > sortList[i] {
			minValue = sortList[i]
		}
		if maxValue < sortList[i] {
			maxValue = sortList[i]
		}
	}

	// 建立新的桶排序数组
	bucket := make([][]int, (maxValue-minValue)/bucketSize+1)

	// 将数据分配到各个桶中
	for i := range sortList {
		bucketTarget := (sortList[i] - minValue) / bucketSize
		bucket[bucketTarget] = append(bucket[bucketTarget], sortList[i])
	}

	// 对桶内数据进行插入排序
	for i := range bucket {
		if len(bucket[i]) <= 0 {
			continue
		}

		bucket[i] = InsertionSort(bucket[i])

		for j := range bucket[i] {
			sortList[sortedIndex] = bucket[i][j]
			sortedIndex++
		}
	}

	return sortList
}
