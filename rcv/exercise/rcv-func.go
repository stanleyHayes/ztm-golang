package main

import "fmt"

const (
	MaxEnergy uint = 100
	MaxHealth uint = 100
	MinEnergy uint = 0
	MinHealth uint = 0
)

type Player struct {
	Health uint
	Energy uint
	Name   string
}

func (player *Player) AddHealth(health uint) {
	player.Health += health
	if player.Health > MaxHealth {
		player.Health = MaxHealth
	}
	fmt.Println(player.Name, "Add", health, "health ->", player.Health)
}

func (player *Player) AddEnergy(energy uint) {
	player.Energy += energy
	if player.Energy > MaxEnergy {
		player.Energy = MaxEnergy
	}
	fmt.Println(player.Name, "Add", energy, "energy ->", player.Energy)
}

func (player *Player) ReduceHealth(health uint) {
	player.Health -= health
	if player.Health < MinHealth {
		player.Health = MinHealth
	}
}

func (player *Player) ReduceEnergy(energy uint) {
	player.Energy -= energy
	if player.Energy < MinEnergy {
		player.Energy = MinEnergy
	}
}

func main() {
	player := Player{Health: 100, Name: "Kunglao", Energy: 100}
	fmt.Println(player)
	player.ReduceEnergy(10)
	player.ReduceEnergy(7)
	fmt.Println(player)
	player.ReduceHealth(120)
	fmt.Println(player)
}
