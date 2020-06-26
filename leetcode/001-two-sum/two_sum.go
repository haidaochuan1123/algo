package leetcode

func twoSum(nums []int, target int) []int {

	numMap := make(map[int]int)
	for idx, num := range nums {
		need := target - num
		if _, ok := numMap[need]; ok {
			return []int{numMap[need], idx}
		}

		numMap[num] = idx
	}

	return nil
}
