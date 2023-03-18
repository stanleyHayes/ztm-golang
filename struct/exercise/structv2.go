package main

import "fmt"

type Coordinate struct {
	x, y int
}

type Rect struct {
	a Coordinate
	b Coordinate
}

func width(rect Rect) int {
	return rect.b.x - rect.a.x
}

func length(rect Rect) int {
	return rect.a.y - rect.b.y
}

func area(rect Rect) int {
	return length(rect) * width(rect)
}

func perimeter(rect Rect) int {
	return (width(rect) * 2) + (length(rect) * 2)
}

func printInfo(rect Rect) {
	fmt.Println("Area is", area(rect))
	fmt.Println("Perimeter is", perimeter(rect))
}

func main() {
	rect := Rect{a: Coordinate{x: 2, y: 4}, b: Coordinate{x: 4, y: 8}}
	printInfo(rect)
}
