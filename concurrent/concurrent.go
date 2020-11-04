package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // 1

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1) // 2位置一定要放在这里
		go func(i int) {
			defer wg.Done() // 3位置一定要放在这里
			fmt.Printf("do some thing, index:[%d] \n", i)
			//time.Sleep(3 * time.Second)
		}(i) // *
	}
	wg.Wait() // 4位置一定要放在这里
	fmt.Println("main finished")
}
