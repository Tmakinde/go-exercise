package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(divisible(2000, 3200))
	fmt.Println(factorial(8))
	fmt.Printf("%v", "Squares: ")
	fmt.Printf("%v", squaredNumber(5))
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

/*
With a given integral number n, write a program to generate a map that contains (i, i*i) such that is an integral number between 1 and n (both included),
and then the program should print the map with representation of the value
*/
func squaredNumber(n int) map[int]int {
	result := make(map[int] int)
	for i := 1; i <= n; i++ {
		result[i] = i * i
	}
	return result
}