package main

import "fmt"

const (
	Small = iota
	Standard
	Large
)

type Lift int

type Picker interface {
	Pick() Lift
}

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) String() string {
	return fmt.Sprintf("Motocycle: %v", string(m))
}
func (c Car) String() string {
	return fmt.Sprintf("Car: %v", string(c))
}
func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

func (m Motorcycle) Pick() Lift {
	return Small
}

func (c Car) Pick() Lift {
	return Standard
}

func (t Truck) Pick() Lift {
	return Large
}

func SendToLift(p Picker) {
	switch p.Pick() {
	case Small:
		fmt.Printf("Send %v to small lift\n", p)

	case Standard:
		fmt.Printf("Send %v to standard lift\n", p)

	case Large:
		fmt.Printf("Send %v to large lift\n", p)
	}
}

func main() {
	car := Car("Sporty")
	truck := Truck("MountainCrusher")
	motorcycle := Motorcycle("Croozer")
	
	SendToLift(car)
	SendToLift(truck)
	SendToLift(motorcycle)
}
