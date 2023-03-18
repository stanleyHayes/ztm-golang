package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	Add = iota
	Modulo
	Subtract
	Multiply
	Divide
	Exponent
)

type Operation int

func (op Operation) Calculate(lhs, rhs int) (int, error) {
	switch op {
	case Add:
		return lhs + rhs, nil
	case Modulo:
		return lhs % rhs, nil
	case Subtract:
		return lhs - rhs, nil
	case Multiply:
		return lhs * rhs, nil
	case Divide:
		if rhs == 0 {
			return 0, errors.New("zero division error: cannot divide by zero")
		}
		return lhs / rhs, nil
	case Exponent:
		return int(math.Pow(float64(lhs), float64(rhs))), nil
	default:
		panic("Unhandled operation")
	}
}

func main() {
	add := Operation(Add)
	lhs, rhs := 6, 9
	result, err := add.Calculate(lhs, rhs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Sum of", lhs, "and", rhs, "is", result)
}
