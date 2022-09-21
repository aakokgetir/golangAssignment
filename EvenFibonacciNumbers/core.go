package EvenFibonacciNumbers

import (
	"errors"
)

var memo = make(map[int]int)

func fib(n int) (int, error) {
	val, ok := memo[n]
	if ok {
		return val, nil
	}
	if n < 1 {
		return 0, errors.New("n can not be smaller then 1")
	} else if n == 1 {
		return 1, nil
	} else if n == 2 {
		return 2, nil
	} else {
		n1, err := fib(n - 1)
		if err != nil {
			return 0, errors.New("reached below 1 in recursion for n - 1")
		}
		n2, err := fib(n - 2)
		if err != nil {
			return 0, errors.New("reached below 1 in recursion for n - 2")
		}
		val = n1 + n2
		memo[n] = val
		return val, nil
	}
}

func Fib(n int) int {
	val, err := fib(n)
	if err != nil {
		panic(err)
	}
	return val
}

func SumShort(max int) int {
	var sum = 0
	for idx, nextFib := 1, Fib(1); nextFib < max; idx, nextFib = idx+1, Fib(idx+1) {
		if nextFib%2 == 0 {
			sum += nextFib
		}
	}
	return sum
}

func SumLong(max int) int {
	var sum = 0
	idx := 1
	nextFib, _ := fib(idx)
	for nextFib < max {
		if nextFib%2 == 0 {
			sum += nextFib
		}
		idx += 1
		nextFib, _ = fib(idx)
	}
	return sum
}
