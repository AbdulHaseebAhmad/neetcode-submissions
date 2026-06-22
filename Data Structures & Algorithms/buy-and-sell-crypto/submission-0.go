func maxProfit(prices []int) int {
	var total = 0 
	for p1 := 0; p1 < len(prices); p1++ {
		 for p2 := len(prices) - 1 ; p2 > p1 ; p2-- {
			if prices[p2] - prices[p1] > total {
				total = prices[p2] - prices[p1]
		} 
	}
	}
	return total
}
