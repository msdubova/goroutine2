package main

import (
	"Gourotines2/minmax"
	"Gourotines2/random"
	"fmt"
)

func main() {
	chRandInt := make(chan []int)
	chMinMax := make(chan minmax.MinMax)
	minMax := &minmax.MinMax{
		Min: 0,
		Max: 0,
	}

	go random.CreateRandom(10, chRandInt, minMax)
	numbers := <-chRandInt
	fmt.Println("Отримали рандомні числа з каналу:", numbers)

	go minmax.CalculateMinMax(numbers, chMinMax)
	minMaxNumbers := <-chMinMax
	fmt.Println("Отримали Min Max числа з рандомних чисел з каналу", minMaxNumbers)
}
