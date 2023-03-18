package main

import "fmt"

func main() {
	var name = "Vladislaus Draguila"
	fmt.Println("My name is", name)

	var myName string = "Kathy"
	fmt.Println("Name =", myName)

	username := "usermaatre"
	fmt.Println("username =", username)

	var sum int
	fmt.Println("The sum is", sum)

	email, password := "hayfordstanley@gmail.com", "Typingmyname123#"
	fmt.Println("Password is", password, "and email is", email)

	email, password = "dev.stanley.hayford@gmail.com", "IloveAdolfHitler1889@Braunau"
	fmt.Println("Password is", password, "and email is", email)

	price1, price2 := 1, 2

	sum = price1 + price2

	price1, price2 = 3, 4
	sum += price1 + price2
	fmt.Println("The sum of", price1, "and", price2, "is", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)

	fmt.Println("Lesson name =", lessonName)
	fmt.Println("Lesson type =", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println("Word1 =", word1, "Word2=", word2)
}
