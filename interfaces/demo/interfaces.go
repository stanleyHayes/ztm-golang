package main

import "fmt"

type Preparer interface {
	Prepare()
}

type Chicken string
type Salad string

func (c Chicken) Prepare() {
	fmt.Println("Cook chicken")
}

func (c Salad) Prepare() {
	fmt.Println("Chop Salad")
	fmt.Println("Add dressing")
}

func PrepareDishes(dishes []Preparer) {
	fmt.Println("Preparing dishes")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--Dish: %v--\n", dish)
		dish.Prepare()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{Chicken("Chicken wings"), Salad("Iceberg Salad")}
	PrepareDishes(dishes)
}
