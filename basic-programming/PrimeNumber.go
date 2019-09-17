package basic_programming

import "fmt"

/*
You are given an integer N. You need to print the series of all prime numbers till N.

Input Format

The first and only line of the input contains a single integer N denoting the number till where you need to find the series of prime number.

Output Format

Print the desired output in single line separated by spaces.

Constraints

1<=N<=1000
 */


/*
 Solution:
 Observation: A prime number N cannot be factored. So to find all the prime numbers til N
 then for each n <= N we try to find if any number factors the n. If yes we abort as we know that n is not a prime.
 */
func PrintPrimeNumbers() {
	var N int
	fmt.Scanf("%d", &N)

	nums := GetPrimeNumbers(N)

	for num := range nums {
		fmt.Printf("%d ", num)
	}
}

func GetPrimeNumbers(to int) []int {
	result := []int{}
	for i := 1; i <= to; i+=1 {
		isFactored := false
		for j := 2; j <= i/2; j+=1 {
			if i % j == 0 {
				isFactored = true
				break
			}
		}

		if !isFactored && i != 1 {
			result = append(result, i)
		}
	}
	return result
}