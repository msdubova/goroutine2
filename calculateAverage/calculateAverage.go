package calculateAverage

import "fmt"

func CalculateAverage(numbers []int, ch chan<- int) {
	fmt.Println("✅ Виконується горутина 2 calculateAverage")
	var total int
	for _, item := range numbers {
		total += item
	}

	var average = total / (len(numbers))
	fmt.Println("Середнє число ", average)
	ch <- average
	fmt.Println("Середнє число передано в канал")
}
