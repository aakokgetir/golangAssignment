package SmallestProduct

import (
	"multiplesOfThreeOrFive/LargestPrimeFactor"
)

func Max[K float64 | float32 | int | int64 | int32 | int8 | int16 | uint8 | uint | uint16](v1 K, v2 K) K {
	if v1 < v2 {
		return v2
	}
	return v1
}

func SmallestProduct(start int, end int) int {
	primes := make(map[int]int)
	if start <= 1 {
		start = 2
	}

	if start >= end {
		return -1
	}

	for i := start; i <= end; i++ {
		primeCounts := make(map[int]int)
		for _, factor := range LargestPrimeFactor.PrimeFactors(i) {
			primeCounts[factor] += 1
		}
		for prime, count := range primeCounts {
			primes[prime] = Max(primes[prime], count)
		}
	}
	product := 1
	for prime, count := range primes {
		for i := 0; i < count; i++ {
			product *= prime
		}
	}
	return product
}
