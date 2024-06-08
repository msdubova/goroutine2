package random

import (
	"Gourotines2/calculateMinMax"
	"fmt"
	"math/rand"
)

func CreateRandom(n int, ch chan<- []int, minMaxNumbers *calculateMinMax.MinMax) {
	fmt.Println("✅ Виконується горутина 1 createRandom, починаю створювати числа")
	numbers := make([]int, n)

	for i := range numbers {
		numbers[i] = rand.Intn(100)
	}

	fmt.Println("Створено рандомні числа", numbers)
	ch <- numbers
	fmt.Println("Рандомні числа передано в канал")

	fmt.Println("Перша горутина друкує результат другої  - Min Max", *minMaxNumbers)
}
