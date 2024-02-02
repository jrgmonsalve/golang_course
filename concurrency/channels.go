package main

import "fmt"

func channels() {
	ch := make(chan int, 3)
	ch <- 1
	fmt.Println(<-ch)
}
