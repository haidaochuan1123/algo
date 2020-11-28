package sort

// InsertionSort 插入排序算法
func InsertionSort(sortList []int) []int {

	for i := range sortList {
		if i < 1 {
			continue
		}

		key := sortList[i]
		// 待排序元素key，如果小于之前1到i-1的有序列表，进行排序
		if sortList[i] < sortList[i-1] {
			for j := i; j > 0; j-- {
				if key < sortList[j-1] {
					sortList[j], sortList[j-1] = sortList[j-1], sortList[j]
				} else {
					break
				}
			}
		}

		// fmt.Printf("第%d次排序结果为:%v \n", i, sortList)
	}

	return sortList
}
