package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	var diff float64
	for i := 0; ; i++ {
		diff = z
		z = z - (z*z-x)/(2*z)
		if math.Abs(diff-z) < 1e-6 {
			break
		}
	}
	return z
}

func main() {
	estValue := Sqrt(2)
	realValue := math.Sqrt(2)
	diff := math.Abs(estValue - realValue)
	fmt.Println(estValue, realValue, diff)
}
