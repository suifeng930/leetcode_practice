package main

func main() {

	nums := []int{7, 1, 5, 3, 6, 4}
	println(maxProfit(nums))
}

func maxProfit(prices []int) int {

	minPrice := prices[0]
	maxPrice := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if (prices[i] - minPrice) > maxPrice {
			maxPrice = prices[i] - minPrice
		}
	}
	return maxPrice
}

func maxProfit01(prices []int) int {

	minPrice := -1
	maxPrice := 0
	for _, price := range prices {
		if minPrice == -1 {
			minPrice = price
			continue
		}
		maxPrice = max(maxPrice, price-minPrice)
		minPrice = min(minPrice, price)
	}
	return maxPrice
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
