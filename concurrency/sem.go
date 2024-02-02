package main

import (
	"fmt"
	"sync"
	"time"
)

func semMain() {
	ch := make(chan int, 2)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		ch <- 1
		wg.Add(1)
		go doSome(i, &wg, ch)
	}
	wg.Wait()

}

func doSome(i int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	fmt.Printf("Starting %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Done %d\n", i)
	x := <-ch
	fmt.Println(x)
}
