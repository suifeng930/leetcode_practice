package main

import "fmt"

func main() {

	nums1 := []int{0}
	m := 0
	nums2 := []int{2}
	n := 1
	value := merge_tst(nums1, m, nums2, n)
	fmt.Println(value)
}

func merge_tst(nums1 []int, m int, nums2 []int, n int) []int {

	length := len(nums1)
	for n > 0 {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[length-1] = nums1[m-1]
			m--
			length--
		} else {
			nums1[length-1] = nums2[n-1]
			length--
			n--
		}
	}
	return nums1
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	length := len(nums1)
	for n > 0 {

		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[length-1] = nums1[m-1]
			m--
		} else {
			nums1[length-1] = nums2[n-1]
			n--
		}
		length--
	}

}
