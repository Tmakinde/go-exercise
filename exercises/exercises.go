package exercises

import (
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"net/http"
	"github.com/gin-gonic/gin"
	"os"
)

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

func basicFileOperations() {
	// Create a file
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer file.Close()
	// Write to the file
	// range of letter returns index and unicode code point of each letter
	character := "Hello, World!"
	for index, letter  := range character {
		file.WriteString(string(letter))
		if index < len(character) - 1 {
			file.WriteString("\n")
		}
	}
}

func readFromFile() (string, error) {
	file, err := os.Open("test.txt")
	if err != nil {
		return "", err
	}

	defer file.Close()
	var letters []string
	reader := bufio.NewScanner(file)

	for reader.Scan() {
		letters = append(letters, reader.Text())
	}
	return strings.Join(letters, ""), nil
}

func appendToFile() (string, error) {
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		return "", err
	}
	defer file.Close()

	numbers := []int{1,2,3,4,5}

	fileinfo, _ := os.Stat("test.txt")
	if fileinfo.Size() > 0 {
		file.WriteString("\n")
	}

	for _, number := range numbers {
		_, err := file.WriteString(strconv.Itoa(number) + "\n")
		if err != nil {
			return "something went wrong", err
		}
	}
	return "Numbers appended successfully", nil
}

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	id := c.Query("id")

	if id != "" {
		for _, a := range albums {
			if a.ID == id {
				c.IndentedJSON(http.StatusOK, a)
				return
			}
		}
	}
    c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(c *gin.Context) {
	var newAlbum album

	// bind incoming JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// add new album to slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}