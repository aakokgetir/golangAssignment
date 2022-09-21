package SumOfMultiplesOfThreeOrFive

func Sum(max int) int {
	var sum int = 0
	for i := 1; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
