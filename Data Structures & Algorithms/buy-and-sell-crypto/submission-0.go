func maxProfit(prices []int) int {
    max_profit := 0
    running_min_price := 101 // this is the maximum possible stock price + 1

    for _, price := range prices {
        if running_min_price > price {
            running_min_price = price
        }

        trade_value := price - running_min_price
        if trade_value > max_profit {
            max_profit = trade_value
        }
    }

    return max_profit
}
