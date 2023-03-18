package main

import "fmt"

func Sum(vars ...int) int {
	total := 0
	for _, number := range vars {
		total = number
	}
	return total
}

func Max(numbers ...int) int {
	maximum := numbers[0]
	for _, number := range numbers {
		if number > maximum {
			maximum = number
		}
	}
	return maximum
}
func Min(numbers ...int) int {
	minimum := numbers[0]
	for _, number := range numbers {
		if number < minimum {
			minimum = number
		}
	}
	return minimum
}

func main() {
	numbers := []int{1, 2, 6, -4, 0}
	fmt.Println("Numbers:", numbers)
	fmt.Println("Sum:", Sum(numbers...))
	fmt.Println("Max:", Max(numbers...))
	fmt.Println("Min:", Min(numbers...))
}
