package main

// import "fmt"
import "github.com/Rishu-89/Go_Concept/dynamicArray"

func main() {
	prices := [6]float64{10.23, 21.7,5,6}
	prices[4]=19
	prices[5]=12.4
	// fmt.Println(prices)
	// featuredPrices:=prices[1:3]			// including one excluding last
	// fmt.Println(featuredPrices)
	// featuredPrices=prices[1:]
	// fmt.Println(featuredPrices)
	// featuredPrices=prices[:3]
	// fmt.Println(featuredPrices)
	// featuredPrices[1]=9777
	// fmt.Println(prices)					// [10.23 9777 12.4 19] Important
	
	// featuredPrice:=prices[2:5]
	// fmt.Println(len(featuredPrice),cap(featuredPrice))  // output 3 4
	// 													// slice always remembar the last not first

	dynamicArray.DynamicArray();

}