package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type Item struct {
	Name  string
	State bool
}

func activate(item *Item) {
	if !item.State {
		item.State = Active
	}
}

func deactivate(item *Item) {
	if item.State {
		item.State = Inactive
	}
}

func checkout(items *[]Item) {
	for _, item := range *items {
		deactivate(&item)
	}
}

func main() {
	items := []Item{
		{"JBL Speaker", Active},
		{"Alienware Laptop", Active},
		{"Alienware Monitor", Active},
		{"Power Bank", Active},
	}
	fmt.Println("--- Initial States ---")
	fmt.Println(items)
	fmt.Println("--- Updating state of first item ---")
	deactivate(&items[0])
	fmt.Println(items)
	fmt.Println("--- Deactivating all items using checkout ---")
	checkout(&items)
	fmt.Println(items)
}
