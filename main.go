package main

import (
	"fmt"
	basic_programming "github.com/theodesp/go-hackerearth/basic-programming"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	nums := basic_programming.GetPrimeNumbers(N)

	for _, num := range nums {
		fmt.Printf("%d ", num)
	}
}
