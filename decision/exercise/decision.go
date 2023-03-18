package main

import (
	"fmt"
	"log"
	"strings"
)

func accessGranted() string {
	return "Welcome, access granted!"
}

func accessDenied() string {
	return "Sorry, access denied!"
}

func main() {
	var role string
	fmt.Println("Enter your role (admin, manager, member, guest or contractor): ")
	_, err := fmt.Scan(&role)
	if err != nil {
		log.Fatal(err)
	}

	var weekday string
	fmt.Println("Enter your weekday: ")
	_, err = fmt.Scan(&weekday)
	if err != nil {
		log.Fatal(err)
	}

	if strings.ToLower(role) == "admin" || strings.ToLower(role) == "manager" {
		fmt.Println(accessGranted())
	} else if strings.ToLower(weekday) == "saturday" || strings.ToLower(weekday) == "sunday" {
		if strings.ToLower(role) == "contractor" {
			fmt.Println(accessGranted())
		} else {
			fmt.Println(accessDenied())
		}
	} else if strings.ToLower(weekday) == "monday" ||
		strings.ToLower(weekday) == "wednesday" ||
		strings.ToLower(weekday) == "friday" {
		if strings.ToLower(role) == "guest" {
			fmt.Println(accessGranted())
		} else {
			fmt.Println(accessDenied())
		}
	} else if strings.ToLower(weekday) == "monday" ||
		strings.ToLower(weekday) == "tuesday" ||
		strings.ToLower(weekday) == "wednesday" ||
		strings.ToLower(weekday) == "thursday" ||
		strings.ToLower(weekday) == "friday" {
		if strings.ToLower(role) == "member" {
			fmt.Println(accessGranted())
		} else {
			fmt.Println(accessDenied())
		}
	} else {
		fmt.Println(accessDenied())
	}
}
