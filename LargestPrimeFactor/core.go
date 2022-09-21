package LargestPrimeFactor

import "sort"

var numbers = map[int]bool{2: true, 3: true}

func IsPrime(n int) bool {
	isPrime, ok := numbers[n]
	if ok && isPrime {
		return true
	}
	if n < 2 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			numbers[n] = false
			return false
		}
	}
	numbers[n] = true
	return true
}

func MinPrimeFactor(n int) int {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return i
		}
	}
	return n
}

func PrimeFactors(n int) []int {
	if IsPrime(n) {
		return []int{n}
	}
	minPrimeFactor := MinPrimeFactor(n)
	return append(PrimeFactors(n/minPrimeFactor), minPrimeFactor)
}

func LargestPrimeFactor(n int) int {
	primeFactors := PrimeFactors(n)
	sort.Slice(primeFactors, func(i int, j int) bool {
		return primeFactors[i] > primeFactors[j]
	})
	return primeFactors[0]
}
