// package main

// import "fmt"
// // import "github.com/Rishu-89/Go_Concept/dynamicArray"
// // import "github.com/Rishu-89/Go_Concept/structTypeArray"

// func main() {
// 	prices := []float64{10.23, 21.7,5,6}
// 	// prices[4]=19
// 	// prices[5]=12.4
// 	// fmt.Println(prices)
// 	// featuredPrices:=prices[1:3]			// including one excluding last
// 	// fmt.Println(featuredPrices)
// 	// featuredPrices=prices[1:]
// 	// fmt.Println(featuredPrices)
// 	// featuredPrices=prices[:3]
// 	// fmt.Println(featuredPrices)
// 	// featuredPrices[1]=9777
// 	// fmt.Println(prices)					// [10.23 9777 12.4 19] Important

// 	// featuredPrice:=prices[2:5]
// 	// fmt.Println(len(featuredPrice),cap(featuredPrice))  // output 3 4
// 	 													// slice always remembar the last not first
// 	// dynamicArray.DynamicArray();
// 	// structTypeArray.StructArray();

// 	discount:=[]float64{10,20,30,40}
// 	prices=append(prices,discount...)
// 	fmt.Println(prices)
// }

package main

import "fmt"

func main(){
	website:=map[string]string{
		"Google":"https://googe.com",
		"Amazon web services":"https://aws.com",
	}
	fmt.Println(website)
	fmt.Println(website["Amazon web services"])
	website["linkedin"]="https://linked.com"
	fmt.Println(website)
	delete(website,"Google")
	fmt.Println(website)
}