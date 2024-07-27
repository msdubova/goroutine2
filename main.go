package main

import (
	"Gourotines2/minmax"
	"Gourotines2/random"
	"fmt"
	"time"
)

func main() {
	chRandInt := make(chan int)
	chMinMax := make(chan minmax.MinMax)

	go random.CreateRandom(10, chRandInt, chMinMax)

	fmt.Println("Отримали рандомні числа з каналу")

	go minmax.CalculateMinMax(chRandInt, chMinMax)
	// minMaxNumbers := <-chMinMax
	// fmt.Println("Отримали Min Max числа з рандомних чисел з каналу", minMaxNumbers)

	time.Sleep(2 * time.Second)
}
