package main

import "fmt"

func discountedPrice(product string, price float64) float64 {
	var discount float64
	switch(product) {
	case "apples":
	fallthrough
	case "apple":
		discount = .10
	case "bananas":
	// case "banana":
		discount = .20
	}
	discountedprice := price - price * discount
	return discountedprice
}

func main() {
	fmt.Println(discountedPrice("apples", 100))
	fmt.Println(discountedPrice("apple", 100))
	fmt.Println(discountedPrice("orange", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("bananas", 100))
	fmt.Println(discountedPrice("oranges", 100))
}

