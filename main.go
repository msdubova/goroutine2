package main

import (
	"Gourotines2/calculateMinMax"
	random "Gourotines2/createRandom"
	"fmt"
)

func main() {
	chRandInt := make(chan []int)
	chMinMax := make(chan calculateMinMax.MinMax)
	minMax := &calculateMinMax.MinMax{}

	go random.CreateRandom(10, chRandInt, minMax)
	numbers := <-chRandInt
	fmt.Println("Отримали рандомні числа з каналу:", numbers)

	go calculateMinMax.CalculateMinMax(numbers, chMinMax)
	minMaxNumbers := <-chMinMax
	fmt.Println("Отримали Min Max числа з рандомних чисел з каналу", minMaxNumbers)
}
