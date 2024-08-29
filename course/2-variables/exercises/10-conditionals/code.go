package main

import "fmt"

func main() {
	messageLen := 20
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	// don't touch above this line

	if messageLen <= maxMessageLen && messageLen > 0 {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}
