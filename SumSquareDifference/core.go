package SumSquareDifference

func SquaresSum(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i * i
	}
	return result
}

func SumSquare(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func SumSquareDifference(n int) int {
	return SumSquare(n) - SquaresSum(n)
}
