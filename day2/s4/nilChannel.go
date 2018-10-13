package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)

	for {
		select {
		case input := <-c:
			sum = sum + input
		case <-t.C:
			c = nil
			fmt.Println(sum)
			fmt.Println("add() ending...")
			return
		}
	}
}

func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
	fmt.Println("send() ending...")
}

func main() {
	c := make(chan int)
	go add(c)
	go send(c)

	time.Sleep(3 * time.Second)
}
