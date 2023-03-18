package main

import (
	"fmt"
	"log"
)

func main() {
	var age int
	fmt.Println("Enter age: ")
	_, err := fmt.Scan(&age)
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case age == 0:
		fmt.Println("You are a newborn")
	case age > 0 && age < 4:
		fmt.Println("You are a toddler")
	case age > 3 && age < 13:
		fmt.Println("You're a child")
	case age > 12 && age < 18:
		fmt.Println("You're a teenager")
	case age > 17:
		fmt.Println("You're an adult")

	}
}
