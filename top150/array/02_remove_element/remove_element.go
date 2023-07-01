package main

func main() {

	nums := []int{3, 2, 2, 3}
	//nums02 := []int{3, 2, 2, 3}
	val := 3
	println("1:", removeElement(nums, val))

	println("2:", removeElement02(nums, val))

}

func removeElement(nums []int, val int) int {

	var ans = 0
	// 数组指针下标
	for i := 0; i < len(nums); i++ {
		// 遍历数组，做等值比较，如果val不相等，则将当前nums[i]移动到数组下标nums[ans],且ans增加
		if nums[i] != val {
			nums[ans] = nums[i]
			ans++
		}
	}
	return ans
}

func removeElement02(nums []int, val int) int {

	println("sdasd", nums)
	println(len(nums))

	// 数组指针下标
	left, right := 0, len(nums)
	for left < right {
		// 如果nums[left] ==val 将 nums[right-1]覆盖nums[left], right--
		// 注意这个 nums[right-1] 并没有做比较排序，但会再下一次将再次比较 nums[left]==val
		// 相当于两次for loop 分别比较 nums[left]和 nums[right-1]的值
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
