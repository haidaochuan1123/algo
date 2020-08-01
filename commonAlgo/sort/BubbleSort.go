package sort

import "fmt"

// BubbleSort 冒泡排序算法
func BubbleSort(sortList []int) []int {
	for i := range sortList {

		//对于列表中相邻的元素进行比较，每次循环结束，大的元素沉底
		for j := 0; j < len(sortList)-i-1; j++ {
			if sortList[j] > sortList[j+1] {
				sortList[j], sortList[j+1] = sortList[j+1], sortList[j]
			}
		}
		fmt.Printf("第%d次排序结果为:%v \n", i, sortList)
	}
	return sortList
}
