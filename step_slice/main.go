package main

import (
	"bufio"
	"fmt"
	"os"
)

func rotate(input string, num int) {
	// our word is avatar the last airbender and number is 3
	leftFirst := input[0:num] // slice first part which is ava
	leftSecond := input[num:] // slice second part which is tar the last airbender

	// lets say our input is (I like trains) then its length is 13(including blanks) then we slice in like [0: (13-3)]
	// it will like like this: (I like tra)
	// sliced part is (ins) which contains in rightSecond variable in below
	rightFirst := input[0 : len(input)-num] // slice from right side.
	rightSecond := input[len(input)-num:]

	// for example our word is araba sevdasi and rotate number is 4
	fmt.Println("Left Rotation : ", (leftSecond + leftFirst))    // output look like this: a sevdasi arab
	fmt.Println("Right Rotation : ", (rightSecond + rightFirst)) // output look like this:  si araba sevda

}

func main() {
	fmt.Print("enter the string: ")
	inputReader := bufio.NewReader(os.Stdin) // get input from user
	input, _ := inputReader.ReadString('\n') // get whole input as one string
	fmt.Println(input)

	var num int // input number by user
	fmt.Print("enter the number you like rotate: ")
	_, err := fmt.Scan(&num) // get input from user
	if err != nil {          // looks for a error
		fmt.Println("Invalid inputs try again: err:", err)
	}
	fmt.Println(num)
	rotate(input, num)
}
