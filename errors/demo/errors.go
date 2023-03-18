package main

import (
	"errors"
	"fmt"
)

type Stuff struct {
	Values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.Values) {
		return 0, errors.New(fmt.Sprintf("no element at %v", index))
	} else {
		return s.Values[index], nil
	}
}

func main() {
	stuff := Stuff{}
	value, err := stuff.Get(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
