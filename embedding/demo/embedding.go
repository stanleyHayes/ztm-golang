package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type BeltSize int
type Shipping int

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

type Conveyor interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type WarehouseAutomator interface {
	Conveyor
	Shipper
}

type SpamMail struct {
	Amount int
}

func (s *SpamMail) String() string {
	return "Spam mail"
}

func (s *SpamMail) Ship() Shipping {
	return Air
}

func (s *SpamMail) Convey() BeltSize {
	return Small
}

func Automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v conveyor\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

type ToxicWaste struct {
	Weight int
}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

func (t *ToxicWaste) Convey() BeltSize {
	return Large
}

func main() {
	mail := SpamMail{Amount: 40_000}
	Automate(&mail)

	waste := ToxicWaste{Weight: 300}
	Automate(&waste)
}
