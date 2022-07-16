package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 1. uncomment the code below
	var header []byte
	png := []byte{'P', 'N', 'G'}

	// 2. append elements to header to make it equal with the png slice
	header = append(header, png...)
	// 3. compare the slices using the bytes.Equal function
	// 4. print whether they're equal or not
	//bytes equal example fmt.Println(bytes.Equal([]byte("Go"), []byte("Go"))) output will be true
	if bytes.Equal(png, header) == true {
		fmt.Println("They are equal")
	} else if bytes.Equal(png, header) == false {
		fmt.Println("they're not equal")
	}

}
