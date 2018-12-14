package main

import (
	"sync"
	"time"
)

func main() {
	var w sync.WaitGroup

    w.Add(1)
    go func() {
        time.Sleep(5 * time.Second)
    }()

    w.Wait()
}
