package diffsquares

import "math"

func SquareOfSums(n int) int {
	var accumulator int

	for i := 1; i <= n; i++ {
		accumulator += i
	}

	return int(math.Pow(float64(accumulator), 2))
}

func SumOfSquares(n int) int {
	var accumulator float64

	for i := 1; i <= n; i++ {
		accumulator += math.Pow(float64(i), 2)
	}

	return int(accumulator)
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
