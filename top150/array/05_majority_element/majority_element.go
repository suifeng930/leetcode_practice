package main

import "sort"

func main() {

}

func majorityElement(nums []int) int {

	sort.Ints(nums)
	return nums[len(nums)/2]
}

func majorityElement02(nums []int) int {
	val := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == val {
			count++
		} else {
			count--
			if count == 0 {
				val = nums[i]
				count = 1
			}
		}
	}
	return val
}

func majorityElement03(nums []int) int {

	tempMap := make(map[int]int, len(nums))
	for _, num := range nums {
		tempMap[num] += 1
		if tempMap[num] > len(nums)>>1 {
			return num

		}
	}
	return 0
}
