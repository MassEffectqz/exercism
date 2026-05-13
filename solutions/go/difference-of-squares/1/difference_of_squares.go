package differenceofsquares
import (
	"math"
)

func SquareOfSum(n int) int {
	var sum int  
	for i := 1; i <= n; i++ {
		sum += i
	}
	pw := math.Pow(float64(sum), 2)
	return int(pw)
}

func SumOfSquares(n int) int {
	var pw float64
	for i := 1; i <= n; i++ {
		pw += math.Pow(float64(i), 2)
	}
	return int(pw) 
}

func Difference(n int) int {
	n = SquareOfSum(n) - SumOfSquares(n)
	return n 
}