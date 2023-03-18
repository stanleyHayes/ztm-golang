package main

import "fmt"

type Item struct {
	name  string
	price int
}

func printStats(products [4]Item) {
	var cost, totalItems int

	for i := 0; i < len(products); i++ {
		item := products[i]
		cost += item.price

		if item.name != "" {
			totalItems += 1
		}
	}

	fmt.Println("Last item on the list:", products[totalItems-1])
	fmt.Println("Total items", totalItems)
	fmt.Println("Total cost", cost)
}

func main() {
	products := [4]Item{
		{name: "JBL", price: 700},
		{name: "iPhone 14", price: 6000},
		{name: "Samsung Galaxy s23 Ultra", price: 17000},
	}
	printStats(products)
	product := Item{name: "Bread", price: 50}
	products[3] = product
	printStats(products)
}
