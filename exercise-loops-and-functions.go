package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 10; i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func Cqrt(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 10; i++ {
		fmt.Println(z)
		z -= (z*z*z - x) / (3 * z * z)
	}
	return z
}

func Nqrt(x float64, n float64) float64 {
	var z float64 = 1
	for i := 0; i < 100; i++ {
		fmt.Println(z)
		z -= (math.Pow(z, n) - x) / ((n) * math.Pow(z, n-1))
	}
	return z
}

func main() {
	fmt.Println(Nqrt(math.Pow(10, 5), 5))
}
