package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"

		//  Location       Size           Beds           Baths          Price
		//  ===========================================================================
		//  New York       100            2              1              100000
		//  New York       150            3              2              200000
		//  Paris          200            4              3              400000
		//  Istanbul       500            10             5              1000000
		data = `New York,100,2,1,100000
New York,150,3,2,200000 
Paris,200,4,3,400000 
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		locs   []string
		sizes  []int
		beds   []int
		baths  []int
		prices []int
	)

	// HINTS
	//  + strings.Split function can separate a string into a []string
	rows := strings.Split(data, "\n") // row will be number of each data with seperated enter which is 4

	for _, row := range rows { // for each rows(4)
		cols := strings.Split(row, separator) // spliting each colums (locs, sizes, beds, baths, prices)
		// cols[0] = locs, cols[1] = sizes, cols[2] = beds, cols[3] = baths, cols[4] = prices
		locs = append(locs, cols[0]) // string array of locs
		// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int
		x, _ := strconv.Atoi(cols[1])
		sizes = append(sizes, x) // int array of size
		// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int
		x, _ = strconv.Atoi(cols[2])
		beds = append(beds, x) // int array of beds
		// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int
		x, _ = strconv.Atoi(cols[3])
		baths = append(baths, x) // int array of baths
		// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int
		x, _ = strconv.Atoi(cols[4])
		prices = append(prices, x) // int array of prices
	}

	//  5. Print the header
	//
	//     1. Separate it by using the separator
	//
	//     2. Print each column 15 character wide ("%-15s")
	for _, head := range strings.Split(header, separator) { //
		fmt.Printf("%-15s", head)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75)) //to make this : ==================================

	for i := range rows {
		// print datas in order 4 times
		fmt.Println()
		fmt.Printf("%-15s", locs[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", prices[i])
		fmt.Println()
		// without println its looks bad so I add two println
	}
}
