package fizzbuzz

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Print("Enter an integer number: ")
	var input int
	fmt.Scanf("%d", &input)

	for i := 1; 1 <= input; i++ {
		fmt.Println(fizzbuzz(i))
	}
}

func fizzbuzz(i int) string {
	var result string

	if i%3 == 0 {
		result = "Fizz"
	}

	if i%5 == 0 {
		result += "Buzz"
	}

	if len(result) == 0 {
		result = strconv.Itoa(i)
	}

	return result
}
