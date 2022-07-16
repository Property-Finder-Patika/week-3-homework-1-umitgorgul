package main

import "fmt"

func SieveForPrimeDetect(n int) []int {
	// check if number is prime, if number is prime (true) if number isn't prime (false)
	numbers := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		numbers[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// If numbers[p] is not changed, then it is a prime
		if numbers[p] == true {
			// Update all multiples of p
			for i := p * 2; i <= n; i += p {
				numbers[i] = false
			}
		}
	}

	// return all prime numbers <= n
	var primes []int
	for p := 2; p <= n; p++ {
		if numbers[p] == true {
			primes = append(primes, p)
		}
	}

	return primes
}

func main() {
	var num int // input number by user
	fmt.Print("enter the number: ")
	_, err := fmt.Scan(&num) // get input from user
	if err != nil {          // looks for a error
		fmt.Println("Invalid inputs try again: err:", err)
	}
	fmt.Println(SieveForPrimeDetect(num))
}
