package matematica

import (
	"fmt"
	"strconv"
)

// WeightedAverage get value
func WeightedAverage(numbers ...float64) float64 {
	total := 0.0

	for _, value := range numbers {
		total += value
	}
	average := total / float64(len(numbers))
	averageRound, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", average), 64)
	return averageRound
}
