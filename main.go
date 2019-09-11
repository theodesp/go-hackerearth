package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	A := new(big.Int)
	B := new(big.Int)
	var result []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		sum := new(big.Int)
		n, _ := fmt.Sscanf(scanner.Text(), "%s %s", A, B)
		if n == 0 {
			break
		}
		sum.Add(A, B)
		result = append(result, sum.String())
	}

	for _, val := range result {
		fmt.Println(val)
	}
}
