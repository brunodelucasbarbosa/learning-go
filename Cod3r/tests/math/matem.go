package matem

import (
	"fmt"
	"strconv"
)

func Media(nums ...float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	valor, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(sum / float64(len(nums)))), 64)
	return valor;
}