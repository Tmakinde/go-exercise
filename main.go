package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(divisible(2000, 3200))
}

/*
Write a program which will find all such numbers which are divisible by 7  but are not a multiple of 5, between 2000 and 3200 (both included).
The numbers obtained should be printed in a comma-separated sequence on a single line.
*/

func divisible(min int, max int) string {
	var result strings.Builder

	for i := min; i <= max; i++ {
		if (i%7 == 0 && i%5 !=0) {
			if result.Len() > 0 {
				result.WriteString(", ")
			}
			result.WriteString(strconv.Itoa(i))
		}
	}
	return result.String()
}
