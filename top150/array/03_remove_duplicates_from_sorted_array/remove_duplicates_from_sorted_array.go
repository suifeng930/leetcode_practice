package main

func main() {

	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			nums[slow+1] = nums[fast]
			slow++
		}
		fast++

	}
	return slow + 1

}
