package main

import "fmt"

func main() {
	prices := [4]float64{10.23, 21.7}
	prices[3]=19
	prices[2]=12.4
	fmt.Println(prices)
	featuredPrices:=prices[1:3]			// including one excluding last
	fmt.Println(featuredPrices)
	featuredPrices=prices[1:]
	fmt.Println(featuredPrices)
	featuredPrices=prices[:3]
	fmt.Println(featuredPrices)
	featuredPrices[1]=9777
	fmt.Println(prices)					// [10.23 9777 12.4 19] Important
	
}