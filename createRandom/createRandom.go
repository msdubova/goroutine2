package random

import (
	"fmt"
	"math/rand"
)

func CreateRandom(n int, ch chan<- []int) {
	fmt.Println("✅ Виконується горутина 1 createRandom, починаю створювати числа")
	numbers := make([]int, n)

	for i := range numbers {
		numbers[i] = rand.Intn(100)
	}

	fmt.Println("Створено рандомні числа", numbers)
	ch <- numbers
	fmt.Println("Рандомні числа передано в канал")

}
