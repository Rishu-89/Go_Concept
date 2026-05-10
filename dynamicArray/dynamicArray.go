package dynamicArray

import "fmt"

func DynamicArray() {
	prices := []float64{10, 20, 30}
	fmt.Println(prices)
	prices=append(prices, 40)
	fmt.Println(prices)
	prices=prices[1:] // remove first element from list
	fmt.Println(prices)
	prices=prices[:2] // remove last element from list
	fmt.Println(prices)
}