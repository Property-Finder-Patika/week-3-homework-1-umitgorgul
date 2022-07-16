package main

import "fmt"

func SieveForPrimeDetect(n int) []int {
	// check if number is prime, if number is prime (true) if number isn't prime (false)
	numbers := make([]bool, n+1) // numbers len is n+1 and each value in this array is starting at 0 and ending in n like [0, 1, 2, 4, 5, ... n]
	for i := 2; i < n+1; i++ {   // we make all numbers true except 0
		numbers[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// start from 2 then make every number false which are multiples of i(2 then 3 then 5 ...)
		if numbers[p] == true { // check is p still true if not then it isn't a prime number after all so try next number
			// make all numbers false which are multiples of p
			for i := p * 2; i <= n; i += p {
				numbers[i] = false
			}
		}
	}

	var primes []int          // create array of prime numbers inside input number
	for p := 2; p <= n; p++ { // start from 2 check if number is true then if it true append it in prime array
		if numbers[p] == true {
			primes = append(primes, p)
		}
		// then return prime array to console
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
