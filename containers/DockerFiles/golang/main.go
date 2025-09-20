package main

import (
	"fmt"
	"sort"
)

func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, x := range nums {
		i := sort.SearchInts(dp, x)
		if i == len(dp) {
			dp = append(dp, x)
		} else {
			dp[i] = x
		}
	}
	return len(dp)
}

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Println(lengthOfLIS(nums))
}
