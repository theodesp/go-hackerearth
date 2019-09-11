package basic_programming

import (
	"bufio"
	"fmt"
	"os"
)

/*
Problem Description

Given a series of integer pairs  and , output the sum of  and .

Input Format

Each line contains two integers,  and . One input file may contain several pairs  where .

Output Format

Output a single integer per line - The sum of  and .

Constraints


Input
1 2
2 5
10 14

Output
3
7
24
*/

func PrintAPlusB() {
	var A int
	var B int
	var result []int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d %d", &A, &B)
		result = append(result, A + B)
	}

	for _, val := range result {
		fmt.Println(val)
	}
}
