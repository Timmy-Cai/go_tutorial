package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // 1

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 2
		go func(i int) {
			defer wg.Done() // 3
			fmt.Printf("do some thing, index:[%d] \n", i)
			time.Sleep(3 * time.Second)
		}(i) // *
	}
	wg.Wait() // 4
	fmt.Println("main finished")
}
