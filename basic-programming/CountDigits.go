package basic_programming

import (
	"fmt"
	"strconv"
)

func PrintCountDigits()  {
	var S string
	fmt.Scanf("%s", &S)

	digitsMap := CountDigits(S)

	for i := 0;i<10; i+= 1 {
		fmt.Printf("%d %d\n", i, digitsMap[i])
	}
}


func CountDigits(s string) map[int]int   {
	result := make(map[int]int)
	for i := 0;i<10; i+= 1 {
		result[i] = 0
	}

	for _, char := range s {
		n, _ := strconv.Atoi(string(char))
		existing := result[n]
		result[n] = existing + 1
	}

	return result
}