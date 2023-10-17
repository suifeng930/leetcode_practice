package main

func main() {

	nums := []int{2, 3, 1, 1, 4}
	value := canJump(nums)
	v2 := canJumpV2(nums)
	println(value)
	println(v2)

}

func canJumpV2(nums []int) bool {
	arrLen := len(nums)
	if arrLen == 1 {
		return true
	}
	// 默认当前下标为最大可跳跃的长度
	maxJump := nums[0]
	// 遍历数组，当前数组下标小于maxJump; 且当前下标小于数组长度
	for i := 0; i < maxJump && i < arrLen; i++ {
		// 更新当前下标可跳跃的值
		nowJump := nums[i] + i + 1
		// 如果当前下标可跳跃的值大于最大下标，更新最大下标
		if nowJump > maxJump {
			maxJump = nowJump
		}
	}
	if maxJump >= arrLen {
		return true
	}
	return false

}

func canJump(nums []int) bool {

	flag := 0
	arrLen := len(nums)
	for i, val := range nums {
		if flag >= arrLen-1 {
			return true
		}
		if i <= flag {
			flag = compareToMax(flag, i+val)
		}
	}
	return false
}

func compareToMax(a, b int) int {
	if a > b {
		return a
	}
	return b

}
