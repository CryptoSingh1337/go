package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	x := 5

	fmt.Println(x)

	test := make(chan int)
	test <- 1
	test <- 2

	for i := range test {
		fmt.Println(i)
	}
}
