package sort

import (
	"math/rand"
	"time"
)

// generateRandomNumber 生成指定大小的数组
func generateRandomNumber(maxNum, arrayLen int) []int {
	var arrs []int
	for i := 0; i < arrayLen; i++ {
		value := rand.Intn(maxNum)
		arrs = append(arrs, value)
	}

	return arrs
}

// generateRandomNumberRange 不重复数组
func generateRandomNumberRange(start, end, arrayLen int) []int {
	if end < start {
		return nil
	}

	arrs := make([]int, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(arrs) < arrayLen {
		value := r.Intn(end - start)
		exist := false
		for _, v := range arrs {
			if v == value || value < start {
				exist = true
				break
			}
		}

		if !exist {
			arrs = append(arrs, value)
		}
	}

	return arrs
}
