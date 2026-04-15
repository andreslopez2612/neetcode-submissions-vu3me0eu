func maxProfit(prices []int) int {
	var num int
	for i := 0; i < len(prices); i++{
		for j := i+1; j < len(prices); j++{
			profit := prices[j] - prices[i]
			if profit > 0 && profit > num{
				num = profit
			}
		}
	}

	return num
}
