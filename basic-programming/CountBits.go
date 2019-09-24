package basic_programming

import (
	"fmt"
)

/*
Given a number N, print the number of set bits in the binary representation of this number.

Input:
The first contains a single integer T denoting the number of test cases. Each test case contains a single integer N

Output:
For each test case, print a single integer denoting the number of set bits in the binary representation of the given N .
 */

func PrintCountBits() {
	var N int
	fmt.Scanf("%d", &N)

	for i := 0;i < N; i += 1 {
		var num int
		fmt.Scanf("%d", &num)
		fmt.Println(CountBits(num))
	}
}

func CountBits(num int) int  {
	counter := 0
	for num > 0 {
		num = num & (num - 1) // flips rightmost bit everytime
		counter += 1
	}
	return counter
}