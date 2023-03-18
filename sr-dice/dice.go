package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var dice int
	var sides int
	var rolls int

	fmt.Println("How many dice will you use")
	_, err := fmt.Scan(&dice);
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("How many sides is each dice")
	_, err = fmt.Scan(&sides);
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("How many times are you rolling the dice")
	_, err = fmt.Scan(&rolls);
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	sum := 0
	for count < rolls {
		for i := 0; i < dice; i++ {
			roll := rand.Intn(sides) + 1
			sum += roll
		}
		if sum%2 == 0 {
			fmt.Println("Even")
		} else if sum%2 == 1 {
			fmt.Println("Odd")
		}

		if sum == 7 {
			fmt.Println("Lucky 7")
		}
		if rolls == 2 && sum == 2 {
			fmt.Println("Snake eyes")
		}
		count++
	}

}
