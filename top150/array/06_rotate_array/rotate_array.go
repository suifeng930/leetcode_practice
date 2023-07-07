package main

func main() {

}

func rotate(nums []int, k int) {

	newNums := make([]int, len(nums))
	for i, val := range nums {
		newNums[(i+k)%len(nums)] = val
	}
	copy(nums, newNums)

}

// [1,2,3,4,5,6,7]   k=3
//  0 1 2 3 4 5 6
//  3 4 5 6 7 8 9
//  3%7 =3
//  4%7 =4
//  5%7 =5
//  6%7 =6
//  7%7 =0
//  8%7 =1
//  9%7 =2

func rotate02(nums []int, k int) {

	n := len(nums)

	k = k % n
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func rotate03(nums []int, k int) {
	k = k % len(nums) //       [1,2,3,4,5,6,7]   k=3
	// 翻转数组所有元素           [7,6,5,4,3,2,1]   k=3
	reverse(nums, 0, len(nums)-1)
	// 翻转 [0, k mod (n-1)]     [5,6,7,4,3,2,1]
	reverse(nums, 0, k-1)
	// 翻转 [ k mod (n), n-1]   [5,6,7,1,2,3,4]
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}

}
