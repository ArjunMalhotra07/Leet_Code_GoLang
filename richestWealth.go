//https://leetcode.com/problems/richest-customer-wealth/

package main

import "fmt"

func maximumWealthHelper() {
	f := fmt.Println
	f("PROGRAM 4 : Richest Wealth LeetCode Problem")

	matrix := [][]int{{1, 2, 3}, {3, 2, 1}}
	f(maximumWealth(matrix))
	f()
	f()
}
func maximumWealth(accounts [][]int) int {

	ans := 0
	for i := 0; i < len(accounts); i++ {
		sum := 0
		for j := 0; j < len(accounts[i]); j++ {
			sum = sum + accounts[i][j]
		}
		if sum > ans {
			ans = sum
		}

	}
	return ans

}
