package arrays

import (
	"fmt"
)

func SearchInsertHelper() {
	f := fmt.Println
	array := []int{1, 3, 5, 6}
	fmt.Printf("PROGRAM 4 : Where will 2 be inserted in the array %v", array)
	f()
	f(searchInsert(array, 2))
	f()
	f()

}

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
		if i+1 < len(nums) && nums[i] < target && nums[i+1] > target {
			return i + 1
		}
		if i+1 == len(nums) && nums[len(nums)-1] < target {
			return len(nums)
		}
	}
	return 0
}
