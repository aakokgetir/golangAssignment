package main

import (
	"fmt"
	"multiplesOfThreeOrFive/EvenFibonacciNumbers"
	"multiplesOfThreeOrFive/LargestPalindromeProduct"
	"multiplesOfThreeOrFive/LargestPrimeFactor"
	"multiplesOfThreeOrFive/LargestProductInASeries"
	"multiplesOfThreeOrFive/SmallestProduct"
	"multiplesOfThreeOrFive/SomePrime"
	"multiplesOfThreeOrFive/SpecialPythagoreanTriplet"
	"multiplesOfThreeOrFive/SumOfMultiplesOfThreeOrFive"
	"multiplesOfThreeOrFive/SumSquareDifference"
	"multiplesOfThreeOrFive/SummationOfPrimes"
)

func main() {
	fmt.Println("1. Find the sum of all the multiples of 3 or 5 below 1000.")
	fmt.Println(SumOfMultiplesOfThreeOrFive.Sum(1000))
	fmt.Println()

	fmt.Println("2. Find the sum of the even-valued terms in fibonacci sequence whose values do not exceed four million.")
	fmt.Println(EvenFibonacciNumbers.SumShort(4000000))
	fmt.Println(EvenFibonacciNumbers.SumLong(4000000))
	fmt.Println()

	fmt.Println("3. What is the largest prime factor of the number 600851475143 ?")
	fmt.Println(LargestPrimeFactor.LargestPrimeFactor(600851475143))
	fmt.Println()

	fmt.Println("4. Find the largest palindrome made from the product of two 3-digit numbers.")
	fmt.Println(LargestPalindromeProduct.LargestPalindromeNumbersForProduct(3))
	fmt.Println()

	fmt.Println("5. What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?")
	fmt.Println(SmallestProduct.SmallestProduct(1, 20))
	fmt.Println()

	fmt.Println("6. Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.")
	fmt.Println(SumSquareDifference.SumSquareDifference(100))
	fmt.Println()

	fmt.Println("7. What is the 10 001st prime number?")
	fmt.Println(SomePrime.NthPrime(10001))
	fmt.Println()

	fmt.Println("8. Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?")
	part, product := LargestProductInASeries.LargestProductInASeries(LargestProductInASeries.Number, 4)
	fmt.Println(part, "=", product)
	part, product = LargestProductInASeries.LargestProductInASeries(LargestProductInASeries.Number, 13)
	fmt.Println(part, "=", product)
	fmt.Println()

	fmt.Println("9. There exists exactly one Pythagorean triplet for which a + b + c = 1000. Find the product abc.")
	a, b, c := SpecialPythagoreanTriplet.TripletFor(1000)
	fmt.Println(a, b, c, "=", a*b*c)
	fmt.Println()

	fmt.Println("10. Find the sum of all the primes below two million.")
	fmt.Println(SummationOfPrimes.SummationOfPrimes(2000000))
	fmt.Println()

}
