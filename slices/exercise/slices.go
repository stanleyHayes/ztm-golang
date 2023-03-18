package main

import "fmt"

type Part string

func UnloadAssemblyLine(parts []Part) {
	for i := 0; i < len(parts); i++ {
		part := parts[i]
		fmt.Println(part)
	}
}

func main() {
	parts := []Part{"Pipe", "Bolt", "Sheet"}
	UnloadAssemblyLine(parts)
	parts = append(parts, "Washer", "Wheel")
	UnloadAssemblyLine(parts)
	parts = parts[3:]
	UnloadAssemblyLine(parts)
}
