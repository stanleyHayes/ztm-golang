package main

import "fmt"

type Space struct {
	Occupied bool
}

type ParkingLot struct {
	Spaces []Space
}

func (lot *ParkingLot) OccupySpace(space int) {
	lot.Spaces[space-1].Occupied = true
}

func (lot *ParkingLot) VacateSpace(space int) {
	lot.Spaces[space-1].Occupied = false
}

func main() {
	lot := ParkingLot{
		Spaces: make([]Space, 10),
	}

	fmt.Println(lot)
	lot.OccupySpace(1)
	lot.OccupySpace(2)
	fmt.Println("After occupied:", lot)
	lot.VacateSpace(2)
	fmt.Println("After Vacate:", lot)
}
