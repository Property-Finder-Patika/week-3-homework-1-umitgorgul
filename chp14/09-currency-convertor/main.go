package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//   1. Declare a few constants with iota. They're going to be the keys of the array.
	const (
		EUR = iota // 0
		GBP        // 1
		JPY        // 2
	)

	currency := [...]float64{
		EUR: 0.88,   // currency[0]
		GBP: 0.78,   // currency[1] // gbp higher than eur irl tho
		JPY: 113.02, // currency[2]
	}

	args := os.Args[1:] // getting input
	if len(args) != 1 { // check error
		fmt.Println("Please provide the amount to be converted.")
		return
	}

	input, err := strconv.ParseFloat(args[0], 64) //str to float so we can convert in currency
	if err != nil {
		fmt.Println("Invalid amount. It should be a number.")
		return
	}

	fmt.Printf("%.2f USD is %.2f EUR\n", input, currency[EUR]*input) // input * 0.88
	fmt.Printf("%.2f USD is %.2f GBP\n", input, currency[GBP]*input) // input * 0.78
	fmt.Printf("%.2f USD is %.2f JPY\n", input, currency[JPY]*input) // input * 113.02
}
