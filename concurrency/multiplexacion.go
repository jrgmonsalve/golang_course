package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	d1 := 5 * time.Second
	d2 := 2 * time.Second
	go DoSomething(d1, c1, 1)
	go DoSomething(d2, c2, 2)
	for i := 0; i < 2; i++ {
		fmt.Println("index i:", i)
		select {
		case chMsg1 := <-c1:
			fmt.Println(chMsg1)
		case chMsg2 := <-c2:
			fmt.Println(chMsg2)
		}
		fmt.Println("end:", i)
	}

}

func DoSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
