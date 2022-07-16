package main

import "fmt"

func main() {
	var num int // input number by user
	fmt.Print("enter the number: ")
	_, err := fmt.Scan(&num) // get input from user
	if err != nil {          // looks for a error
		fmt.Println("Invalid inputs try again: err:", err)
	}

	i := 2
	// it's simply start from 2 and try divide the number with i until it can't
	// then its increase i and try divide number again until number can't be divided
	// while dividing it print all dividers to console
	for i <= num {
		for num%i == 0 {
			fmt.Println(num, i)
			a := num / i // a = number after it divided
			num = a
		}
		i++
	}

}
