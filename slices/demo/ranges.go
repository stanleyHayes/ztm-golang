package main

import "fmt"

func main() {
	slice := []string{"Hello", "world", "!"}

	for i, element := range slice {
		fmt.Println(i, element)

		for _, letter := range element {
			fmt.Printf("\t%q\n", letter)
		}
	}
}
