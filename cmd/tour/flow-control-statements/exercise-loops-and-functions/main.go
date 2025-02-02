package main

import (
	"fmt"
	"math"
)

const closeMargin = 1e-9

func Sqrt(x float64) float64 {
	value, difference := x/2, 0.

	for i := 1; i <= 10; i++ {
		value -= (value*value - x) / (2 * value)
		fmt.Println(i, "回目の計算値: ", value)
		if math.Abs(value-difference) < closeMargin {
			break
		}
		difference = value
	}
	return value
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(5) == math.Sqrt(5))
}
