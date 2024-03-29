// https://leetcode.com/problems/first-missing-positive/
package Cyclic_Sort

import "fmt"

func FirstMissingPositiveHelper() {
	f := fmt.Println
	f("PROGRAM 6 : To find first missing positive number in given array")

	arr := []int{1, 2, 0}
	f("test array ", arr)
	ans := firstMissingPositive(arr)
	f("Number missing", ans)
	f()
	f()
}
func firstMissingPositive(nums []int) int {
	i := 0

	for i < len(nums) {
		correct := nums[i] - 1
		if nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[correct] {
			swap(nums, i, correct)
		} else {
			i++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func swap(nums []int, i int, correct int) {
	temp := nums[i]
	nums[i] = nums[correct]
	nums[correct] = temp
}
