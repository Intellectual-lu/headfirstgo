package avg

import "fmt"

func average() {
	numbers := [3]float64{56.1, 22.5, 36.2}
	var total float64 = 0
	for _, number := range numbers {
		total += number
	}
	fmt.Printf("total: %.2f", total)
	fmt.Println()
	fmt.Printf("average: %.2f", total/float64(len(numbers)))
}
