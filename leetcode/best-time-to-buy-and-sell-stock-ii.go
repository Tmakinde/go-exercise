// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/
func maxProfit(prices []int) int {
 // buying price is first price
 // if buying price is greather than current market price, then buying price is current market price
 // if buying price is less than current market price, then :
    // -calculate profit, profit += current price - buying price
    // buying price is current market price

    profit := 0
    buyingPrice := prices[0]

    if (len(prices) == 0) {
        return profit
    }

    for i := 1; i < len(prices); i++ {
        currentPrice := prices[i]
        if (buyingPrice >= currentPrice) {
            buyingPrice = currentPrice
        } else {
            profit += currentPrice - buyingPrice
            buyingPrice = currentPrice
        }
    }
    return profit
}