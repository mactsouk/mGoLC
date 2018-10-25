package main

import (
	"fmt"
)

func main() {
	numbers := make(chan int, 5)
	numbers <- 0
	numbers <- 1
	numbers <- 2
	numbers <- 3
	numbers <- 4

	// This will not work
	// numbers <- 15

	// This will work
	select {
	case numbers <- 15:
	default:
		fmt.Println("Not enough space in the channel!")
		num := <-numbers
		fmt.Println(num)
	}

	select {
	case numbers <- 15:
	default:
		fmt.Println("Not enough space in the channel!")
		num := <-numbers
		fmt.Println(num)
	}

	// You can read a buffered channel as usual
	for i := 0; i < 10; i++ {
		select {
		case num := <-numbers:
			fmt.Println(num)
		default:
			break
		}
	}
}
