package main

import (
	"fmt"
)

func wpMain() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	nWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))
	for w := 1; w <= nWorkers; w++ {
		go Worker(w, jobs, results)
	}

	for _, task := range tasks {
		jobs <- task
	}
	close(jobs)

	for a := 1; a <= len(tasks); a++ {
		fmt.Println(<-results)
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "started job", job)
		fib := Fibonacci(job)
		fmt.Println("worker", id, "finished job", job, "with result", fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
