## 多数元素

>
> 给定一个大小为n 的数组 nums ,返回其中的多数元素。多数元素是指在数组中出现次数大于 n/2 的元素
> 你可以假设数组是非空的，并且给定的数组总是存在多数元素
>
>

**示例1：**

```
输入：nums = [3,2,3]
输出：3
```

**示例2：**

```
输入：nums = [2,2,1,1,1,2,2]
输出：2
```

## 解法

### 解法1

>
> 排序思路： 既然数组中有出现 大于n/2 的元素，那么排好序的数组中，相同元素总是相邻的，即存在长度 > n/2的长串由相同元素构成的连续子数组。
> 因此 取 nums[n/2]的值一定是这个排好序的子数组元素；
> 
>  

> 
> 时间复杂度： O(n*logn)
> 空间复杂度： O(n*logn)


```go
func majorityElement(nums []int) int {

sort.Ints(nums)
return nums[len(nums)/2]
}
```

### 解法2

>
> 哈希表： key为数组元素，value 为数组元素出现的次数
> 
> 遍历整个数组，对记录每个数值出现的次数（利用map,其中key为数值，value为出现的次数）；接着遍历每个map中的元素，寻找value值大于 nums.length/2的值，返回
> 
> 

>
> 时间复杂度：O(n)
> 空间复杂度：O(n)
> 
```go
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

```


### 解法3

>
> 摩尔投票法
> 
> 候选人 初始化为 nums[0];票数初始化为1.当遇到与cond_num 相同的数，则票数count=count+1;
> 否则票数 count=count-1 .当票数减为0时，更换候选人，并将票数 count 重置为1 遍历完数组后，cond_num即为最终答案。
> 
> 

>
> 投票法时遇到相同的则 **票数+1** 遇到不同的则 **票数-1**；且多数元素的个数大于 [n/2]；其余元素的个数总和一定是小于 [n/2]的
> 因此【多数元素的个数 - 其余元素的个数总和】 的结果肯定是 >=1的。这就相当于每个 多数元素和其他元素两梁互相抵消，抵消到最后肯定还剩余至少一个 多数元素
> 
> 

> 
> 时间复杂度：O(n)
> 空间复杂度：O(1)
> 

```go
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
```