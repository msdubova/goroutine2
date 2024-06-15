package minmax

import "fmt"

type MinMax struct {
	Min int
	Max int
}

func CalculateMinMax(numbers []int, ch chan<- MinMax) {
	fmt.Println("✅ Виконується горутина 2 calculateMinMax")

	var min int = numbers[0]
	var max int = numbers[0]

	for _, item := range numbers {
		if item < min {
			min = item
		}
		if item > max {
			max = item
		}
	}

	fmt.Println("Min число ", min)
	fmt.Println("Max число ", max)
	result := MinMax{Min: min, Max: max}
	ch <- result

	fmt.Println("Min Max  числа передано в канал", result)
}
