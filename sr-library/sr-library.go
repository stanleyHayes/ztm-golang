package main

import "time"

type Title string
type Name string

type Book struct {
	Name
}

type Member struct {
	Name
	books map[Title]LendAudit
}

type LendAudit struct {
	Checkout time.Time
	Checkin  time.Time
}

type Library struct {
	Members []Member
	Books   []Book
}

func main() {

}
