package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	books := [4]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	input := os.Args[1:] // get user input

	searchedBook := strings.ToLower(input[0])

	fmt.Println("Search Results:")

	var found bool
	i := 0

	for i <= len(books)-1 {
		book := books[i]
		if strings.Contains(strings.ToLower(book), searchedBook) {
			fmt.Printf("We have the book: %q\n", book)
			found = true
		}
		i++
	}

	if found == false {
		fmt.Printf("We don't have the book: %q\n", searchedBook)
	}
}
