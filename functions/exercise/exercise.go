package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello", name)
}

func fGreet() string {
	return "Heil Fuhher"
}

func add(a, b, c int) int {
	return a + b + c
}

func doubleLuck(a, b int) (int, int) {
	return a * 2, b * 2
}

func main() {
	a, b, c := 1, 2, 3
	greet("Adolf Hitler")
	message := fGreet()
	fmt.Println(message)
	fmt.Println(add(a, b, c))
	fmt.Println(doubleLuck(add(a, b, c), add(a*2, b*2, c*2)))
}
