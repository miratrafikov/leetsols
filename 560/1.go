package main

import "fmt"

func main() {
	arr := []int{1,2,3}
	fmt.Println(subarraySum(arr, 3))
}

func subarraySum(nums []int, k int) int {
	positionToSumSoFar := make([]int, len(nums))
	positionToSumSoFar[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		positionToSumSoFar[i] = positionToSumSoFar[i - 1] + nums[i]
	}
	numberOfSubarraysEqualToK := 0
	if nums[0] == k {
		numberOfSubarraysEqualToK++
	}
	for i := 1; i < len(nums); i++ {
		if positionToSumSoFar[i] == k {
			numberOfSubarraysEqualToK++
		}
		for j := 0; j < i; j++ {
			if positionToSumSoFar[i] - positionToSumSoFar[j] == k {
				numberOfSubarraysEqualToK++
			}
		}
	}
	return numberOfSubarraysEqualToK
}
