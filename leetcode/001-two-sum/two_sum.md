https://leetcode-cn.com/problems/add-two-numbers/

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。并且数组中同一个元素不能使用两遍。

 

示例:
```
给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
```

解析：
正常暴力解法需要遍历两遍数组，时间复杂度为O(n^2)；
由于限定了输出结果只有一个答案且数组元素唯一，因此可以借助golang map。
target－num[i]作为map的key，遍历数组，如果key存在，将其返回即可。
时间复杂度为O(n)，空间复杂度O(n)


结果
```
go test  -v
=== RUN   TestTwoSum
nums = [3 4 6] target = 6 result = []
nums = [3 4 6] target = 7 result = [0 1]
nums = [2 7 11 15] target = 9 result = [0 1]
nums = [1 2 3 4 5] target = 5 result = [1 2]
nums = [3] target = 3 result = []
--- PASS: TestTwoSum (0.00s)
PASS
ok      github.com/svoid/algo/leetcode/001-two-sum      0.006s
```