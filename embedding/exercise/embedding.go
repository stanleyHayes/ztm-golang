package main

import "fmt"

type Bytes int
type Celsius float32

type BandwidthUsage struct {
	Amount []Bytes
}

type CpuTemperature struct {
	Temperatures []Celsius
}

type MemoryUsage struct {
	Amount []Bytes
}

func (b *BandwidthUsage) AverageBandwidth() int {
	sum := 0
	for i := 0; i < len(b.Amount); i++ {
		sum += int(b.Amount[i])
	}
	return sum / len(b.Amount)
}

func (c *CpuTemperature) AverageCPUTemperature() int {
	sum := 0
	for i := 0; i < len(c.Temperatures); i++ {
		sum += int(c.Temperatures[i])
	}
	return sum / len(c.Temperatures)
}

func (m *MemoryUsage) AverageMemoryUsage() int {
	sum := 0
	for i := 0; i < len(m.Amount); i++ {
		sum += int(m.Amount[i])
	}
	return sum / len(m.Amount)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemperature
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{
		Amount: []Bytes{50000, 100000, 130000, 80000, 90000},
	}
	temperatures := CpuTemperature{
		Temperatures: []Celsius{50, 51, 53, 551, 52},
	}

	memory := MemoryUsage{Amount: []Bytes{
		800000, 000000, 810000, 820000, 800000},
	}

	dashboard := Dashboard{BandwidthUsage: bandwidth, CpuTemperature: temperatures, MemoryUsage: memory}

	fmt.Printf("Average bandwidth usage: %v\n", dashboard.AverageBandwidth())
	fmt.Printf("Average bandwidth usage: %v\n", dashboard.AverageCPUTemperature())
	fmt.Printf("Average bandwidth usage: %v\n", dashboard.AverageMemoryUsage())

}
