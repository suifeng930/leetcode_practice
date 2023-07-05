package main

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}

	println(removeDuplicates(nums))
	for _, value := range nums {
		print(value)
	}

}

func removeDuplicates(nums []int) int {

	if len(nums) <= 2 {
		return len(nums)
	}
	slow, fast := 2, 2
	for fast < len(nums) {

		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
