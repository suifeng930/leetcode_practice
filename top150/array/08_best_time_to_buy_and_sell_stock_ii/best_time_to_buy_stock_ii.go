package main

func main() {

	prices := []int{7, 1, 5, 3, 6, 4}
	count := maxProfit(prices)
	println(count)

}

func maxProfit(prices []int) int {

	count := 0
	if len(prices) < 2 {
		return count
	}
	for i := 1; i < len(prices); i++ {
		count += max(0, prices[i]-prices[i-1])
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}
