package main

import (
	"ztm/pkg/display"
	"ztm/pkg/message"
)

func main() {
	message.Hi()
	display.Display("Hello from display")
	message.Exciting("an exciting message")
}
