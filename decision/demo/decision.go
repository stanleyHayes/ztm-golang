package main

import (
	"fmt"
	"log"
)

func average(a, b, c int) float64 {
	return float64(a+b+c) / float64(3)
}

func main() {
	var quiz1, quiz2, quiz3 int

	fmt.Println("Enter what you obtained in quiz 1")
	_, err := fmt.Scan(&quiz1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter what you obtained in quiz 2")
	_, err = fmt.Scan(&quiz2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter what you obtained in quiz 3")
	_, err = fmt.Scan(&quiz3)
	if err != nil {
		log.Fatal(err)
	}

	if quiz1 > quiz2 {
		fmt.Println("Quiz1 scored higher than quiz 2")
	} else if quiz1 < quiz2 {
		fmt.Println("Quiz2 scored higher than quiz 1")
	} else {
		fmt.Println("Quiz 1 & Quiz 2 have the same score")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("Acceptable grades")
	} else {
		fmt.Println("Grades not acceptable")
	}
}
