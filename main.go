package main

import "fmt"

func main() {
	var word = "Hello, World!\n"
	for _, letter := range word {
		fmt.Println(string(letter))
		fmt.Println(letter)
	}
}
