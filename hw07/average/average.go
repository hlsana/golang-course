package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers(ch chan int) {
	for {
		num := rand.Intn(100)
		ch <- num
		time.Sleep(1 * time.Second)
	}
}

func averageNumbers(ch chan int, avgCh chan float64) {
	var sum, count int
	for num := range ch {
		sum += num
		count++
		avg := float64(sum) / float64(count)
		avgCh <- avg
	}
}

func displayAverage(avgCh chan float64) {
	for avg := range avgCh {
		fmt.Printf("Average: %.2f\n", avg)
	}
}

func main() {
	numCh := make(chan int)
	avgCh := make(chan float64)

	go generateNumbers(numCh)
	go averageNumbers(numCh, avgCh)
	go displayAverage(avgCh)

	select {}
}
