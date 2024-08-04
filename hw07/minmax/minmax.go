package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbersInRange(ch chan int, min, max int) {
	for {
		num := rand.Intn(max-min+1) + min
		ch <- num
		time.Sleep(1 * time.Second)
	}
}

func findMinMax(ch chan int, resultCh chan [2]int) {
	var min, max int
	min, max = <-ch, <-ch

	for num := range ch {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		resultCh <- [2]int{min, max}
	}
}

func main() {
	numCh := make(chan int)
	resultCh := make(chan [2]int)

	go generateNumbersInRange(numCh, 1, 100)
	go findMinMax(numCh, resultCh)

	for {
		result := <-resultCh
		fmt.Printf("Smallest: %d, Largest: %d\n", result[0], result[1])
	}
}
