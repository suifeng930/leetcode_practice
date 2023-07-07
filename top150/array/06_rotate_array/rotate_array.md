
## 轮转数组

>
> 给定一个整数数组 nums ,将数组中的元素向又轮转 k 个位置，其中 k 是非负数
> 
>

>
> **进阶**
>  尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题
>  你可以使用空间复杂度为 O(1) 的原地算法解决这个问题吗？
> 

**示例1:**

```
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
```

**示例2:**

```
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释: 
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]
```


### 解法1

>
> 使用额外的数组来将每个元素放在正确的位置， 用n 表示数组的长度，我们遍历原数组，将原数组下角标为i 的元素放至 新数组下标为 (i+k) mod n 的位置
> 最后将新数组拷贝至原数组即可。
> 
> 

>
> 时间复杂度： O(n) 其中 n 为数组的长度  
> 空间复杂度： O(n)

```go
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
```

### 解法2

>
> 环状替代， 将被替换的元素保存在变量temp中，从而避免额外的数组开销

```go
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

```

### 解法3

```go
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

```