package random

import (
	"Gourotines2/minmax"
	"fmt"
	"math/rand"
	"time"
)

func CreateRandom(n int, chRandInt chan<- int, chMinMax <-chan minmax.MinMax) {
	fmt.Println("✅ Виконується горутина 1 createRandom, починаю створювати числа")
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, n)

	for i := 0; i < n; i++ {
		num := rand.Intn(100)
		numbers[i] = num
		chRandInt <- num
	}

	fmt.Println("Створено рандомні числа", numbers)
	fmt.Println("Рандомні числа передано в канал")
	close(chRandInt)

	minMaxNumbers := <-chMinMax

	fmt.Println("Перша горутина друкує результат другої  - Min Max", minMaxNumbers.Min, minMaxNumbers.Max)

}
