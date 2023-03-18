package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func main() {
	products := [4]Product{
		{Name: "JBL", Price: 700},
		{Name: "iPhone 14", Price: 6000},
		{Name: "Samsung Galaxy s23 Ultra", Price: 17000},
	}

	fmt.Println(products[3])
	fmt.Println("There are", len(products), "products available")

	total := 0.0
	for _, product := range products {
		total += product.Price
	}

	fmt.Println("The total cost of products is", total)
	product := Product{Name: "Alienware 17", Price: 25000}
	products[3] = product

	total = 0.0
	for _, product := range products {
		total += product.Price
	}
	fmt.Println("The total cost of products is", total)
}
