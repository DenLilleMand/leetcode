package main

//Say you have an array for which the ith element is the price of a given stock on day i.
//
//Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).
//
//Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

// Looks a lot like a dynamic programming problem, can likely be solved with
// a greedy algorithm. If done by brute force, we have to calculate all of
// the different variations and compare min/max
// What is a combination ? e.g. buy 7 sell day 1, buy day 2, sell day 3

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit := 0
	for i := 1; i < len(prices); i++ {
		prev_p := prices[i-1]
		p := prices[i]
	}
	return profit
}

func main() {
	input := []int{7, 1, 5, 3, 6, 4}

}
