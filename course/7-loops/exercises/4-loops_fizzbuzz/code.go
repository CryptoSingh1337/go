package main

import "fmt"

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		str := ""
		if i%3 == 0 {
			str += "fizz"
			if i%5 == 0 {
				str += "buzz"
			}
		} else if i%5 == 0 {
			str = "buzz"
		}
		if str == "" {
			fmt.Println(i)
		} else {
			fmt.Println(str)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
