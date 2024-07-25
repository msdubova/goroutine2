package minmax

import "fmt"

type MinMax struct {
	Min int
	Max int
}

func CalculateMinMax(chRandInt <-chan int, chMinMax chan<- MinMax) {
	fmt.Println("✅ Виконується горутина 2 calculateMinMax")

	var min, max int
	first := true
	for num := range chRandInt {
		if first {
			min, max = num, num
			first = false
		} else {
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}
	}
	chMinMax <- MinMax{Min: min, Max: max}
	close(chMinMax)
	fmt.Println("Min Max  числа передано в канал")
}
