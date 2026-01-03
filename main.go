package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(divisible(2000, 3200))
	fmt.Println(factorial(8))
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

/*
Write a program which can compute the factorial of a given numbers. The results should be printed in a comma-separated sequence on a single line.
*/
func factorial(n int) int {
	if n == 0 || n == 1 || n < 0 {
		return 1
	}

	result := 1
	for i := n; i >= 1; i-- {
		result *= i
	}
	return result
}
