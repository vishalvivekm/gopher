package main

import "fmt"

type Item struct {
	Name  string
	Price float64
}

func getItemsInPriceRange(items []Item, minPrice, maxPrice float64) []Item {

	var ansSlice []Item
	for _, ele := range items {
        // if( items[index].Price >= minPrice && items[index].Price <= maxPrice ) {
		// 	ansSlice = append(ansSlice, items[index])
		// }
		if( ele.Price >= minPrice && ele.Price <= maxPrice ) {
			ansSlice = append(ansSlice, ele)
		}
	}
      /*for index, _ := range items {
        if( items[index].Price >= minPrice && items[index].Price <= maxPrice ) {
			ansSlice = append(ansSlice, items[index])
		}
}*/
	return ansSlice
}

func main() {
	items := []Item{
		{Name: "Apple", Price: 0.5},
		{Name: "Banana", Price: 0.25},
		{Name: "Orange", Price: 0.75},
		{Name: "Pineapple", Price: 1.5},
	}

	fmt.Println(getItemsInPriceRange(items, 0.0, 1.0))
	fmt.Println(getItemsInPriceRange(items, 0.5, 1.0))
	fmt.Println(getItemsInPriceRange(items, 0.75, 1.5))
}
