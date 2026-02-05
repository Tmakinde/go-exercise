// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
    if len(prices) == 0{
        return 0
    }

    profit := 0
    todayPrice := 0

    buyingPrice := prices[0]

    for x := 1; x <= len(prices)-1; x++ {
        todayPrice = prices[x]

        if todayPrice > buyingPrice {
            if  (todayPrice - buyingPrice) > profit {
                // best time to sell
                profit = todayPrice - buyingPrice
            }
        } else {
            // best time to buy
            buyingPrice = todayPrice
        }
    }

    return profit
}