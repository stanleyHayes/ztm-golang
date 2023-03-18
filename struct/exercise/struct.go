package main

import "fmt"

type Rectangle struct {
	Breadth float64
	Length  float64
}

func (r *Rectangle) Area() float64 {
	return r.Breadth * r.Length
}
func (r *Rectangle) Perimeter() float64 {
	return r.Breadth*2 + r.Length*2
}

func main() {
	rect := Rectangle{
		Length:  2,
		Breadth: 4,
	}
	area := rect.Area()
	perimeter := rect.Perimeter()

	fmt.Println(area, perimeter)

	rect.Breadth *= 2
	rect.Length *= 2

	area = rect.Area()
	perimeter = rect.Perimeter()

	fmt.Println(area, perimeter)
}
