package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

var counter int

func main() {

	for range 10000 {
		wg.Go(increment)
	}
	wg.Wait()

	fmt.Println("Counter value is", counter)

}

func increment() {
	mu.Lock()
	defer mu.Unlock()
	counter = counter + 1
}
