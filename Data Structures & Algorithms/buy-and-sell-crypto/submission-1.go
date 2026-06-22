func maxProfit(prices []int) int {
	var minPrice = prices[0]
	maxProfit := 0

	for p1 := 0; p1 < len(prices); p1++ {
		if prices[p1] < minPrice{
			minPrice = prices[p1]
		} else if prices[p1]-minPrice > maxProfit {
			 maxProfit = prices[p1] - minPrice
		}
	}
	return maxProfit
}
