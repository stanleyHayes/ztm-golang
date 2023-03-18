package main

import "fmt"

func main() {
	var color = "black"
	birthYear, age := 1993, 29

	var (
		first = "S"
		last  = "H"
	)

	var ageInDays int
	ageInDays = age * 365

	fmt.Println(color, birthYear, first, last, ageInDays)

	a, b := 5, 8
	sum := a + b
	difference := a - b
	product := a * b
	quotient := a / b

	fmt.Println(sum, difference, product, quotient)
}
